package enderecos_test

import (
	// É possível importar o pacote com um alias antes da importação
	// Ao utilizar o . no nome do alias, o Go utilizará as funções do pacote importado como função padrão
	// Geralmente esse tipo de importação é utilizado em testes
	. "introducao-testes/enderecos"
	"testing"
)

/*
	TESTE DE UNIDADE
	Testa uma pequena parte do código, geralmente uma função

	Testes em Go precisam obrigatóriamente estar em arquivos separados daqueles que estão testando
	O nome do arquivo deve ser sempre com o sufixo _test.go

	Essa é a assinatura de uma função de teste:
	TestNomeDaFuncaoQueIraTestar(t *testing.t){
	}

	Para chamar o teste, basta executar no terminal, dentro da pasta que deseja testar
	"go test"

	Funcionalidades do comando 'go test'
	'go test ./...' 	-> entra em todos os arquivos do projeto e executa os testes
	'go test -v' 		-> modo verboso, é mais descritivo
	'go test --cover' 	-> verifica a cobertura do teste que está rodando
	'go test --coverprofile cobertura.txt' 	-> gera um txt sobre a cobertura dos testes
	'go tool cover --func=cobertura.txt' 	-> lê o resultado de cada função testada
	'go tool cover --html=cobertura.txt'	-> mostra um html com um feedback visual das linhas que não estão cobertas pelo teste
*/

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	// É possível rodar os testes em paralelo
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"avenida ABC", "Avenida"},
		{"RODOVIA ABC", "Rodovia"},
		{"Estrada ABC", "Estrada"},
		{"RUA dos Bobos", "Rua"},
		{"Avenida dos Bobos", "Avenida"},
		{"estrada dos Bobos", "Estrada"},
		{"Rodovia dos Bobos", "Rodovia"},
		{"praça das Rosas", "Tipo Inválido"},
		{"12345mil", "Tipo Inválido"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		TipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if TipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido \"%s\" é diferente do esperado \"%s\"",
				TipoDeEnderecoRecebido,
				cenario.retornoEsperado)
		}
	}

}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste Quebrou 😢")
	}
}
