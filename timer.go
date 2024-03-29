package timer

import (
	"time"
)


// PausyTimer with pause, resume and stop with Durations
type PausyTimer interface {
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

// NewTimer returns a Timer object which implements PausyTimer
func NewTimer(duration time.Duration) (timerWithPause Timer) {
	newTimerWithPause := time.NewTimer(duration)
	timerWithPause.startTime = time.Now()
	timerWithPause.totalDuration = duration
	timerWithPause.timer = newTimerWithPause
	return
}

// Pause stops the timer temporarily if the timer was running
// and it returns the elapsed duration of timer
// pT is simply an abbrevation for pausyTimer ;-)
func (pT *Timer) Pause() (ElapsedDuration time.Duration, PausedSuccessfully bool) {
	// we are fetching currentTime here itself to increase accuracy
	currentTime := time.Now()
	switch pT.state {
	case running:
		// Pause for 291.6 years
		// timerWasActive will become false whenever timer
		// was stopped or expired
		timerWasActive := pT.timer.Reset(2<<61 - 1)
		if timerWasActive {
			pT.elapsedDuration += currentTime.Sub(pT.startTime)
			pT.state = paused
			return pT.elapsedDuration, true
		}
		// Timer was Stopped
		pT.state = stopped
		return pT.totalDuration, false
	// if already paused then return elapsedTime
	case paused:
		return pT.elapsedDuration, false
	}
	// There is nothing to Pause, return Timer duration
	return pT.totalDuration, false
}

// Resume returns the remaining Duration,
// and if the timer was resumed successfully.
func (pT *Timer) Resume() (RemainingDuration time.Duration, ResumedSuccessfully bool) {
	// we are fetching currentTime here itself to increase accuracy
	currentTime := time.Now()

	switch pT.state {
	// Timer is already running, there is nothing to resume
	// return remaining duration and false for resumed
	case running:
		// if timer has completed its run but hasn't reported its stoppage
		if currentTime.Sub(pT.startTime) > pT.totalDuration {
			// TODO : Handle this situation more gracefully
			pT.state = stopped
			pT.elapsedDuration = pT.totalDuration
			return 0, false
		}
		// if the timer is still running
		return pT.totalDuration - currentTime.Sub(pT.startTime), false

	// if the timer was paused then
	case paused:
		resetWasSuccessful := pT.timer.Reset(pT.totalDuration - pT.elapsedDuration)
		// timer was expired but it didn't report
		// Just create a new timer and replace the old one
		if !resetWasSuccessful {
			newTimerWithPause := time.NewTimer(pT.totalDuration - pT.elapsedDuration)
			pT.timer = newTimerWithPause
		}
		// both these field setting is common despite of
		// whether timer was reset Successfully or not
		pT.startTime = time.Now()
		pT.state = running
		return pT.totalDuration - pT.elapsedDuration, true
	}
	// Only option remaining is that the timer had stopped
	// we return what duration more it could have run ,i.e Remaining Duration
	// You can't resume such a timer
	return pT.totalDuration - pT.elapsedDuration, false
}

// Stop literally stops the timer
// we can't resume such a timer
func (pT *Timer) Stop() (ElapsedDuration time.Duration, StoppedSuccessfully bool) {
	// we are fetching currentTime here itself to increase accuracy
	currentTime := time.Now()

	switch pT.state {
	// running timer will have to stop successfully
	case running:
		stopWasSuccessfull := pT.timer.Stop()
		pT.state = stopped
		if stopWasSuccessfull {
			pT.elapsedDuration += currentTime.Sub(pT.startTime)
			return pT.elapsedDuration, true
		}
		// else
		// There is nothing to Stop
		// Timer didn't report its expiry
		pT.elapsedDuration = pT.totalDuration
		return pT.totalDuration, false
	case paused:
		// timer that was paused will always stop Except when expired
		// that expiration is after 291 years, so no worries
		pT.timer.Stop()
		pT.state = stopped
		return pT.elapsedDuration, true
	}
	// if already stopped then return elapsedTime and false
	return pT.elapsedDuration, false
}
