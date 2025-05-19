package main

import "fmt"

type usuario struct {
	nome      string
	sobrenome string
	idade     uint8
	endereco  endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var enderecoExemplo = endereco{logradouro: "Rua das Flores", numero: 123}

	usuario1 := usuario{
		nome:      "Lucas",
		sobrenome: "Silva",
		idade:     25,
		endereco:  enderecoExemplo,
	}

	fmt.Println(usuario1)

	var usuario2 usuario

	usuario2.nome = "Wesley"
	usuario2.sobrenome = "de Andrade Santos"
	usuario2.idade = 31
	usuario2.endereco = endereco{logradouro: "Rua das Palmeiras", numero: 10}

	fmt.Println(usuario2)
	fmt.Println(usuario2.nome)

	usuario3 := usuario{nome: "Débora", idade: 24, endereco: enderecoExemplo}

	fmt.Println(usuario3)

}
