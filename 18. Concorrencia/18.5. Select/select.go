package main

import (
	"fmt"
	"time"
)

func main() {

	canal1, canal2 := make(chan string), make(chan string) // cria dois canais do tipo string

	go func() {
		for {
			time.Sleep(time.Millisecond * 500) // aguarda 1 segundo antes de enviar a mensagem
			canal1 <- "Canal 1"                // envia uma mensagem para o canal
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2) // aguarda 1 segundo antes de enviar a mensagem
			canal2 <- "Canal 2"         // envia uma mensagem para o canal
		}
	}()

	for {
		select {
		case mensagem1 := <-canal1: // recebe a mensagem do canal 1
			fmt.Println(mensagem1) // imprime a mensagem recebida
		case mensagem2 := <-canal2: // recebe a mensagem do canal 2
			fmt.Println(mensagem2) // imprime a mensagem recebida

		}
	}

}
