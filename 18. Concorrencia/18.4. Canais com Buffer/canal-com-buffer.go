package main

import (
	"fmt"
)

func main() {
	canal := make(chan string, 2) // cria um canal do tipo string com buffer de 2
	canal <- "Hello, World!"      // envia uma mensagem para o canal
	canal <- "Programando em Go!" // envia outra mensagem para o canal

	mensagem := <-canal  // recebe a mensagem do canal
	mensagem2 := <-canal // recebe a mensagem do canal

	fmt.Println(mensagem2) // imprime a mensagem recebida
	fmt.Println(mensagem)  // imprime a mensagem recebida

}
