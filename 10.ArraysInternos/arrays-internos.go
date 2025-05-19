package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e slices")

	var array1 [5]int

	for i := 0; i < len(array1)-1; i++ {
		array1[i] = i + 1
		fmt.Println(i)
	}

	array2 := [5]string{"Lucas", "Wesley", "Débora", "João", "Maria"}
	fmt.Println(array2)

	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))  //imprime o tipo do slice
	fmt.Println(reflect.TypeOf(array2)) //imprime o tipo do array

	slice = append(slice, 110) //adiciona um elemento ao slice
	fmt.Println(slice)

	slice2 := array2[0:3] //fatiamento do array
	fmt.Println(slice2)   //imprime o fatiamento do array

	array2[0] = "Denyze Alterado" //altera o valor do array
	fmt.Println(slice2)           //imprime o array alterado

	//Array Interno
	slice3 := make([]float32, 10, 11) //cria um slice com 10 elementos e capacidade 11

	slice3 = append(slice3, 1.5) //adiciona um elemento ao slice
	slice3 = append(slice3, 6)   //adiciona um elemento ao slice

	fmt.Println(slice3)      //imprime o slice
	fmt.Println(len(slice3)) //imprime o tamanho do slice
	fmt.Println(cap(slice3)) //imprime a capacidade do slice

	slice4 := make([]float32, 5) //cria um slice com 5 elementos
	fmt.Println(slice4)          //imprime o slice
	fmt.Println(len(slice4))     //imprime o tamanho do slice
	fmt.Println(cap(slice4))     //imprime a capacidade do slice

	//resumindo: array é uma lista com tamanho fixo, enquanto slice é uma lista com tamanho variável
	//slice é uma abstração de array, ou seja, slice é um array com tamanho variável

}
