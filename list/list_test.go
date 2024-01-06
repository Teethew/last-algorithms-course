package list

import (
	"testing"

	"github.com/Teethew/last-algorithms-course/list/internal/linked/doubly"
	"github.com/Teethew/last-algorithms-course/list/internal/linked/singly"
)

var singlyLinked LinkedList[*singly.Node]
var doublyLinked LinkedList[*doubly.Node]

func setup() {
	singlyLinked = NewSinglyLinkedList()
	doublyLinked = NewDoublyLinkedList()
}

func TestSinglyLinkedList(t *testing.T) {
	setup()

	var list List = singlyLinked

	list.Add(3)
	list.Add(6)
	list.Add(9)

	if list.ToString() != "[ 3  6  9 ]" {
		t.Fatal()
	}

	singlyLinked.Prepend(0)

	if list.ToString() != "[ 0  3  6  9 ]" {
		t.Fatal()
	}

	result, err := list.Get(3)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if result != 9 {
		t.Fatal()
	}

	err = list.Update(3, 12)

	if err != nil {
		t.Fatalf(err.Error())
	}
	
	if list.ToString() != "[ 0  3  6  12 ]" {
		t.Fatal()
	}

	err = list.DeleteAt(3)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if list.ToString() != "[ 0  3  6 ]" {
		t.Fatal()
	}

	err = list.DeleteAt(1)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if list.ToString() != "[ 0  6 ]" {
		t.Fatal()
	}

	err = list.AddAt(1, 3)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if list.ToString() != "[ 0  3  6 ]" {
		t.Fatal()
	}

	err = singlyLinked.Remove(singlyLinked.Head())

	if err != nil {
		t.Fatalf(err.Error())
	}

	if list.ToString() != "[ 3  6 ]" {
		t.Fatal()
	}
}

func TestDoublyLinkedList(t *testing.T) {
	setup()

	var list List = doublyLinked

	list.Add(3)
	list.Add(6)
	list.Add(9)

	if list.ToString() != "[ 3  6  9 ]" {
		t.Fatal()
	}

	doublyLinked.Prepend(0)

	if list.ToString() != "[ 0  3  6  9 ]" {
		t.Fatal()
	}

	result, err := list.Get(3)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if result != 9 {
		t.Fatal()
	}

	err = list.Update(3, 12)

	if err != nil {
		t.Fatalf(err.Error())
	}
	
	if list.ToString() != "[ 0  3  6  12 ]" {
		t.Fatal()
	}

	err = list.DeleteAt(3)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if list.ToString() != "[ 0  3  6 ]" {
		t.Fatal()
	}

	err = list.DeleteAt(1)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if list.ToString() != "[ 0  6 ]" {
		t.Fatal()
	}

	err = list.AddAt(1, 3)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if list.ToString() != "[ 0  3  6 ]" {
		t.Fatal()
	}

	err = doublyLinked.Remove(doublyLinked.Head())

	if err != nil {
		t.Fatalf(err.Error())
	}

	if list.ToString() != "[ 3  6 ]" {
		t.Fatal()
	}
}