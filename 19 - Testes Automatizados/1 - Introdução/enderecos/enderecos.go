package enderecos

import "strings"

// TipoDeEndereco verifica se o endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)

	// Função split divide a string passada de acordo com o separador passado
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0] // o [0] especifica que deseja retornar somente o primeiro índice do slice
	// RUA DOS BOBOS -> parâmetro
	// ["RUA", "DOS", "BOBOS"] -> retorno
	// No caso como foi especificado que deseja somente o priemro índice:
	// ["RUA"]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"
}
