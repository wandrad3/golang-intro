package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}
func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoAprovado(nota1, nota2 float32) bool {
	defer fmt.Println("Média calculada")
	fmt.Println("Entrando na função alunoAprovado")
	media := (nota1 + nota2) / 2
	if media >= 7 {

		return true
	}
	return false
}

func main() {
	fmt.Println("Defer")

	//defer funcao1()
	//funcao2()

	fmt.Println(alunoAprovado(7, 8))

}
