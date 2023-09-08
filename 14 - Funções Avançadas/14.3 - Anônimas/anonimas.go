package main

import "fmt"

func main() {

	// A função anônima precisa do parênteses no final para ser executada
	func(texto string) {
		fmt.Println(texto)
	}("Olá Mundo")

	// É possível armazenar o resultado da função na variável
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Parâmetro passado")

	fmt.Println(retorno)
}
