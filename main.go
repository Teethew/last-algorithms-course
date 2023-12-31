package main

import (
	"fmt"

	"github.com/Teethew/last-algorithms-course/list"
)

func main() {
	var listaLigada = list.NewLinkedList()
	var lista list.List = listaLigada

	lista.Add(3)
	lista.Add(6)
	lista.Add(9)

	fmt.Printf("%s\n", lista.ToString()) // [ 3 6 9 ]
	
	listaLigada.Prepend(0)
	
	fmt.Printf("%s\n", lista.ToString()) // [ 0 3 6 9 ]
}
