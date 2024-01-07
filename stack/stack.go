package stack

import (
	"github.com/Teethew/last-algorithms-course/list"
	"github.com/Teethew/last-algorithms-course/stack/internal/impl"
)

type Stack[T list.Type] interface {
	Push(item T)
	Pop() T
	Peek() T
	ToString() string
}

func NewLinkedStack[T list.Type]() Stack[T] {
	return impl.NewLinkedStack[T]()
}
