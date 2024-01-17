package impl

import (
	"github.com/Teethew/last-algorithms-course/list"
	"github.com/Teethew/last-algorithms-course/list/linked/singly"
)

func NewLinkedStack[T list.Type]() *LinkedStack[T] {
	return &LinkedStack[T]{
		list: list.NewSinglyLinkedList[T](),
	}
}

type LinkedStack[T list.Type] struct {
	list list.LinkedList[T, *singly.Node[T]]
}

func (s *LinkedStack[T]) Push(item T) {
	s.list.Prepend(item)
}

func (s *LinkedStack[T]) Pop() T {
	head := s.list.Head()
	s.list.Remove(head)
	return head.Val
}

func (s *LinkedStack[T]) Peek() T {
	return s.list.Head().Val
}

func (s *LinkedStack[T]) ToString() string {
	return s.list.ToString()
}

func (s *LinkedStack[T]) Length() int {
	return s.list.Length()
}
