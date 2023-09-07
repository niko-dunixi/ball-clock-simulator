package ringbuffer

import (
	"encoding/json"
)

type Queue[T any] interface {
	Pop() T
	Push(value T)
}

type NaiveQueue[T any] struct {
	queue      []T
	startIndex *int
	endIndex   *int
}

// Utilizes a given pre-initialized slice as a queue
func NewNaiveQueue[T any](slice []T) NaiveQueue[T] {
	return NaiveQueue[T]{
		queue:      slice,
		startIndex: new(int),
		endIndex:   new(int),
	}
}

func (nq NaiveQueue[T]) Pop() T {
	value := nq.queue[*nq.startIndex]
	*nq.startIndex = (*nq.startIndex + 1) % len(nq.queue)
	return value
}

func (nq NaiveQueue[T]) Push(value T) {
	nq.queue[*nq.endIndex] = value
	*nq.endIndex = (*nq.endIndex + 1) % len(nq.queue)
}

func (nq NaiveQueue[T]) MarshalJSON() ([]byte, error) {
	presentationBuffer := make([]T, 0, len(nq.queue))
	if *nq.startIndex < *nq.endIndex {
		for i := *nq.startIndex; i < *nq.endIndex; i++ {
			presentationBuffer = append(presentationBuffer, nq.queue[i])
		}
		return json.Marshal(presentationBuffer)
	}
	for i := *nq.startIndex; i < len(nq.queue); i++ {
		presentationBuffer = append(presentationBuffer, nq.queue[i])
	}
	for i := 0; i < *nq.endIndex; i++ {
		presentationBuffer = append(presentationBuffer, nq.queue[i])
	}
	return json.Marshal(presentationBuffer)
}
