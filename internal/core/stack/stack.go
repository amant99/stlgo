package stack

type Stack[T any] interface {
	Push(value T)
	Pop() (T, bool)
	Peek() (T, bool)
	Size() int
	IsEmpty() bool
}
