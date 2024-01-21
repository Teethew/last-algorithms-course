package doubly

type Type interface {
	comparable
}

type Node[T Type] struct {
	Val  T
	next *Node[T]
	prev *Node[T]
}
