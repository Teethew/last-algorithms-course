package list

type List interface {
	Add(value int)
	Get(index int) (int, error)
	Update(index, value int) error
	DeleteAt(index int) error
	AddAt(index, value int) error
	Length() int
	ToString() string
}

type LinkedList[N Node] interface {
	List
	Prepend(value int)
	Head() N
	Tail() N
	Remove(node N) error
}