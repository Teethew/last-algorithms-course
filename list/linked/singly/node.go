package singly

type Type interface {
	any
}

type Node[T Type] struct {
	Val  T
	next *Node[T]
}
