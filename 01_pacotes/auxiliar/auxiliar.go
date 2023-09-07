package auxiliar

import "fmt"

// Escrever registra uma mensagem na tela -> ter um comentário acima da função é uma boa prática, esse comentário deve dizer o que a função faz
func Escrever() {

	fmt.Println("Escrevendo do pacote auxiliar")

	//aqui é possível chamar as funções que estão dentro do mesmo pacote, sem a necessidade de importar
	escrever2()

	//funções que começam com letra minúscula só são visíveis dentro do pacote que ela está -> go não tem public, private ou protected
}
