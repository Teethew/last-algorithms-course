package tree

import (
	"github.com/Teethew/last-algorithms-course/list"
	"github.com/Teethew/last-algorithms-course/queue"
)

type BinaryTreeImpl[T comparable] struct {
	Root *BinaryTreeNode[T]
}

func NewBinaryTree[T comparable]() *BinaryTreeImpl[T] {
	return &BinaryTreeImpl[T]{Root: &BinaryTreeNode[T]{}}
}

func (t *BinaryTreeImpl[T]) PreOrder() []T {
	path := list.NewArrayList[T](10)
	t.preOrderTraversal(t.Root, path)
	return path.ToArray()
}

func (t *BinaryTreeImpl[T]) preOrderTraversal(node *BinaryTreeNode[T], path list.List[T]) {
	if node == nil {
		return
	}

	// pre
	path.Add(node.Val)
	// recursive
	t.preOrderTraversal(node.Left, path)
	t.preOrderTraversal(node.Right, path)
	// post
}

func (t *BinaryTreeImpl[T]) InOrder() []T {
	path := list.NewArrayList[T](10)
	t.inOrderTraversal(t.Root, path)
	return path.ToArray()
}

func (t *BinaryTreeImpl[T]) inOrderTraversal(node *BinaryTreeNode[T], path list.List[T]) {
	if node == nil {
		return
	}

	// pre
	// recursive
	t.inOrderTraversal(node.Left, path)
	path.Add(node.Val)
	t.inOrderTraversal(node.Right, path)
	// post
}

func (t *BinaryTreeImpl[T]) PostOrder() []T {
	path := list.NewArrayList[T](10)
	t.postOrderTraversal(t.Root, path)
	return path.ToArray()
}

func (t *BinaryTreeImpl[T]) postOrderTraversal(node *BinaryTreeNode[T], path list.List[T]) { // corrigir
	if node == nil {
		return
	}

	// pre
	// recursive
	t.postOrderTraversal(node.Left, path)
	t.postOrderTraversal(node.Right, path)
	// post
	path.Add(node.Val)
}

func (t *BinaryTreeImpl[T]) BFS(value T) bool {
	q := queue.NewQueue[*BinaryTreeNode[T]]()
	q.Enqueue(t.Root)

	for q.Length() > 0 {
		if current := q.Dequeue(); current != nil {

			if value == current.Val {
				return true
			}

			if current.Left != nil {
				q.Enqueue(current.Left)
			}
			if current.Right != nil {
				q.Enqueue(current.Right)
			}
		} else {
			break
		}
	}

	return false
}

// func (t *BinaryTreeImpl[T]) AddLeaf(value T) {
// 	node := &BinaryTreeNode[T]{Val: value}

// 	if t.root == nil {
// 		t.root = node
// 		return
// 	}

// }

// func add[T any](n *BinaryTreeNode[T], value T) {

// }
