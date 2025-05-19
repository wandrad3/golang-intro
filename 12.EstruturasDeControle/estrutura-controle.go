package main

import "fmt"

func diaDaSemana(numero int) string {

	var diaDasemana string
	switch numero {
	case 1:
		diaDasemana = "Domingo"
	case 2:
		diaDasemana = "Segunda-feira"
	case 3:
		diaDasemana = "Terça-feira"
	case 4:
		diaDasemana = "Quarta-feira"
	case 5:
		diaDasemana = "Quinta-feira"
	case 6:
		diaDasemana = "Sexta-feira"
	case 7:
		diaDasemana = "Sábado"
		//fallthrough
		//não existe break no switch, o fallthrough é usado para continuar a execução do próximo case

	default:
		diaDasemana = "Número inválido"

	}
	return diaDasemana
}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-feira"
	case numero == 3:
		return "Terça-feira"
	case numero == 4:
		return "Quarta-feira"
	case numero == 5:
		return "Quinta-feira"
	case numero == 6:
		return "Sexta-feira"
	case numero == 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func main() {
	fmt.Println("Estruturas de controle")

	// Estruturas de controle são usadas para controlar o fluxo de execução do programa.
	// Elas permitem que o programa execute diferentes blocos de código com base em condições específicas.
	// As principais estruturas de controle em Go são: if, else, switch, for e select.

	numero := 10
	if numero > 15 { //parenteses no if são opcionais
		fmt.Println("O número é maior que 15")
	} else {
		fmt.Println("O número não é maior que 15")
	} //em go os condicionais devem ter chaves

	if outroNumero := numero; outroNumero > 0 { // if com inicialização
		fmt.Println("O número é maior que zero")
		//a variavel não pode ser usada fora do escopo do bloco if
	}

	// Switch - é uma estrutura de controle que permite executar diferentes blocos de código com base no valor de uma variável.
	fmt.Printf("Dia da semana é " + diaDaSemana(3)) //chama a função diaDaSemana

	diaDaSemana := diaDaSemana2(5)               //chama a função diaDaSemana2
	fmt.Printf("Dia da semana é " + diaDaSemana) //chama a função diaDaSemana2

}
