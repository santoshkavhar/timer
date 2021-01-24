package timer

import (
	"time"
)

// PausyInterface with pause, resume and stop with Durations
type PausyInterface interface {
	Pause() (time.Duration, bool)  // Duration Elapsed
	Resume() (time.Duration, bool) // Duration Remaining
	Stop() (time.Duration, bool)   // Duration Elapsed
}

// timerState will help us keep track of timer satus
type timerState uint8

const (
	running = timerState(iota)
	paused
	stopped
)

// Timer will help us implement a timer with Pause functionality
type Timer struct {
	// timer stores the underlying timer
	timer *time.Timer
	// startTime is the time when timer was started or restarted
	startTime time.Time
	// elapsedDuration is the Duration that has been completed by the timer
	elapsedDuration time.Duration
	// totalDuration is the Duration set by user
	totalDuration time.Duration
	// state helps us know the current status of timer
	state timerState
}

// NewTimer returns a timer which implements PausyInterface
func NewTimer(duration time.Duration) (timerWithPause Timer) {
	newTimerWithPause := time.NewTimer(duration)
	timerWithPause.startTime = time.Now()
	timerWithPause.totalDuration = duration
	timerWithPause.timer = newTimerWithPause
	return
}
