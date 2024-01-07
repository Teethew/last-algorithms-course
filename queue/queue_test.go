package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue[string]()

	queue.Enqueue("habiloaldo")

	if queue.ToString() != "[ habiloaldo ]" {
		t.Fatal()
	}

	queue.Enqueue("jefferson")

	if queue.ToString() != "[ habiloaldo  jefferson ]" {
		t.Fatal()
	}

	if queue.Peek() != "habiloaldo" {
		t.Fatal()
	}

	_ = queue.Deque()

	if queue.ToString() != "[ jefferson ]" {
		t.Fatal()
	}
}
