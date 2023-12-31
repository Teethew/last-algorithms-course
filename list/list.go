package list

type List interface {
	Add(value int)
	Get(index int) int
	Update(index, value int)
	DeleteAt(index int)
	AddAt(index, value int)
	ToString() string
	Length() int
}