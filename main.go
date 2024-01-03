package main

import (
	"fmt"
	"log"

	"github.com/Teethew/last-algorithms-course/list"
)

func main() {
	// TODO: move to unit tests

	var listaLigada = list.NewSinglyLinkedList()
	var lista list.List = listaLigada

	lista.Add(3)
	lista.Add(6)
	lista.Add(9)

	fmt.Printf("%s\n", lista.ToString()) // [ 3 6 9 ]

	listaLigada.Prepend(0)

	fmt.Printf("%s\n", lista.ToString()) // [ 0 3 6 9 ]

	result, err := lista.Get(3)

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("%d\n", result) // 9

	err = lista.Update(3, 12)

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("%s\n", lista.ToString()) // [ 0 3 6 12 ]

	err = lista.DeleteAt(3)

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("%s\n", lista.ToString()) // [ 0 3 6 ]

	err = lista.DeleteAt(1)

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("%s\n", lista.ToString()) // [ 0 6 ]

	err = lista.AddAt(1, 3)

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("%s\n", lista.ToString()) // [ 0 3 6 ]

	err = listaLigada.Remove(listaLigada.Head())

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("%s\n", lista.ToString()) // [ 3 6 ]
}
