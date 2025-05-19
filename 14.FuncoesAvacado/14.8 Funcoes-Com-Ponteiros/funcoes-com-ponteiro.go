package main

import "fmt"

func inverterSinal(numero int) int {
	numero = numero * -1
	return numero
}

func inverSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	fmt.Println("Ponteiro em funções")
	numero := 10
	numeroInvertido := inverterSinal(numero)
	fmt.Println("Número original:", numero)
	fmt.Println("Número invertido:", numeroInvertido)

	novoNumero := 20
	fmt.Println("Número original:", novoNumero)
	inverSinalComPonteiro(&novoNumero)
	fmt.Println("Número invertido:", novoNumero)
	// Ponteiro é uma variável que armazena o endereço de outra variável
}
