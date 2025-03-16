package stack

import "testing"

func TestStack(t *testing.T) {
	stack := NewLinkedStack[string]()

	stack.Push("habiloaldo")

	if stack.String() != "[ habiloaldo ]" {
		t.Fatal()
	}

	stack.Push("jefferson")

	if stack.String() != "[ jefferson  habiloaldo ]" {
		t.Fatal()
	}

	if stack.Peek() != "jefferson" {
		t.Fatal()
	}

	_ = stack.Pop()

	if stack.String() != "[ habiloaldo ]" {
		t.Fatal()
	}
}
