package stackadapter

import "sync"

type arrayStack[T any] struct {
	data []T
	lock sync.Mutex
}

func NewArrayStack[T any]() *arrayStack[T] {
	return &arrayStack[T]{data: make([]T, 0)}
}

func (s *arrayStack[T]) Push(value T) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data = append(s.data, value)
}

func (s *arrayStack[T]) Pop() (T, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val, true
}

func (s *arrayStack[T]) Peek() (T, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	return s.data[len(s.data)-1], true
}

func (s *arrayStack[T]) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return len(s.data)
}

func (s *arrayStack[T]) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return len(s.data) == 0
}
