package factory

import (
	stackadapter "github.com/amant99/stlgo/internal/adapter/stack"
	"github.com/amant99/stlgo/internal/core/stack"
)

func NewStack[T any]() stack.Stack[T] {
	return stackadapter.NewArrayStack[T]()
}
