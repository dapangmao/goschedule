package main

import (
	"errors"
	"time"
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

type scheduled interface {
	nextRun() (time.Duration, error)
}

// Job defines a running job and allows to stop a scheduled job or run it.
type Job struct {
	id        int
	err       error
	sched  scheduled
	isStopped bool
}

type recurrent struct {
	units  int
	period time.Duration
	done   bool
}

func (r *recurrent) nextRun() (time.Duration, error) {
	if r.units == 0 || r.period == 0 {
		return 0, errors.New("cannot set recurrent time with 0")
	}
	if !r.done {
		r.done = true
		return 0, nil
	}
	return time.Duration(r.units) * r.period, nil
}

type daily struct {
	hour int
	min  int
	sec  int
}


func (d daily) nextRun() (time.Duration, error) {
	now := time.Now()
	year, month, day := now.Date()
	date := time.Date(year, month, day, d.hour, d.min, d.sec, 0, time.Local)
	if now.Before(date) {
		return date.Sub(now), nil
	}
	date = time.Date(year, month, day+1, d.hour, d.min, d.sec, 0, time.Local)
	return date.Sub(now), nil
}

type weekly struct {
	day time.Weekday
	d   daily
}

func (w weekly) nextRun() (time.Duration, error) {
	now := time.Now()
	year, month, day := now.Date()
	numDays := w.day - now.Weekday()
	if numDays == 0 {
		numDays = 7
	} else if numDays < 0 {
		numDays += 7
	}
	date := time.Date(year, month, day+int(numDays), w.d.hour, w.d.min, w.d.sec, 0, time.Local)
	return date.Sub(now), nil
}



type ErrorMsg struct {
	err error
	id int
}



type Scheduler struct {
	jobMap map[int]chan string
}


func (s *Scheduler) runNew(job *Job) {
	var newChan = make(chan string)
	s.jobMap[job.id] = newChan
	go s.RunJob(job, newChan)

}
func (s *Scheduler) Run() {
	for current := range ui2sched {
		job, command := current.j, current.cmd
		if command == "add" {
			s.runNew(job)
		} else if command == "delete" || command == "modify" {
			s.jobMap[job.id] <- "delete"
			delete(s.jobMap, job.id)
			if command == "modify" {
				s.runNew(job)
			}
		} else {
			s.jobMap[job.id] <- command
		}
	}
}

// Run sets the job to the schedule and returns the pointer to the job so it may be
// stopped or executed without waiting or an error.
func (s *Scheduler) RunJob(j *Job, cmdChan<- chan string)  {
	next, err := j.sched.nextRun()
	if err != nil {
		sched2ui <- ErrorMsg{err, j.id}
	}
	for {
		select {
		case cmd := <- cmdChan:
			switch cmd {
			case "stop":
				j.isStopped = true
			case "start":
				j.isStopped = false
			case "runNow":
				go crawl()
			case "delete":
				return
			}
		case <-time.After(next):
			if !j.isStopped {
				go crawl()
				next, err = j.sched.nextRun()
				if err != nil {
					sched2ui <- ErrorMsg{err, j.id}
				}
			}
		}
	}
}


func crawl() {
	doc, _ := goquery.NewDocument("http://metalsucks.net")
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
}
