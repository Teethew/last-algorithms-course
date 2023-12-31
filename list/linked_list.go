package list

import "fmt"

type Node struct {
	Val  int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil, length: 0}
}

func (l *LinkedList) Add(value int) {
	node := Node{Val: value, next: nil}

	if l.head == nil {
		l.head = &node
		l.length++
		return
	}

	for n := l.head; ; n = n.next {
		if n.next == nil {
			n.next = &node
			l.length++
			break
		}
	}
}

func (l *LinkedList) Prepend(value int) {
	node := Node{Val: value, next: l.head}
	l.head = &node
}

func (l *LinkedList) Get(index int) int {
	return 0
}

func (l *LinkedList) Update(index, value int) {}

func (l *LinkedList) DeleteAt(index int) {}

func (l *LinkedList) Remove(node *Node) {}

func (l *LinkedList) AddAt(index, value int) {}

func (l *LinkedList) Length() int {
	return l.length
}

func (l *LinkedList) ToString() (str string) {
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
