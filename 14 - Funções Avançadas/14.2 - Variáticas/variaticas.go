package main

import "fmt"

// Funções variáticas podem receber n parametros
func soma(numeros ...int) {
	fmt.Println(numeros)
}

func main() {
	soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 100, 60, 800)
}
