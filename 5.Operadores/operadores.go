package main

import (
	"fmt"
)

func main() {

	// Operadores Aritméticos
	soma := 1 + 2
	subtracao := 3 - 1
	multiplicacao := 2 * 3
	divisao := 10 / 2
	resto := 10 % 3

	fmt.Println("Soma:", soma)
	fmt.Println("Subtração:", subtracao)
	fmt.Println("Multiplicação:", multiplicacao)
	fmt.Println("Divisão:", divisao)
	fmt.Println("Resto da Divisão:", resto)

	// Operadores Relacionais

	//Atribuição
	var a int = 10
	var b int = 20
	maior := a > b
	menor := a < b

	maiorIgual := a >= b
	menorIgual := a <= b
	igual := a == b
	diferente := a != b
	fmt.Println("Maior:", maior)
	fmt.Println("Menor:", menor)
	fmt.Println("Maior ou Igual:", maiorIgual)
	fmt.Println("Menor ou Igual:", menorIgual)
	fmt.Println("Igual:", igual)
	fmt.Println("Diferente:", diferente)

	println("---------------------")
	//operadores lógicos
	println((a > b) && (a < b)) // E
	println((a > b) || (a < b)) // OU
	println(!(a > b))           // NÃO

	// Operadores unarios
	numero := 10
	numero++
	fmt.Println("Número incrementado:", numero)
	numero -= 15 // decrementa 15
	fmt.Println("Número decrementado:", numero)
	numero += 5 // incrementa 5
	fmt.Println("Número incrementado:", numero)

	//ternario

	// Go não tem operador ternário, mas podemos simular com uma função
	// resultado := func(a, b int) int {
	// 	if a > b {
	// 		return a
	// 	}

}
