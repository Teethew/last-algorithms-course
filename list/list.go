package list

type List interface {
	// Adds new item to the list with the specified value
	Add(value int)

	// Returns the value from the specified index
	Get(index int) (int, error)

	// Updates the item in the specified index with the specified value
	Update(index, value int) error

	// Deletes the item at the specified index
	DeleteAt(index int) error

	// Adds the specified value at the specified index 
	AddAt(index, value int) error
	
	// Returns the length of the list
	Length() int
	
	// Returns a string representation of the list
	ToString() string
}

type LinkedList[N Node] interface {
	// Inherit methods from Interface List
	List
	
	// Adds a new node with the value specified at the start of the LinkedList
	Prepend(value int)
	
	// Returns the head (the first node) of the LinkedList
	Head() N
	
	// Returns the tail (the last node) of the LinkedList
	Tail() N
	
	// Removes the specified node of the LinkedList
	Remove(node N) error
}