package ballclock

import "github.com/niko-dunixi/ball-clock/simulator/internal/lib/ballclock/ringbuffer"

type Unit int

const (
	Minute Unit = iota
	Hour
	Day
)

type Ballclock interface {
	Tick(unit Unit, quantity uint)
}

type SimpleBallclock struct {
	// mainBufferStartIndex int
	// mainBufferEndIndex   int
	// mainBuffer           []int
	MainBuffer      ringbuffer.Queue[int] `json:"Main"`
	SinglesRegister []int                 `json:"Min"`
	FifthsRegister  []int                 `json:"FiveMin"`
	HoursRegister   []int                 `json:"Hour"`
}

func New(ballCount int) SimpleBallclock {
	mainBuffer := make([]int, ballCount)
	for i := 0; i < len(mainBuffer); i++ {
		mainBuffer[i] = i + 1
	}

	return SimpleBallclock{
		// mainBufferStartIndex: 0,
		// mainBufferEndIndex:   0,
		// mainBuffer:           mainBuffer,
		MainBuffer:      ringbuffer.NewNaiveQueue[int](mainBuffer),
		SinglesRegister: make([]int, 0, 4),
		FifthsRegister:  make([]int, 0, 11),
		HoursRegister:   make([]int, 0, 11),
	}
}

func (s *SimpleBallclock) TickMinutes(count uint) {
	for i := 0; i < int(count); i++ {
		s.TickMinute()
	}
}

func (s *SimpleBallclock) TickMinute() {
	// tickBall := s.mainBuffer[s.mainBufferStartIndex]
	// s.mainBuffer[s.mainBufferStartIndex] = tickBall * -1
	// s.mainBufferStartIndex = (s.mainBufferStartIndex + 1) % len(s.mainBuffer)
	tickBall := s.MainBuffer.Pop()
	// Handle single minutes register
	if len(s.SinglesRegister) != cap(s.SinglesRegister) {
		s.SinglesRegister = append(s.SinglesRegister, tickBall)
		return
	}
	for i := len(s.SinglesRegister) - 1; i >= 0; i-- {
		currentBall := s.SinglesRegister[i]
		s.MainBuffer.Push(currentBall)
		// s.mainBuffer[s.mainBufferEndIndex] = currentBall
		// s.mainBufferEndIndex = (s.mainBufferEndIndex + 1) % len(s.mainBuffer)
	}
	s.SinglesRegister = s.SinglesRegister[:0]
	// Handle fifths minutes register
	if len(s.FifthsRegister) != cap(s.FifthsRegister) {
		s.FifthsRegister = append(s.FifthsRegister, tickBall)
		return
	}
	for i := len(s.FifthsRegister) - 1; i >= 0; i-- {
		currentBall := s.FifthsRegister[i]
		// s.mainBuffer[s.mainBufferEndIndex] = currentBall
		// s.mainBufferEndIndex = (s.mainBufferEndIndex + 1) % len(s.mainBuffer)
		s.MainBuffer.Push(currentBall)
	}
	s.FifthsRegister = s.FifthsRegister[:0]
	// Handle hours register
	if len(s.FifthsRegister) != cap(s.FifthsRegister) {
		s.FifthsRegister = append(s.FifthsRegister, tickBall)
		return
	}
	for i := len(s.FifthsRegister) - 1; i >= 0; i-- {
		currentBall := s.FifthsRegister[i]
		// s.mainBuffer[s.mainBufferEndIndex] = currentBall
		// s.mainBufferEndIndex = (s.mainBufferEndIndex + 1) % len(s.mainBuffer)
		s.MainBuffer.Push(currentBall)
	}
	s.FifthsRegister = s.FifthsRegister[:0]
	// We've cycled fully, so the ball is returned to the main buffer
	s.MainBuffer.Push(tickBall)
	// s.mainBuffer[s.mainBufferEndIndex] = tickBall
	// s.mainBufferEndIndex = (s.mainBufferEndIndex + 1) % len(s.mainBuffer)
}
