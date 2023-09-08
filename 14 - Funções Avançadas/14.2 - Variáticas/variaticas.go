package main

import "fmt"

// Funções variáticas podem receber n parametros
func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}

// Combinando parâmetros fixos com parâmetros variáveis
// Só é possível ter 1 parâmetro variático por função e deve ser o último parâmetro
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	// A chamada da função envia os parâmetros como um slice
	totalSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 100, 60, 800)
	fmt.Println(totalSoma)

	soma2 := soma(2, 4)
	fmt.Println(soma2)

	escrever("Olá, Mundo", 10, 2, 4, 6)

}
