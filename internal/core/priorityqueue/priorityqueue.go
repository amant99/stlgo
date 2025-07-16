package priorityqueue

type Comparator[T any] func(a, b T) bool

type PriorityQueue[T any] interface {
    PushValue(value T)
    PopValue() (T, bool)
    Peek() (T, bool)
    Size() int
    IsEmpty() bool
}