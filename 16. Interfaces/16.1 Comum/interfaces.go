package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A area da forma é: %0.2f\n", f.area())
}

func (r retangulo) area() float64 {
	return r.base * r.altura
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func (q quadrado) area() float64 {
	return q.base * q.base
}

type retangulo struct {
	base   float64
	altura float64
}
type circulo struct {
	raio float64
}
type quadrado struct {
	base float64
}

func main() {

	r := retangulo{base: 10, altura: 5}
	escreverArea(r)
	c := circulo{raio: 10}
	escreverArea(c)

	q := quadrado{base: 10}
	escreverArea(q)
	fmt.Println("A area do quadrado é: ", q.area())

}
