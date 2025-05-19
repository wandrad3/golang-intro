package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	peso      uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Heranca")

	var estudante1 = estudante{
		pessoa: pessoa{
			nome:      "Wesley",
			sobrenome: "de Andrade Santos",
			idade:     31,
			peso:      92,
			altura:    173,
		},
		curso:     "Tecnologia de Segurança da Infomação",
		faculdade: "Fatec Araraquara",
	}

	fmt.Println(estudante1)

	fmt.Println(estudante1.nome)
}
