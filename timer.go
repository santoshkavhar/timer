package timer

import (
	"time"
)

// Timer implementation with pause
type pausyTimerInterface interface {
	Pause() (time.Duration, bool)  // Duration Elapsed
	Resume() (time.Duration, bool) // Duration Remaining
	Stop() (time.Duration, bool)   // Duration Elapsed
}

type timerState uint8

const (
	running = timerState(iota)
	paused
	stopped
)

type pausyTimer struct {
	timer           *time.Timer
	startTime       time.Time
	elapsedDuration time.Duration
	totalDuration   time.Duration
	state           timerState
}
