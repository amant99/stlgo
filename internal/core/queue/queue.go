package queue

type Queue[T any] interface {
	Enqueue(value T)
	Dequeue() (T, bool)
	Peek() (T, bool)
	Size() int
	IsEmpty() bool
}
