package main

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)


func init() {
	Fetch = func(id int) {
		sched2ui <- Feedback{"fetched", time.Now(), id}
	}

	timeAfter = func(t time.Duration) <-chan time.Time {
		return time.After(time.Second)
	}
}

func TestScheduler(t *testing.T) {

	var id = 1
	sched := NewScheduler()
	go sched.Serve()
	go getStats()

	s := "every 5 hour"
	p := Parser{id, s}
	job, _ := p.Parse()

	ui2sched <- Command{job, add}
	ui2sched <- Command{job, stop}
	ui2sched <- Command{job, restart}

	s = "every 10 hour"
	p = Parser{id, s}
	job, _ = p.Parse()
	ui2sched <- Command{job, update}

	id = 2
	s = "every Monday 10:30"
	p = Parser{id, s}
	job, _ = p.Parse()
	ui2sched <- Command{job, add}

	id = 3
	s = "every 2:30:4"
	p = Parser{id, s}
	job, _ = p.Parse()
	ui2sched <- Command{job, add}

	time.Sleep(time.Second * 1)

	job = NewActionOnlyJob(1)
	ui2sched <- Command{job, remove}


	t.Log("Debug the result map", stats)

	assert.Equal(t, stats[1].message, "remove")
	assert.Equal(t, stats[2].message, "fetched")
	assert.Equal(t, stats[3].message, "fetched")
	assert.Equal(t, len(sched.jobMap), 2)

}
