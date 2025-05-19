package main

import (
	"fmt"
	"time"
)

func main() {
	// um canal é uma estrutura de dados que permite a comunicação entre goroutines
	// canais são usados para enviar e receber dados entre goroutines

	canal := make(chan string)          // cria um canal do tipo string
	go escrever("Hello, World!", canal) // inicia uma goroutine que escreve no canal

	fmt.Println("Aguardando mensagem...") // imprime uma mensagem enquanto aguarda a mensagem do canal

	//for {
	//	mensagem, aberto := <-canal
	//	if !aberto {
	//		fmt.Println("Canal está fechado") // verifica se o canal está aberto
	//	break
	//	}
	//	fmt.Println(mensagem) // recebe a mensagem do canal e imprime
	//
	//}

	for mensagem := range canal { // recebe a mensagem do canal e imprime
		fmt.Println(mensagem) // imprime a mensagem recebida
	}
	fmt.Println("Canal fechado") // imprime uma mensagem quando o canal é fechado
}

func escrever(texto string, canal chan string) {
	//time.Sleep(time.Second * 3) // aguarda 3 segundos antes de enviar a mensagem
	for i := 0; i < 5; i++ {

		canal <- texto // envia o texto para o canal
		// O programa continua executando enquanto a goroutine está em execução
		time.Sleep(1 * time.Second)
	}
	close(canal) // fecha o canal após enviar todas as mensagens
}
