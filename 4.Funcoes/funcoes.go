package main

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func subtrair(n1 int8, n2 int8) int8 {
	return n1 - n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8, int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2
	divisao := n1 / n2

	return soma, subtracao, multiplicacao, divisao
}

func main() {
	//somar
	soma := somar(10, 20)
	println(soma)

	//subtrair
	subtracao := subtrair(20, 10)
	println(subtracao)

	//chamar funcao sem retorno
	//imprimir()

	var f = func(texto string) string {
		println(texto)
		return texto
	}
	f("Texto de teste")

	resultado := f("Outro texto")

	println(resultado) //imprime o endereço da função

	resultadoSoma, resultadoSubtracao, resultadoMultiplicacao, _ := calculosMatematicos(10, 5)

	println(resultadoSoma, resultadoSubtracao, resultadoMultiplicacao) //imprime o endereço da função
}
