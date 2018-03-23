package main

import (
	"time"
)


const (
	remove = "remove"
	add = "add"
	stop = "stop"
	restart = "restart"
	runNow = "runNow"
	update = "update"
)


var timeAfter = time.After
var ui2sched = make(chan Command)
var sched2ui = make(chan Feedback)


type Command struct {
	j   *Job
	cmd string
}



type scheduled interface {
	nextRun() time.Duration
}

// Job defines a running job and allows to stop a scheduled job or run it.
type Job struct {
	id        int
	sched     scheduled
	isStopped bool
	isOneTime bool
}



type recurrent struct {
	quantity int
	unit     time.Duration
}

func (r *recurrent) nextRun() time.Duration {
	return time.Duration(r.quantity) * r.unit
}

type daily struct {
	hour int
	min  int
	sec  int
}

func (d daily) nextRun() time.Duration{
	now := time.Now()
	year, month, day := now.Date()
	date := time.Date(year, month, day, d.hour, d.min, d.sec, 0, time.Local)
	if now.Before(date) {
		return date.Sub(now)
	}
	date = time.Date(year, month, day+1, d.hour, d.min, d.sec, 0, time.Local)
	return date.Sub(now)
}

type weekly struct {
	day time.Weekday
	d   *daily
}

func (w weekly) nextRun() time.Duration {
	now := time.Now()
	year, month, day := now.Date()
	numDays := w.day - now.Weekday()
	if numDays == 0 {
		numDays = 7
	} else if numDays < 0 {
		numDays += 7
	}
	date := time.Date(year, month, day+int(numDays), w.d.hour, w.d.min, w.d.sec, 0, time.Local)
	return date.Sub(now)
}

type Feedback struct {
	message string
	time time.Time
	id  int
}

type Scheduler struct {
	jobMap map[int]chan string
}

func NewNilJob(id int) *Job {
	return &Job{id, nil, false, false}
}


func NewScheduler() *Scheduler {
	j := make(map[int]chan string)
	return &Scheduler{j}
}


func (s *Scheduler) runNewJob(job *Job) {
	var newJobChan = make(chan string)
	s.jobMap[job.id] = newJobChan
	go s.RunJob(job, newJobChan)
	sched2ui <- Feedback{add , time.Now(), job.id}
}

func (s *Scheduler) Serve() {
	for current := range ui2sched {
		job, command := current.j, current.cmd
		if command == add {
			s.runNewJob(job)
		} else if command == remove || command == update {
			s.jobMap[job.id] <- remove
			delete(s.jobMap, job.id)
			if command == update {
				s.runNewJob(job)
			}
		} else {
			s.jobMap[job.id] <- command
		}
	}
}

// Run sets the job to the schedule and returns the pointer to the job so it may be
// stopped or executed without waiting or an error.
func (s *Scheduler) RunJob(j *Job, jobCmdChan <- chan string) {
	next := j.sched.nextRun()
	for {
		select {
		case cmd := <- jobCmdChan:
			sched2ui <- Feedback{cmd , time.Now(),j.id}
			switch cmd {
			case stop:
				j.isStopped = true
			case restart:
				j.isStopped = false
			case runNow:
				go fetch(j.id)
			case remove:
				return
			}
		case <- timeAfter(next):
			if !j.isStopped {
				go Fetch(j.id)
				next= j.sched.nextRun()
			}
			if j.isOneTime {
				return
			}
		}
	}
}

