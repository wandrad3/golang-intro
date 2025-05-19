package main

import (
	"errors"
	"fmt"
)

func main() {

	var numero int16 = 100
	fmt.Println(numero)

	var numero2 uint = 1000 //unsigned int aceita somente numeros positivos

	fmt.Println(numero2)

	//alias
	// int32 = rune
	var numero3 rune = 123456 //rune é um alias para int32
	fmt.Println(numero3)

	//BYTE = uint8
	var numero4 byte = 123 //byte é um alias para uint8
	fmt.Println(numero4)

	//float
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000.456789
	fmt.Println(numeroReal2)

	numeroReal3 := 123.456 // o compilador identifica o tipo
	fmt.Println(numeroReal3)

	//string
	var str1 string = "Texto"
	fmt.Println(str1)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'A' // char é um alias para rune
	fmt.Println(char)

	//valor inicial

	var text string

	fmt.Println(text) // string vazia

	var inteiro int

	fmt.Println(inteiro) // int 0

	//booleano

	var booleano1 bool
	fmt.Println(booleano1) // false

	var booleano2 bool = true
	fmt.Println(booleano2) // true

	var erro error
	fmt.Println(erro) // nil
	// nil é o valor zero para ponteiros, slices, maps, channels e interfaces

	var erro1 error = errors.New("Erro Interno")
	fmt.Println(erro1)
	// erro1 é um erro com mensagem personalizada

}
