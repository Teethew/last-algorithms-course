package main

import (
	"fmt"

	"github.com/Teethew/last-algorithms-course/search"
)

func main() {
	arr := []int{3,6,7}

	fmt.Println(search.LinearSearch(arr, 3))
	fmt.Println(search.LinearSearch(arr, 0))
}