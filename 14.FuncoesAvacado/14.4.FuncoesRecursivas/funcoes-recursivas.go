package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-1) + fibonacci(posicao-2)
}

func main() {
	fmt.Println("Recursividade")
	// A recursividade é uma técnica de programação onde uma função chama a si mesma.

	posicao := uint(10)
	fmt.Println("Fibonacci na posição", posicao, "é", fibonacci(posicao))

	posicao2 := uint(10)
	for i := uint(0); i <= posicao2; i++ {
		fmt.Println("Fibonacci na posição", i, "é", fibonacci(uint(posicao2)))
	}

}
