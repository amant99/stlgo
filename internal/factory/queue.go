package factory

import (
    "github.com/amant99/stlgo/internal/core/queue"
    queueadapter "github.com/amant99/stlgo/internal/adapter/queue"
)

func NewQueue[T any]() queue.Queue[T] {
    return queueadapter.NewArrayQueue[T]()
}
