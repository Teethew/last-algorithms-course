package tree

import (
	"slices"
	"testing"
)

func setupBinaryTree() *BinaryTreeImpl[int] {
	b := NewBinaryTree[int]()
	b.Root = &BinaryTreeNode[int]{
		Val: 1,
		Left: &BinaryTreeNode[int]{
			Val: 2,
		},
		Right: &BinaryTreeNode[int]{
			Val: 3,
		},
	}

	return b
}

func TestBSTFind(t *testing.T) {
	b := &BST{t: &BinaryTreeImpl[int]{
		Root: &BinaryTreeNode[int]{
			Val:   2,
			Left:  &BinaryTreeNode[int]{Val: 1},
			Right: &BinaryTreeNode[int]{Val: 3},
		},
	}}

	if !b.Find(2) {
		t.Error("expected to find value 2")
	}

	if b.Find(54) {
		t.Error("not expected to find value 54")
	}
}

func TestPreOrderTraversal(t *testing.T) {
	b := setupBinaryTree()

	if slices.Compare(b.PreOrder(), []int{1, 2, 3}) != 0 {
		t.Error("unexpected sequence at pre order traversal")
	}
}

func TestInOrderTraversal(t *testing.T) {
	b := setupBinaryTree()

	if slices.Compare(b.InOrder(), []int{2, 1, 3}) != 0 {
		t.Error("unexpected sequence at in order traversal")
	}
}

func TestPostOrderTraversal(t *testing.T) {
	b := setupBinaryTree()

	if slices.Compare(b.PostOrder(), []int{2, 3, 1}) != 0 {
		t.Error("unexpected sequence at post order traversal")
	}
}

func TestLeverOrderTraversal(t *testing.T) {
	// implement
}

func TestBreadthFirstSearch(t *testing.T) {
	b := setupBinaryTree()

	if !b.BFS(3) {
		t.Error("expected to find value 3")
	}

	if b.BFS(54) {
		t.Error("unexpected value 54")
	}
}

func TestBST(t *testing.T) {
	b := NewBST(nil)
	b.Insert(54)
	b.Insert(17)
	b.Insert(5)
	b.Insert(128)
	b.Insert(65)
	b.Insert(256)
	b.Insert(39)

	expected := []int{5, 39, 17, 65, 256, 128, 54}
	result := b.PostOrder()

	if slices.Compare(expected, result) != 0 {
		t.Errorf("insert: expected: %v, instead got %v", expected, result)
	}

	if !b.IsValid() {
		t.Error("insert: expected newly created BST to be valid")
	}

	if !slices.IsSorted(b.InOrder()) {
		t.Error("expected BST in order traversal to present elements in increasing order")
	}

	b = NewBST(
		&BinaryTreeImpl[int]{
			Root: &BinaryTreeNode[int]{
				Val:  5,
				Left: &BinaryTreeNode[int]{Val: 1},
				Right: &BinaryTreeNode[int]{
					Val:   4,
					Left:  &BinaryTreeNode[int]{Val: 3},
					Right: &BinaryTreeNode[int]{Val: 6},
				},
			},
		},
	)

	if b.IsValid() {
		t.Error("invalid BST considered valid")
	}
}

func TestMinHeap(t *testing.T) {
	h := NewMinHeap()
	h.Insert(7)
	h.Insert(4)
	h.Insert(3)
	h.Insert(1)

	if h.Length() != 4 {
		t.Errorf("expected length to be 4, instead got: %d", h.Length())
	}

	if expected := []int{1, 3, 4, 7}; slices.Compare[[]int, int](h.ToArray(), expected) != 0 {
		t.Errorf("expected %v, instead got %v", expected, h.ToArray())
	}

	if value, err := h.Pop(); err != nil || value != 1 {
		t.Errorf("expected to pop 1 without errors, instead got \nerr: %v\nvalue: %v", err, value)
	}

	if h.GetMin() != 3 {
		t.Errorf("expected GetMin to return 3, instead returned %d", h.GetMin())
	}

	h.Pop()
	h.Pop()
	h.Pop()

	if _, err := h.Pop(); err == nil {
		t.Errorf("expected error on popping from empty heap, instead got nil")
	}
}
