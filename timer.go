package timer

import (
	"time"
)

// PausyInterface is the interface with pause, resume and stop with Durations
type PausyInterface interface {
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

type Timer struct {
	timer           *time.Timer
	startTime       time.Time
	elapsedDuration time.Duration
	totalDuration   time.Duration
	state           timerState
}

func NewTimer(duration time.Duration) (timerWithPause Timer) {
	newTimer := time.NewTimer(duration)
	timerWithPause.startTime = time.Now()
	timerWithPause.totalDuration = duration
	timerWithPause.timer = newTimer
	return
}
