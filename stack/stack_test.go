package stack

import "testing"

func TestStack(t *testing.T) {
	stack := NewLinkedStack[string]()

	stack.Push("habiloaldo")

	if stack.ToString() != "[ habiloaldo ]" {
		t.Fatal()
	}

	stack.Push("jefferson")

	if stack.ToString() != "[ jefferson  habiloaldo ]" {
		t.Fatal()
	}

	if stack.Peek() != "jefferson" {
		t.Fatal()
	}

	_ = stack.Pop()

	if stack.ToString() != "[ habiloaldo ]" {
		t.Fatal()
	}
}
