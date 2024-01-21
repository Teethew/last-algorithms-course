package tree

type TreeNode[T comparable] struct {
	Val  T
	Children []*TreeNode[T]
}

type BinaryTreeNode[T comparable] struct {
	Val   T
	Left  *BinaryTreeNode[T]
	Right *BinaryTreeNode[T]
}
