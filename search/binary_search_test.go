package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 6, 7, 9, 12}

	if BinarySearch(arr, 3) != 2 {
		t.Fatalf("Value 3 not found at index 2 as expected.")
	}

	if BinarySearch(arr, 9) != 5 {
		t.Fatalf("Value 9 not found at index 5 as expected.")
	}

	
	if BinarySearch(arr, 0) != -1 {
		t.Fatalf("Value 0 found at index %d. Expected -1.", BinarySearch(arr, 0))
	}
}