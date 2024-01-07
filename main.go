package main

import (
	"fmt"

	"github.com/Teethew/last-algorithms-course/list"
)

func main() {
	arr := list.NewArrayList[int](2)

	fmt.Println(arr.ToString()) // []

	arr.Add(1)
	arr.Add(2)
	arr.Add(3)

	fmt.Println(arr.ToString()) // [ 1  2  3 ]

	arr.AddAt(2, 5)

	fmt.Println(arr.ToString()) // [ 1  2  5  3 ]

	arr.DeleteAt(1)

	fmt.Println(arr.ToString()) // [ 1  5  3 ]

	arr.Set(2, 7)

	fmt.Println(arr.ToString()) // [ 1  5  7 ]
}
