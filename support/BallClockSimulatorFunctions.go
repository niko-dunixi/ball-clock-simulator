package support

import (
	"fmt"
	"time"
)

// ------
// Functions vs Methods
//
// The overall accepted difference between functions and methods is that a function
// takes an input and gives an output it doesn't require state. Methods refer to essentially
// the same thing, but in OOP they act upon state of an object. I've separated these from the
// rest of the ball clock methods because they don't effect state of the struct.
// ------

func CalculateDaysUntilResetString(ballCount int) string {
	clock, days, seconds := CalculateDaysUntilReset(ballCount)
	// We have to calculate millis ourselves https://github.com/golang/go/issues/5491
	millis := int(seconds * 1e3)
	// https://golang.org/pkg/fmt/
	resultString := fmt.Sprintf("%d balls cycle after %d days.\nCompleted in %d milliseconds (%.3f seconds)\n", clock.count, days, millis, seconds)
	return resultString
}

func CalculateDaysUntilReset(ballCount int) (*BallClock, int, float64) {
	clock := NewBallClock(ballCount)
	days := 0
	start := time.Now()
	for {
		clock.TickDay()
		days++
		if clock.IsInitialOrdering() {
			break
		}
	}
	duration := time.Since(start)
	return &clock, days, duration.Seconds()
}
