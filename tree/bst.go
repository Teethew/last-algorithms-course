package tree

import (
	"slices"
)

type BST struct {
	t *BinaryTreeImpl[int]
}

func NewBST(tree *BinaryTreeImpl[int]) *BST {
	if tree == nil {
		tree = &BinaryTreeImpl[int]{
			Root: nil,
		}
	}

	return &BST{t: tree}
}

func (b *BST) Insert(value int) {
	if b.t.Root == nil {
		b.t.Root = &BinaryTreeNode[int]{
			Val: value,
		}

		return
	}

	insert(b.t.Root, value)
}

func insert(node *BinaryTreeNode[int], value int) {
	if value <= node.Val {
		if node.Left == nil {
			n := &BinaryTreeNode[int]{
				Val: value,
			}

			node.Left = n
			return
		}

		insert(node.Left, value)
	} else if value > node.Val {
		if node.Right == nil {
			n := &BinaryTreeNode[int]{
				Val: value,
			}

			node.Right = n
			return
		}

		insert(node.Right, value)
	}
}

func (b *BST) IsValid() bool {
	slice := b.InOrder()
	return slices.IsSorted(slice)
}

func (b *BST) Find(value int) bool {
	return find(b.t.Root, value)
}

func find(node *BinaryTreeNode[int], value int) bool {
	if node == nil {
		return false
	}

	if node.Val == value {
		return true
	}

	if node.Val > value {
		return find(node.Left, value)
	} else {
		return find(node.Right, value)
	}
}

func (b *BST) PostOrder() []int {
	return b.t.PostOrder()
}

func (b *BST) PreOrder() []int {
	return b.t.PreOrder()
}

func (b *BST) InOrder() []int {
	return b.t.InOrder()
}
