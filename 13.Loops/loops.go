package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	//i := 0
	//for i < 10 {
	//	fmt.Println(fmt.Sprintf("Incrementando i: %d", i))
	//	time.Sleep(time.Second)
	//	i++
	//}

	//for j := 0; j < 10; j++ {
	//	fmt.Println(fmt.Sprintf("Incrementando j: %d", j))
	//	time.Sleep(time.Second)
	//}

	nomes := [5]string{"Lucas", "Wesley", "Débora", "João", "Maria"}
	for i := 0; i < len(nomes); i++ {
		fmt.Println("Nome ", nomes[i])
	}
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "Palavra" {
		fmt.Println("Letra ", indice, letra)
	}

	usuario := map[string]string{
		"nome":      "Wesley",
		"sobrenome": "de Andrade Santos",
		"idade":     "31",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
