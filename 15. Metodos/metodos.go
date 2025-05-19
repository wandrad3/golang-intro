package main

import (
	"fmt"
)

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuario Nome: %s, Idade: %d\n", u.nome, u.idade)
}
func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {

	usuario1 := usuario{"Wesley", 31}
	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Debora", 24}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()

	fmt.Println("Maior de idade: ", maiorDeIdade)
	usuario2.fazerAniversario()
	println(usuario2.idade)
	usuario2.salvar()

}
