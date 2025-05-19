package main

import "fmt"

func closure() func() {
	texto := "Texto dentro da função closure"
	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	texto := "Texto dentro da função main"
	fmt.Println(texto)
	funcaoNova := closure()
	funcaoNova()

}
