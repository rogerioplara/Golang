package main

import "fmt"

// Variável passada por valor (cópia)
func inverterSinal(numero int) int {
	return numero * -1
}

// Variável passada por referência (endereço)
func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)

	// Dessa forma a variável numero não é alterada, é criada uma cópia
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	fmt.Println("------------")

	// Com ponteiros é possível alterar o valor da variável original
	novoNumero := 40
	fmt.Println(novoNumero) // 40
	// Executando a função e alterando a variável
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero) // -40

}
