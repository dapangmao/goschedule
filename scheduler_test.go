package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

	job = NewActionOnlyJob(1)
	ui2sched <- Command{job, remove}

	time.Sleep(time.Second * 2)

	Entries.Lock()
	t.Log("Debug the result map", Entries.data)

	assert.Equal(t, "remove", Entries.data[1].message)
	assert.Equal(t, "fetched", Entries.data[2].message)
	assert.Equal(t, "fetched", Entries.data[3].message)
	Entries.Unlock()

}
