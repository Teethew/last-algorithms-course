package main

import (
	"fmt"

	"github.com/Teethew/last-algorithms-course/tree"
)

func main() {
	h := tree.NewMaxHeap()
	h.Insert(7)
	h.Insert(4)
	h.Insert(3)
	h.Insert(1)
	fmt.Println(h)
}
