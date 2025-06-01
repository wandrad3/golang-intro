// TESTE DE UNIDADE
package enderecos

import "testing"

type CenarioDeTeste struct {
	enderecoRecebido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	CenarioDeTestes := []CenarioDeTeste{
		{enderecoRecebido: "Rua das Flores, 123", retornoEsperado: "Rua"},
		{enderecoRecebido: "Avenida Brasil, 456", retornoEsperado: "Avenida"},
		{enderecoRecebido: "Estrada do Sol, 789", retornoEsperado: "Estrada"},
		{enderecoRecebido: "Rodovia dos Imigrantes, 101", retornoEsperado: "Rodovia"},
		{enderecoRecebido: "Alameda Santos, 202", retornoEsperado: "Alameda"},
		{enderecoRecebido: "Travessa da Paz, 303", retornoEsperado: "Travessa"},
		{enderecoRecebido: "Largo da Liberdade, 404", retornoEsperado: "Largo"},
		{enderecoRecebido: "Vila Mariana, 505", retornoEsperado: "Vila"},
		{enderecoRecebido: "Praça da Sé, 606", retornoEsperado: "Praça"},
		{enderecoRecebido: "Avenida Paulista, 1578", retornoEsperado: "Avenida"},
		{enderecoRecebido: "Rua Desconhecida, 999", retornoEsperado: "Rua"},
		{enderecoRecebido: "Tipo Inválido, 888", retornoEsperado: "tipo inválido"},
	}

	for _, cenario := range CenarioDeTestes {
		tipoRecebido := TipoDeEndereco(cenario.enderecoRecebido)
		if tipoRecebido != cenario.retornoEsperado {
			t.Errorf("Endereço: %s, Tipo recebido: %s, Tipo esperado: %s", cenario.enderecoRecebido, tipoRecebido, cenario.retornoEsperado)
		}

	}
}
