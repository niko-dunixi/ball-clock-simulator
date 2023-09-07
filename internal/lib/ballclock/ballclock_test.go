package ballclock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBallClock(t *testing.T) {

	t.Run("Count(30),Ticks(325)", func(t *testing.T) {
		ballClock := New(30)
		ballClock.TickMinutes(325)
		// assert.ElementsMatch(t, []int{}, ballClock.SinglesRegister, "Minutes register was wrong")
		assert.ElementsMatch(t, []int{22, 13, 25, 3, 7}, ballClock.FifthsRegister, "Fifths Minutes register was wrong")
		// assert.ElementsMatch(t, []int{6, 12, 17, 4, 15}, ballClock.hoursRegister)
		// assert.ElementsMatch(t, []int{11, 5, 26, 18, 2, 30, 19, 8, 24, 10, 29, 16, 21, 28, 1, 23, 14, 27, 9}, ballClock.mainBuffer)
	})
}
