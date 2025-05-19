package main

import (
	"fmt"
	"sync"
	"time"
)

// Go routines are lightweight threads managed by the Go runtime.
// They are used to perform concurrent tasks in Go programs.
// Go routines are created using the `go` keyword followed by a function call.

func main() {
	var waitGroup sync.WaitGroup
	// WaitGroup is used to wait for a collection of goroutines to finish.

	waitGroup.Add(2) // Adiciona 1 ao contador do WaitGroup

	go func() {
		escrever("Hello, World!") //goroutine
		waitGroup.Done()          // Decrementa o contador do WaitGroup

	}()
	// O programa continua executando enquanto a goroutine está em execução
	// O programa principal não espera a goroutine terminar
	go func() {
		escrever("Programando em Go!")
		//goroutine
		waitGroup.Done() // Decrementa o contador do WaitGroup

	}()
	waitGroup.Wait() // Espera todas as goroutines terminarem

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(1 * time.Second)
	}
}
