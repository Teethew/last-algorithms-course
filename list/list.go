package list

type List[T Type] interface {
	// Adds new item to the list with the specified value
	Add(value T)

	// Returns the value from the specified index
	Get(index int) (T, error)

	// Updates the item in the specified index with the specified value
	Set(index int, value T) error

	// Deletes the item at the specified index
	DeleteAt(index int) error

	// Adds the specified value at the specified index
	AddAt(index int, value T) error

	// Returns the length of the list
	Length() int

	// Returns a string representation of the list
	ToString() string

	// Returns an slice representation of the list
	ToArray() []T
}
