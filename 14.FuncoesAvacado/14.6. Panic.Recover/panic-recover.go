package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Recuperando a execução")
	}

}
func alunoEstaAprovado(nota1, nota2 float32) bool {
	defer recuperarExecucao() // Descomentar para ver o defer funcionando
	media := (nota1 + nota2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MEDIA É EXATAMENTE 6")
}

func main() {
	fmt.Println("PANIC E RECOVER")

	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("FIM DO PROGRAMA")
	// A função panic interrompe a execução do programa e imprime uma mensagem de erro.

}
