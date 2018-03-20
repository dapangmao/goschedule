package main

import (
	"errors"
	"time"
	"strings"
	"strconv"
)


var ui2sched = make(chan command)
var sched2ui = make(chan ErrorMsg)


type command struct {
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
	d   daily
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

type ErrorMsg struct {
	err error
	id  int
}

type Scheduler struct {
	jobMap map[int]chan string
}


func NewScheduler() *Scheduler {
	j := make(map[int]chan string)
	return &Scheduler{j}
}


func (s *Scheduler) runNewJob(job *Job) {
	var newJobChan = make(chan string)
	s.jobMap[job.id] = newJobChan
	go s.RunJob(job, newJobChan)
}

func (s *Scheduler) Serve() {
	for current := range ui2sched {
		job, command := current.j, current.cmd
		if command == "add" {
			s.runNewJob(job)
		} else if command == "delete" || command == "modify" {
			s.jobMap[job.id] <- "delete"
			delete(s.jobMap, job.id)
			if command == "modify" {
				s.runNewJob(job)
			}
		} else {
			s.jobMap[job.id] <- command
		}
	}
}

// Run sets the job to the schedule and returns the pointer to the job so it may be
// stopped or executed without waiting or an error.
func (s *Scheduler) RunJob(j *Job, jobCmdChan <-chan string) {
	next := j.sched.nextRun()

	for {
		select {
		case cmd := <- jobCmdChan:
			switch cmd {
			case "stop":
				j.isStopped = true
			case "start":
				j.isStopped = false
			case "runNow":
				go fetch()
			case "delete":
				return
			}
		case <- time.After(next):
			if !j.isStopped {
				go fetch()
				next= j.sched.nextRun()
			}
		}
		if j.isOneTime {
			return
		}
	}
}

// every 5 hours
// every day at 10:30am
// every monday at 10:30
func Parse(s string, id int) (*Job, error) {
	notValid:
		return nil, errors.New("Not a valid input format for a job.")
	n := len(s)
	if n == 0 {goto notValid}
	s = strings.ToLower(s)
	splits := strings.Split(s, " ")
	if len(splits) < 2 {goto notValid}

	if !strings.Contains(s, "at") {
		isOneTime := true
		q, u := splits[0], splits[1]
		if splits[0] == "every" {
			q, u = splits[1], splits[2]
			isOneTime = false
		}
		quantity, err := strconv.Atoi(q)
		if err != nil {goto notValid}
		var unit time.Duration
		switch u {
		case "hour":
			unit = time.Hour
		case "minute":
			unit = time.Minute
		case "second":
			unit = time.Second
		default:
			goto notValid
		}
		current := &Job{id, &recurrent{quantity, unit}, false, isOneTime}
		return current, nil
	}

	if strings.Contains(s, "day") {
		idx := 0
		for i, x := range splits {
			if x == "at" || x == "on" {
				idx = i
				break
			}
		}

	}



}