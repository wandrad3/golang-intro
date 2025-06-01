// TESTE DE UNIDADE
package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Avenida Paulista, 1578, São Paulo - SP"

	tipoEsperado := "Avenida"

	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)
	if tipoDeEnderecoRecebido != tipoEsperado {
		t.Errorf("Tipo de endereço recebido: %s, esperado: %s", tipoDeEnderecoRecebido, tipoEsperado)
	}

}
