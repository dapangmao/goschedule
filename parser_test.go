package main

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)



func TestParseRecurrent(t *testing.T) {
	s := "every 5 hours"
	p := Parser{1, s}
	result, err := p.Parse()
	assert.Nil(t, err,"Error should be nil")
	assert.Equal(t, result.id, 1)
	assert.Equal(t, result.sched.nextRun(), 5 * time.Hour)
	assert.Equal(t, result.isStopped, false)
	assert.Equal(t, result.isOneTime, false)
}

