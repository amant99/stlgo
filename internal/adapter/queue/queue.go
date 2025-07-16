package queueadapter

import "sync"

type arrayQueue[T any] struct {
	data []T
	lock sync.Mutex
}

func NewArrayQueue[T any]() *arrayQueue[T] {
	return &arrayQueue[T]{data: make([]T, 0)}
}

func (q *arrayQueue[T]) Enqueue(value T) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.data = append(q.data, value)
}

func (q *arrayQueue[T]) Dequeue() (T, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	var zero T
	if len(q.data) == 0 {
		return zero, false
	}
	val := q.data[0]
	q.data = q.data[1:]
	return val, true
}

func (q *arrayQueue[T]) Peek() (T, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()
	var zero T
	if len(q.data) == 0 {
		return zero, false
	}
	return q.data[0], true
}

func (q *arrayQueue[T]) Size() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.data)
}

func (q *arrayQueue[T]) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return len(q.data) == 0
}
