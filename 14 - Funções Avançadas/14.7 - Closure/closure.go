package main

import "fmt"

/*
	Closure -> permite acessar a variável fora do escopo
*/

func closure() func() {
	texto := "Dentro da função Closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função Main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
