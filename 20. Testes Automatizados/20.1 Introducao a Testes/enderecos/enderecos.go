package enderecos

import "strings"

// TipoDeEndereco recebe um endereço e retorna o tipo de endereço (rua, avenida, etc.)
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia", "alameda", "travessa", "largo", "vila", "praça"}

	endereco = strings.ToLower(endereco)
	primeiraPalavra := strings.Split(endereco, " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if strings.EqualFold(primeiraPalavra, tipo) {
			enderecoTemUmTipoValido = true
			break
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavra)
	}
	return "tipo inválido"

}
