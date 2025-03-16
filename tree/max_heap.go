package tree

import "fmt"

type MaxHeap struct {
	capacity int
	length   int
	values   []int
}

func NewMaxHeap() *MaxHeap {
	initialCapacity := 10

	s := make([]int, initialCapacity)

	return &MaxHeap{
		capacity: initialCapacity,
		length:   0,
		values:   s,
	}
}

func (h *MaxHeap) Insert(n int) {
	if h.length == h.capacity {
		h.doubleCapacity()
	}
	h.values[h.length] = n
	h.heapifyUp(h.length)
	h.length++
}

func (h *MaxHeap) Delete(n int) (err error) {
	if h.length == 0 {
		err = fmt.Errorf("runtime error: (*MaxHeap).Delete() called on empty heap")
		return
	}

	h.length--

	if h.length == 0 {
		return
	}

	h.values[0] = h.values[h.length]
	h.heapifyDown(0)
	return
}

func (h *MaxHeap) Pop() (v int, err error) {
	if h.length == 0 {
		return v, fmt.Errorf("runtime error: (*MaxHeap).Pop() called on empty heap")
	}

	v = h.values[0]

	return v, h.Delete(0)
}

func (h *MaxHeap) GetMax() (v int) {
	if h.length > 0 {
		v = h.values[0]
	}
	return
}

func (h *MaxHeap) Length() int {
	return h.length
}

func (h *MaxHeap) ToArray() []int {
	return h.values[:h.length]
}

func (h *MaxHeap) heapifyUp(i int) {
	if i == 0 {
		return
	}

	parent := parent(i)
	parentValue := h.values[parent]
	value := h.values[i]

	if parentValue < value {
		h.values[i] = parentValue
		h.values[parent] = value
		h.heapifyUp(parent)
	}
}

func (h *MaxHeap) heapifyDown(i int) {
	if i >= h.length {
		return
	}

	left := left(i)
	right := right(i)

	if left >= h.length || right >= h.length {
		return
	}

	leftValue := h.values[left]
	rightValue := h.values[right]
	currentValue := h.values[i]

	if leftValue < rightValue && currentValue < rightValue {
		h.values[i] = rightValue
		h.values[right] = currentValue
		h.heapifyDown(right)

	} else if rightValue < leftValue && currentValue < leftValue {
		h.values[i] = leftValue
		h.values[left] = currentValue
		h.heapifyDown(left)

	}
}

func (h *MaxHeap) doubleCapacity() {
	s := make([]int, h.capacity*2)
	copy(s, h.values)
	h.values = s
	h.capacity *= 2
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}