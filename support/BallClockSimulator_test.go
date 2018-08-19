package support

import (
	"testing"
	"fmt"
)

func TestClockMatchesPDFIterations(t *testing.T) {
	ballClock := NewBallClock(30)
	ballClock.TickMinutes(325)
	fmt.Println(ballClock.JsonString())
	assertIntSlicesEqual(t, ballClock.Min, []int{})
	assertIntSlicesEqual(t, ballClock.FiveMin, []int{22, 13, 25, 3, 7})
	assertIntSlicesEqual(t, ballClock.Hour, []int{6, 12, 17, 4, 15})
	assertIntSlicesEqual(t, ballClock.Main, []int{11, 5, 26, 18, 2, 30, 19, 8, 24, 10, 29, 20, 16, 21, 28, 1, 23, 14, 27, 9})
}

func TestClockCycleDaysOne(t *testing.T) {
	_, daysUntilReset, _ := CalculateDaysUntilReset(30)
	if daysUntilReset != 15 {
		t.FailNow()
	}
}

func TestClockCycleDaysTwo(t *testing.T) {
	_, daysUntilReset, _ := CalculateDaysUntilReset(45)
	if daysUntilReset != 378 {
		t.FailNow()
	}
}