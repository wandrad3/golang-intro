package main

import (
	"fmt"
	"math"
)

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	// Generica
	// Uma função generica é uma função que pode receber qualquer tipo de dado como parametro.
	// O tipo de dado é definido como interface{}.
	// A interface{} é uma interface vazia, ou seja, ela pode receber qualquer tipo de dado.
	// A interface{} é o mesmo que o tipo any.

	generica(10)
	generica("Texto")
	generica(true)
	generica(10.5)
	generica(math.Pi)
	generica([]int{1, 2, 3})
	generica(map[string]int{"chave": 1})
}
