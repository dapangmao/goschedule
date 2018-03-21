package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
	"fmt"
)

func TestScheduler(t *testing.T) {
	fetch = func() {
		sched2ui <- Feedback{"this is test", 1}
	}

	sched := NewScheduler()
	go sched.Serve()


	fmt.Print("here")
	s := "every 1 seconds"
	p := Parser{1, s}
	job, _ := p.Parse()
	ui2sched <- Command{job, "add"}

	for {
		select {
		case x := <- sched2ui:
			assert.Equal(t, x.message, "this is test")
			assert.Equal(t, x.id, 1)
			return
		case <- time.After(time.Second * 5):
			t.Error("The scheduler did not get result")
			return
		}
	}
}
