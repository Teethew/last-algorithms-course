package singly

import "fmt"

type SinglyLinkedList[T Type] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func NewLinkedList[T Type]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// List Interface implementation methods

func (l *SinglyLinkedList[T]) Add(value T) {
	node := &Node[T]{Val: value, next: nil}

	if l.head == nil {
		l.head = node
		l.tail = node
		l.length++
		return
	}

	l.tail.next = node
	l.tail = node
	l.length++
}

func (l *SinglyLinkedList[T]) Get(index int) (value T, err error) {

	node, err := l.getNode(index)

	if err != nil {
		return node.Val, err
	}

	return node.Val, nil
}

func (l *SinglyLinkedList[T]) Set(index int, value T) error {

	node, err := l.getNode(index)

	if err != nil {
		return err
	}

	node.Val = value

	return nil
}

func (l *SinglyLinkedList[T]) DeleteAt(index int) error {

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
	l.length--

	return nil
}

func (l *SinglyLinkedList[T]) AddAt(index int, value T) error {

	previous, err := l.getNode(index - 1)

	if err != nil {
		return err
	}

	newNode := &Node[T]{Val: value, next: previous.next}

	previous.next = newNode
	l.length++

	return nil
}

func (l *SinglyLinkedList[T]) Length() int {
	return l.length
}

func (l *SinglyLinkedList[T]) ToString() (str string) {
	str = "["

	if l.length == 0 {
		str += "]"
		return
	}

	for n := l.head; ; n = n.next {
		str += fmt.Sprintf(" %v ", n.Val)
		if n.next == nil {
			break
		}
	}

	str += "]"

	return
}

// Linked List Interface implementation methods

func (l *SinglyLinkedList[T]) Prepend(value T) {
	node := Node[T]{Val: value, next: l.head}
	l.head = &node
	l.length++
}

func (l *SinglyLinkedList[T]) Head() *Node[T] {
	return l.head
}

func (l *SinglyLinkedList[T]) Tail() *Node[T] {
	return l.tail
}

func (l *SinglyLinkedList[T]) Remove(node *Node[T]) error {
	var prev *Node[T]

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
	l.length--

	return nil
}

// Private methods

func (l *SinglyLinkedList[T]) getNode(index int) (*Node[T], error) {
	if index < l.length {
		for i, n := 0, l.head; i < l.length; i, n = i+1, n.next {
			if i == index {
				return n, nil
			}
		}
	}

	return &Node[T]{}, outOfBoundsError(index)
}

func outOfBoundsError(index int) error {
	return fmt.Errorf("index %d out of bounds", index)
}
