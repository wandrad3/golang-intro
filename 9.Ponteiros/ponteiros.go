package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")
	// Ponteiros são variáveis que armazenam o endereço de memória de outra variável.
	// Eles são úteis para manipular dados de forma eficiente e evitar cópias desnecessárias.
	// Ponteiros são representados pelo símbolo "&" antes do nome da variável.

	var variavel1 int = 10
	var ponteiro1 *int = &variavel1 // Cria um ponteiro que aponta para a variável variavel1

	fmt.Println("Valor da variável:", variavel1)    // Imprime o valor da variável
	fmt.Println("Endereço da variável:", ponteiro1) // Imprime o endereço da variável

	variavel1 = 20                                           // Altera o valor da variável
	fmt.Println("Valor da variável:", variavel1)             // Imprime o novo valor da variável
	fmt.Println("Endereço da variável:", ponteiro1)          // O endereço da variável permanece o mesmo
	fmt.Println("Valor apontado pelo ponteiro:", *ponteiro1) // Imprime o valor apontado pelo ponteiro (Desreferência)

	//valor padrao de um ponteiro
	var ponteiro2 *int                            // Cria um ponteiro sem inicializá-lo
	fmt.Println("Valor do ponteiro2:", ponteiro2) // Imprime o valor padrão do ponteiro (nil)

	// Go é uma linguagem segura em relação ao uso de ponteiros, pois não permite aritmética de ponteiros,
	// como em linguagens como C ou C++. Isso reduz a possibilidade de erros e vulnerabilidades.

	// Exemplo de função que altera o valor de uma variável usando ponteiros
	//func alterarValor(valor *int) {
	//	*valor = 50 // Altera o valor da variável para a qual o ponteiro aponta
	//}

	//fmt.Println("Antes de alterar:", variavel1)
	//alterarValor(&variavel1) // Passa o endereço da variável para a função
	//fmt.Println("Depois de alterar:", variavel1)

}
