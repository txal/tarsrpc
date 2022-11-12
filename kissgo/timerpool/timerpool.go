// @author kordenlu
// @创建时间 2018/03/09 11:03
// 功能描述:

package timerpool

import (
	"sync"
	"time"
)

var timerPool sync.Pool

// AcquireTimer returns time from pool if possible.
func AcquireTimer(d time.Duration) *time.Timer {
	v := timerPool.Get()
	if v == nil {
		return time.NewTimer(d)
	}
	tm := v.(*time.Timer)
	if tm.Reset(d) {
		// active timer?
		return time.NewTimer(d)
	}
	return tm
}

// ReleaseTimer returns timer into pool.
func ReleaseTimer(tm *time.Timer) {
	if !tm.Stop() {
		// tm.Stop() returns false if the timer has already expired or been stopped.
		// We can't be sure that timer.C will not be filled after timer.Stop(),
		// see https://groups.google.com/forum/#!topic/golang-nuts/-8O3AknKpwk
		//
		// The tip from manual to read from timer.C possibly blocks caller if caller
		// has already done <-timer.C. Non-blocking read from timer.C with select does
		// not help either because send is done concurrently from another goroutine.
		return
	}
	timerPool.Put(tm)
}
