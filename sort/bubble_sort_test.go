package sort

import (
	"slices"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	input := BigExample
	expected := make([]int, len(input))
	_ = copy(expected, input)
	slices.Sort(expected)

	BubbleSort(input)

	result := input

	if !slices.Equal[[]int, int](result, expected) {
		t.Fatalf("Expected: %v. Got %v", expected, result)
	}

}
