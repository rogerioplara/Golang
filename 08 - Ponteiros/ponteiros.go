package main

import "fmt"

func main() {

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// Ponteiro é uma referência de memória
	var variavel3 int
	var ponteiro *int

	fmt.Println(variavel3, ponteiro) // 0 <nil>

	variavel3 = 100
	ponteiro = &variavel3 // 100 endereço de variavel3
	fmt.Println(variavel3, ponteiro)

	// É possível visualizar o valor por dentro do ponteiro
	// Desreferenciação
	fmt.Println(*ponteiro) // 100

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro) // 150 150
}
