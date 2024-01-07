package doubly

type Type interface {
	int | string | struct{}
}

type Node[T Type] struct {
	Val  T
	next *Node[T]
	prev *Node[T]
}
