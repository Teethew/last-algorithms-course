package queue

import (
	"github.com/Teethew/last-algorithms-course/list"
	"github.com/Teethew/last-algorithms-course/queue/internal/impl"
)

type Queue[T list.Type] interface {
	Enqueue(item T)
	Deque() T
	Peek() T
	ToString() string
}

func NewQueue[T list.Type]() Queue[T] {
	return impl.NewQueueImpl[T]()
}
