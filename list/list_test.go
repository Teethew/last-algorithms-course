package list

import (
	"testing"

	"github.com/Teethew/last-algorithms-course/list/linked/doubly"
	"github.com/Teethew/last-algorithms-course/list/linked/singly"
)

var singlyLinkedList LinkedList[int, *singly.Node[int]]
var doublyLinkedList LinkedList[int, *doubly.Node[int]]

func setup() {
	singlyLinkedList = NewSinglyLinkedList[int]()
	doublyLinkedList = NewDoublyLinkedList[int]()
}

func TestSinglyLinkedList(t *testing.T) {
	setup()

	var list List[int] = singlyLinkedList

	list.Add(3)
	list.Add(6)
	list.Add(9)

	if list.String() != "[ 3  6  9 ]" {
		t.Fatal()
	}

	singlyLinkedList.Prepend(0)

	if list.String() != "[ 0  3  6  9 ]" {
		t.Fatal()
	}

	result, err := list.Get(3)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if result != 9 {
		t.Fatal()
	}

	err = list.Set(3, 12)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if list.String() != "[ 0  3  6  12 ]" {
		t.Fatal()
	}

	err = list.DeleteAt(3)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if list.String() != "[ 0  3  6 ]" {
		t.Fatal()
	}

	err = list.DeleteAt(1)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if list.String() != "[ 0  6 ]" {
		t.Fatal()
	}

	err = list.AddAt(1, 3)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if list.String() != "[ 0  3  6 ]" {
		t.Fatal()
	}

	err = singlyLinkedList.Remove(singlyLinkedList.Head())

	if err != nil {
		t.Fatalf(err.Error())
	}

	if list.String() != "[ 3  6 ]" {
		t.Fatal()
	}
}

func TestDoublyLinkedList(t *testing.T) {
	setup()

	var list List[int] = doublyLinkedList

	list.Add(3)
	list.Add(6)
	list.Add(9)

	if list.String() != "[ 3  6  9 ]" {
		t.Fatal()
	}

	doublyLinkedList.Prepend(0)

	if list.String() != "[ 0  3  6  9 ]" {
		t.Fatal()
	}

	result, err := list.Get(3)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if result != 9 {
		t.Fatal()
	}

	err = list.Set(3, 12)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if list.String() != "[ 0  3  6  12 ]" {
		t.Fatal()
	}

	err = list.DeleteAt(3)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if list.String() != "[ 0  3  6 ]" {
		t.Fatal()
	}

	err = list.DeleteAt(1)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if list.String() != "[ 0  6 ]" {
		t.Fatal()
	}

	err = list.AddAt(1, 3)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if list.String() != "[ 0  3  6 ]" {
		t.Fatal()
	}

	err = doublyLinkedList.Remove(doublyLinkedList.Head())

	if err != nil {
		t.Fatalf(err.Error())
	}

	if list.String() != "[ 3  6 ]" {
		t.Fatal()
	}
}

func TestArrayList(t *testing.T) {
	arr := NewArrayList[int](2)

	if arr.String() != "[]" {
		t.Fatal()
	}

	arr.Add(1)
	arr.Add(2)
	arr.Add(3)

	if arr.String() != "[ 1  2  3 ]" {
		t.Fatal()
	}

	arr.AddAt(2, 5)

	if arr.String() != "[ 1  2  5  3 ]" {
		t.Fatal()
	}

	arr.DeleteAt(1)

	if arr.String() != "[ 1  5  3 ]" {
		t.Fatal()
	}

	arr.Set(2, 7)

	if arr.String() != "[ 1  5  7 ]" {
		t.Fatal()
	}
}
