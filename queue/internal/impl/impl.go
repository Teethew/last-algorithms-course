package impl

import (
	"github.com/Teethew/last-algorithms-course/list"
	"github.com/Teethew/last-algorithms-course/list/linked/singly"
)

func NewQueueImpl[T list.Type]() *QueueImpl[T] {
	return &QueueImpl[T]{
		list: list.NewSinglyLinkedList[T](),
	}
}

type QueueImpl[T list.Type] struct {
	list list.LinkedList[T, *singly.Node[T]]
}

func (q *QueueImpl[T]) Enqueue(item T) {
	q.list.Add(item)
}

func (q *QueueImpl[T]) Dequeue() (item T) {
	head := q.list.Head()

	if head == nil {
		return item
	}

	result := head.Val
	q.list.Remove(head)
	return result
}

func (q *QueueImpl[T]) Peek() T {
	return q.list.Head().Val
}

func (q *QueueImpl[T]) Length() int {
	return q.list.Length()
}

func (q *QueueImpl[T]) ToString() string {
	return q.list.ToString()
}
