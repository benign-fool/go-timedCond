package timedCond

import (
	"sync"
	"time"
)

//Wait waits for either the supplied sync.Cond to be signalled or the
//supplied duration to be timed out.  It returns a bool based on the
//condition being signalled.  True for signalled, false for timedout.
func Wait(cond *sync.Cond, timeout time.Duration) (condSuccess bool) {
	condChan := make(chan interface{}, 1)
	go startCond(cond, condChan)
	timer := time.NewTimer(timeout)

	select {
	case <-timer.C:
		return false
	case <-condChan:
		return true
	}
}

func startCond(cond *sync.Cond, condChan chan interface{}) {
	defer close(condChan)
	defer cond.L.Unlock()
	cond.L.Lock()
	cond.Wait()
	condChan <- true
}
