package list

import (
	"github.com/Teethew/last-algorithms-course/list/internal/linked/doubly"
	"github.com/Teethew/last-algorithms-course/list/internal/linked/singly"
)

func NewSinglyLinkedList() LinkedList[*singly.Node] {
	return singly.NewSinglyLinkedList()
}

func NewDoublyLinkedList() LinkedList[*doubly.Node] {
	return doubly.NewDoublyLinkedList()
}