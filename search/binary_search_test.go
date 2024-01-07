package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{-1, 0, 3, 5, 9, 12}

	if BinarySearch(arr, 9) != 4 {
		t.Fatalf("Value 9 not found. Expected %d. Got %d", 4, BinarySearch(arr, 9))
	}

	if BinarySearch(arr, 2) != -1 {
		t.Fatalf("Value 0 found at index %d. Expected %d.", BinarySearch(arr, 0), -1)
	}
}
