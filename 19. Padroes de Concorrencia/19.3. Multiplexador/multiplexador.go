package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	canal := multiplexar(escrever("Canal 1"),
		escrever("Canal 2"),
		escrever("Canal 3"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}

func multiplexar(canal1, canal2, canal3 <-chan string) <-chan string {

	canalDeSaida := make(chan string)
	go func() {
		for {
			select {
			case msg1 := <-canal1:
				canalDeSaida <- msg1
			case msg2 := <-canal2:
				canalDeSaida <- msg2
			case msg3 := <-canal3:
				canalDeSaida <- msg3
			}
		}
	}()
	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)
	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
		}
	}()
	return canal

	//funcao que gera o canal
}
