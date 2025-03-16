package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue[string]()

	queue.Enqueue("habiloaldo")

	if queue.String() != "[ habiloaldo ]" {
		t.Fatal()
	}

	queue.Enqueue("jefferson")

	if queue.String() != "[ habiloaldo  jefferson ]" {
		t.Fatal()
	}

	if queue.Peek() != "habiloaldo" {
		t.Fatal()
	}

	_ = queue.Dequeue()

	if queue.String() != "[ jefferson ]" {
		t.Fatal()
	}
}
