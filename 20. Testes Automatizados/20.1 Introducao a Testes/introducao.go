package main

import (
	"fmt"
	"introducao-a-testes/enderecos"
)

func main() {

	tipoEndereco := enderecos.TipoDeEndereco("Estrada das Flores")
	fmt.Println("Tipo de Endereço:", tipoEndereco)
}
