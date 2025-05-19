package main

import (
	"fmt"
	"time"
)

// Go routines are lightweight threads managed by the Go runtime.
// They are used to perform concurrent tasks in Go programs.
// Go routines are created using the `go` keyword followed by a function call.

func main() {

	go escrever("Hello, World!") //goroutine
	// O programa continua executando enquanto a goroutine está em execução
	// O programa principal não espera a goroutine terminar

	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(1 * time.Second)
	}
}
