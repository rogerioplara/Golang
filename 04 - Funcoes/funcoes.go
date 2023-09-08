package main

import "fmt"

// O tipo de retorno de uma função deve ser declarado após os parâmetros
func somar(n1 int32, n2 int32) int32 {
	return n1 + n2
}

// Em Go, as funções podem ter mais de um retorno
func calculosMatematicos(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	// É possível atribuir funções a variáveis, funções também são tipos
	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Texto da função 1")

	// Obrigatóriamente deve ser recebido os dois resultados
	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// Posso também escolher quais os resultados que desejo com o _
	resultadoSoma2, _ := calculosMatematicos(20, 30)
	fmt.Println(resultadoSoma2)

	_, resultadoSubtracao2 := calculosMatematicos(20, 30)
	fmt.Println(resultadoSubtracao2)
}
