package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello, World!")
	auxiliar.Escrever()

	err := checkmail.ValidateFormat("wesleyandrade298@gmail.com")
	fmt.Println(err)
	err2 := checkmail.ValidateFormat("123")
	fmt.Println(err2)
}
