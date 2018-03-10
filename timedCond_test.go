package timedCond

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestWaitCond(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})
	timeoutDuration := 1 * time.Second
	go func() {
		waitForSignalSimulationTimer := time.NewTimer(500 * time.Millisecond)
		<-waitForSignalSimulationTimer.C
		cond.Signal()
	}()
	condSuccess := Wait(cond, timeoutDuration)

	assert.Equal(t, true, condSuccess)
}

func TestWaitTimeout(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})
	timeoutDuration := 500 * time.Millisecond
	go func() {
		waitForSignalSimulationTimer := time.NewTimer(1 * time.Second)
		<-waitForSignalSimulationTimer.C
		cond.Signal()
	}()
	condSuccess := Wait(cond, timeoutDuration)

	assert.Equal(t, false, condSuccess)
}

func TestStartCond(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})
	condChan := make(chan interface{}, 1)

	go startCond(cond, condChan)

	timer := time.NewTimer(1 * time.Second)
	<-timer.C
	cond.Signal()
	<-condChan
}
