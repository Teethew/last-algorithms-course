package search

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{3, 6, 7}

	if LinearSearch(arr, 3) != 0 {
		t.Fatalf("Value 3 not found at index 0 as expected.")
	}

	if LinearSearch(arr, 0) != -1 {
		t.Fatalf("Value 0 found at index 0. Expected -1.")
	}
}
