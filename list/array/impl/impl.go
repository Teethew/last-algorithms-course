package impl

import "fmt"

type Type interface {
	int | string | struct{}
}

type ArrayList[T Type] struct {
	array    []T
	length   int
	capacity int
}

func NewArrayList[T Type](initialCapacity int) *ArrayList[T] {
	return &ArrayList[T]{
		array:    make([]T, initialCapacity),
		length:   0,
		capacity: initialCapacity,
	}
}

func (a *ArrayList[T]) Add(value T) {
	if a.length < a.capacity {
		index := a.length
		a.array[index] = value
		a.length++
		return
	}

	a.array = a.doubleCapacity()
	a.capacity *= 2
	index := a.length
	a.array[index] = value
	a.length++
}

func (a *ArrayList[T]) AddAt(index int, value T) error {
	var new []T

	if index >= a.length {
		return outOfBoundsError(index)

	} else if a.length == a.capacity {
		new = a.doubleCapacity()

	} else {
		new = make([]T, a.capacity)
	}

	a.length++

	for i, j := 0, 0; i < a.length; i, j = i+1, j+1 {
		if i == index {
			j--
			new[i] = value
			continue
		}

		new[i] = a.array[j]
	}

	a.array = new

	return nil
}

func (a *ArrayList[T]) DeleteAt(index int) error {
	if index >= a.length {
		return outOfBoundsError(index)
	}

	new := make([]T, a.capacity)

	a.length--

	for i, j := 0, 0; i <= a.length; i, j = i+1, j+1 {
		if i == index {
			i++
			new[j] = a.array[i]
			continue
		}

		new[j] = a.array[i]
	}

	a.array = new

	return nil
}

func (a *ArrayList[T]) Get(index int) (T, error) {
	if index >= a.length {
		return a.array[0], outOfBoundsError(index)
	}

	return a.array[index], nil
}

func (a *ArrayList[T]) Length() int {
	return a.length
}

func (a *ArrayList[T]) Set(index int, value T) error {
	if index >= a.length {
		return outOfBoundsError(index)
	}

	a.array[index] = value
	return nil
}

func (a *ArrayList[T]) ToString() (str string) {
	str = "["

	if a.length == 0 {
		str += "]"
		str = fmt.Sprintf(
			`{
  length: %d
  values: %v
  capacity: %d
}`, a.length, str, a.capacity)
		return
	}

	for i := 0; i < a.length; i++ {
		str += fmt.Sprintf(" %v ", a.array[i])
	}

	str += "]"

	str = fmt.Sprintf(
		`{
  length: %d
  values: %v
  capacity: %d
}`, a.length, str, a.capacity)

	return
}

func (a *ArrayList[T]) doubleCapacity() []T {
	capacity := a.capacity * 2

	new := make([]T, capacity)

	copy(new, a.array)

	return new
}

func outOfBoundsError(index int) error {
	return fmt.Errorf("index %d out of bounds", index)
}
