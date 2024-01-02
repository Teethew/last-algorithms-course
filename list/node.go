package list

import (
	"github.com/Teethew/last-algorithms-course/list/internal/linked/doubly"
	"github.com/Teethew/last-algorithms-course/list/internal/linked/singly"
)

type Node interface {
	*singly.Node | *doubly.Node
}