package list

import (
	"github.com/Teethew/last-algorithms-course/list/array/impl"
	"github.com/Teethew/last-algorithms-course/list/linked/doubly"
	"github.com/Teethew/last-algorithms-course/list/linked/singly"
)

func NewArrayList[T Type](initialCapacity int) ArrayList[T] {
	return impl.NewArrayList[T](initialCapacity)
}

func NewSinglyLinkedList[T Type]() LinkedList[T, *singly.Node[T]] {
	return singly.NewLinkedList[T]()
}

func NewDoublyLinkedList[T Type]() LinkedList[T, *doubly.Node[T]] {
	return doubly.NewLinkedList[T]()
}
