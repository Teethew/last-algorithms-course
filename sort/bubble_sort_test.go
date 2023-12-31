package sort

import (
	"slices"
	"testing"
)

type Array struct {
	values []int
}

func TestBubbleSort(t *testing.T) {
	input := BigExample
	expected := make([]int, len(input))
	_ = copy(expected, input)
	slices.Sort(expected)

	BubbleSort(input)

	result := &Array{values: input}

	if !result.equals(expected) {
		t.Fatalf("Expected: %d. Got %d", expected, result.values)
	}

}

func (a *Array) equals(a2 []int) bool {
	for i := 0; i < len(a.values); i++ {
		if a.values[i] != a2[i] {
			return false
		}
	}
	return true
}