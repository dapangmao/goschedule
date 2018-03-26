package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseRecurrent(t *testing.T) {
	s := "every 5 hours"
	p := Parser{1, s}
	result, err := p.Parse()
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 1, result.id)
	assert.Equal(t, 5*time.Hour, result.sched.nextRun())
	assert.Equal(t, false, result.isStopped)
	assert.Equal(t, false, result.isOneTime)
}
