package main

import "fmt"

var n int

func init() {
	fmt.Println("Funcao init")
	n = 10
	// A função init é executada antes da função main
	// Ela é usada para inicializar variáveis ou executar código que deve ser executado antes do programa principal
	// Ela não pode receber parâmetros e não pode retornar valores
	// Ela é executada automaticamente pelo Go
	// Não é necessário chamá-la
}

func main() {
	fmt.Println("Funcao main")
	fmt.Println("Valor de n:", n)

}
