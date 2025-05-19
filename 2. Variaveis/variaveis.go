package main

import "fmt"

func main() {
	var nome string = "Wesley Andrade"
	idade := 31 //inferencia de tipo

	fmt.Println(nome)
	fmt.Println(idade)

	var (
		sexo   string  = "Masculino"
		altura float32 = 1.80
		peso   float32 = 80.0
	)
	fmt.Println(sexo, altura, peso)

	variavel1, variavel2 := "Wesley", 31
	fmt.Println(variavel1, variavel2)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	sexo, nome = nome, sexo
	fmt.Println(sexo, nome)
}
