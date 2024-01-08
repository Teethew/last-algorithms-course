package doubly

type Type interface {
	any
}

type Node[T Type] struct {
	Val  T
	next *Node[T]
	prev *Node[T]
}
