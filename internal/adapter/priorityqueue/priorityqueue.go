package pqadapter

import (
	"container/heap"
	"sync"
)

type item[T any] struct {
	value T
	index int
}

type priorityQueue[T any] struct {
	data       []*item[T]
	comparator func(a, b T) bool
	lock       sync.Mutex
}

func NewPriorityQueue[T any](cmp func(a, b T) bool) *priorityQueue[T] {
	pq := &priorityQueue[T]{
		data:       []*item[T]{},
		comparator: cmp,
	}
	heap.Init(pq) // pq now implements heap.Interface directly
	return pq
}

// These methods are required to implement heap.Interface
func (pq *priorityQueue[T]) Len() int { return len(pq.data) }
func (pq *priorityQueue[T]) Less(i, j int) bool {
	return pq.comparator(pq.data[i].value, pq.data[j].value)
}
func (pq *priorityQueue[T]) Swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
	pq.data[i].index = i
	pq.data[j].index = j
}
func (pq *priorityQueue[T]) Push(x any) {
	itm := x.(*item[T])
	pq.data = append(pq.data, itm)
}
func (pq *priorityQueue[T]) Pop() any {
	old := pq.data
	n := len(old)
	itm := old[n-1]
	pq.data = old[0 : n-1]
	return itm
}

// Our custom public methods
func (pq *priorityQueue[T]) PushValue(value T) {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	heap.Push(pq, &item[T]{value: value})
}

func (pq *priorityQueue[T]) PopValue() (T, bool) {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	var zero T
	if len(pq.data) == 0 {
		return zero, false
	}
	itm := heap.Pop(pq).(*item[T])
	return itm.value, true
}

func (pq *priorityQueue[T]) Peek() (T, bool) {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	var zero T
	if len(pq.data) == 0 {
		return zero, false
	}
	return pq.data[0].value, true
}

func (pq *priorityQueue[T]) Size() int {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	return len(pq.data)
}

func (pq *priorityQueue[T]) IsEmpty() bool {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	return len(pq.data) == 0
}
