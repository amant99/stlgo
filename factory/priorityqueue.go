package factory

import (
	pqadapter "github.com/amant99/stlgo/internal/adapter/priorityqueue"
	"github.com/amant99/stlgo/internal/core/priorityqueue"
)

func NewPriorityQueue[T any](cmp priorityqueue.Comparator[T]) priorityqueue.PriorityQueue[T] {
	return pqadapter.NewPriorityQueue(cmp)
}
