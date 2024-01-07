package list

type LinkedList[T Type, N Node[T]] interface {
	// Inherit methods from Interface List
	List[T]

	// Adds a new node with the value specified at the start of the LinkedList
	Prepend(value T)

	// Returns the head (the first node) of the LinkedList
	Head() N

	// Returns the tail (the last node) of the LinkedList
	Tail() N

	// Removes the specified node of the LinkedList
	Remove(node N) error
}
