package main

import "fmt"

func somaNummeros(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	fmt.Println(texto)
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

// nao pode ter mais de 1 variatico por funçao. o parametro variatico tem que ser o ultimo
func main() {
	fmt.Println("Funções Variáticas")
	// Variática é uma função que pode receber um número variável de parâmetros

	// O operador "..." é usado para indicar que a função pode receber um número variável de argumentos.
	totalDaSoma := somaNummeros(1, 2, 3, 4, 5)
	fmt.Println("Total da soma:", totalDaSoma)

	escrever("Olá", 1, 2, 3, 4, 5)

}
