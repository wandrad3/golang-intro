package main

import (
	"fmt"
	"time"
)


func main() {

	canal := escrever("Hello, World!") //goroutine

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}




func escrever(texto string) <-chan {
	canal := make(chan string)
	go func() {
		for {
			canal <- Sprintf("Valor recebido %s", texto)
			time.Sleep(500 * time.Millisecond)
		}
	}()
	return canal
}
