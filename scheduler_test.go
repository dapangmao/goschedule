package main

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

var id = 1

func init() {
	Fetch = func() {
		sched2ui <- Feedback{"fetched", time.Now(), id}
	}

	timeAfter = func(t time.Duration) <-chan time.Time {
		return time.After(time.Second)
	}
}



func TestScheduler(t *testing.T) {

	id = 1
	sched := NewScheduler()
	go sched.Serve()

	go func(){
		var j int
		var wanted = []string{"add",  "stop", "restart",  "fetched",   "add", "fetched"}
		for x := range sched2ui {
			assert.Equal(t, x.message, wanted[j])
			t.Log(x)
			j++
		}
	}()

	s := "every 5 hour"
	p := Parser{id, s}
	job, _ := p.Parse()
	
	ui2sched <- Command{job, add}
	ui2sched <- Command{job, stop}
	ui2sched <- Command{job, restart}

	time.Sleep(time.Second * 1)

	s = "every 10 hour"
	p = Parser{1, s}
	job, _ = p.Parse()
	ui2sched <- Command{job, update}
	time.Sleep(time.Second * 1)
}
