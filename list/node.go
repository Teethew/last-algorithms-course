package list

import (
	"github.com/Teethew/last-algorithms-course/list/linked/doubly"
	"github.com/Teethew/last-algorithms-course/list/linked/singly"
)

type Type interface {
	comparable
}

type Node[T Type] interface {
	*singly.Node[T] | *doubly.Node[T]
}
