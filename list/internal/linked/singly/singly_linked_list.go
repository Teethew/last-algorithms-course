package singly

import "fmt"

type SinglyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// List Interface implementation methods

func (l *SinglyLinkedList) Add(value int) {
	node := Node{Val: value, next: nil}

	if l.head == nil {
		l.head = &node
		l.tail = &node
		l.length++
		return
	}

	l.tail.next = &node
	l.tail = &node
	l.length++
}

func (l *SinglyLinkedList) Get(index int) (value int, err error) {

	node, err := l.getNode(index)

	if err != nil {
		return 0, err
	}

	return node.Val, nil
}

func (l *SinglyLinkedList) Update(index, value int) error {

	node, err := l.getNode(index)

	if err != nil {
		return err
	}

	node.Val = value

	return nil
}

func (l *SinglyLinkedList) DeleteAt(index int) error {

	if index == 0 {
		if l.length > 0 {
			l.head = l.head.next
			l.length--
			return nil
		} else {
			return outOfBoundsError(index)
		}
	}

	if index >= l.length {
		return outOfBoundsError(index)
	}

	prevNode, err := l.getNode(index - 1)

	if err != nil {
		return err
	}

	prevNode.next = prevNode.next.next

	return nil
}

func (l *SinglyLinkedList) AddAt(index, value int) error {

	previous, err := l.getNode(index - 1)

	if err != nil {
		return err
	}

	newNode := &Node{Val: value, next: previous.next}

	previous.next = newNode

	return nil
}

func (l *SinglyLinkedList) Length() int {
	return l.length
}

func (l *SinglyLinkedList) ToString() (str string) {
	str = "["

	if l.length == 0 {
		str += "]"
		return
	}

	for n := l.head; ; n = n.next {
		str += fmt.Sprintf(" %d ", n.Val)
		if n.next == nil {
			break
		}
	}

	str += "]"

	return
}

// SinglyLinkedList methods

// Adds a new node with the value specified at the start of the SinglyLinkedList
func (l *SinglyLinkedList) Prepend(value int) {
	node := Node{Val: value, next: l.head}
	l.head = &node
	l.length++
}

// Returns the head (the first node) of the SinglyLinkedList
func (l *SinglyLinkedList) Head() *Node {
	return l.head
}

// Returns the tail (the last node) of the SinglyLinkedList
func (l *SinglyLinkedList) Tail() *Node {
	return l.tail
}

// Removes the specified node of the SinglyLinkedList
func (l *SinglyLinkedList) Remove(node *Node) error {
	var prev *Node

	if node == l.head {
		l.head = l.head.next
		return nil
	}

	for n := l.head; ; n = n.next {
		if n.next == nil {
			return fmt.Errorf("node not found")
		}

		if &n.next == &node {
			prev = n
			break
		}
	}

	prev.next = prev.next.next

	return nil
}

// Private methods

func (l *SinglyLinkedList) getNode(index int) (*Node, error) {
	if index < l.length {
		for i, n := 0, l.head; i < l.length; i, n = i+1, n.next {
			if i == index {
				return n, nil
			}
		}
	}

	return nil, outOfBoundsError(index)
}

func outOfBoundsError(index int) error {
	return fmt.Errorf("index %d out of bounds", index)
}
