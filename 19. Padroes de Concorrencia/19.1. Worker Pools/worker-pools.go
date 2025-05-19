package main

import "fmt"

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-1) + fibonacci(posicao-2)
}

func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func main() {
	tarefas := make(chan int, 20)
	resultados := make(chan int, 20)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 20; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := 0; i < 20; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}
