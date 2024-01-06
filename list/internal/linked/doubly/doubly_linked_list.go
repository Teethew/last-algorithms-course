package doubly

import "fmt"

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

// List Interface implementation methods

func (l *DoublyLinkedList) Add(value int) {
	node := &Node{
		Val:  value,
		next: nil,
		prev: nil,
	}

	if l.head == nil {
		l.head = node
		l.tail = node
		l.length++
		return
	}

	node.prev = l.tail
	l.tail.next = node
	l.tail = node
	l.length++
}

func (l *DoublyLinkedList) Get(index int) (int, error) {

	node, err := l.getNode(index)

	if err != nil {
		return 0, err
	}

	return node.Val, nil
}

func (l *DoublyLinkedList) Update(index, value int) error {

	node, err := l.getNode(index)

	if err != nil {
		return err
	}

	node.Val = value

	return nil
}

func (l *DoublyLinkedList) DeleteAt(index int) error {

	if index == 0 {
		if l.length > 0 {
			l.head.next.prev = nil
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

	node, err := l.getNode(index)

	if err != nil {
		return err
	}

	if node == l.head {
		node.next.prev = nil
		l.head = node.next
		l.length--
		return nil
	}

	if node == l.tail {
		node.prev.next = nil
		l.tail = node.prev
		l.length--
		return nil
	}

	node.prev.next = node.next
	node.next.prev = node.prev
	l.length--

	return nil
}

func (l *DoublyLinkedList) AddAt(index, value int) error {

	if index == 0 {
		l.Prepend(value)
		return nil
	}

	node, err := l.getNode(index)

	if err != nil {
		return err
	}

	newNode := &Node{Val: value, next: node, prev: node.prev}

	node.prev.next = newNode
	node.prev = newNode
	l.length++

	return nil
}

func (l *DoublyLinkedList) Length() int {
	return l.length
}

func (l *DoublyLinkedList) ToString() (str string) {
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

// Linked List Interface implementation methods

func (l *DoublyLinkedList) Prepend(value int) {
	node := Node{
		Val:  value,
		next: l.head,
		prev: nil,
	}
	l.head.prev = &node
	l.head = &node
	l.length++
}

func (l *DoublyLinkedList) Head() *Node {
	return l.head
}

func (l *DoublyLinkedList) Tail() *Node {
	return l.tail
}

func (l *DoublyLinkedList) Remove(node *Node) error {
	var prev *Node

	if node == l.head {
		l.head = l.head.next
		l.head.prev = nil
		l.length--
		return nil
	}

	for n := l.head; ; n = n.next {
		if n.next == nil {
			return fmt.Errorf("node not found")
		}

		if &n == &node {
			prev = n.prev
			break
		}
	}

	node.next.prev = prev
	prev.next = node.next
	l.length--

	return nil
}

// Private methods

func (l *DoublyLinkedList) getNode(index int) (*Node, error) {
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
