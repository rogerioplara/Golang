package main

import "fmt"

// É possível nomear o retorno na declaração da função
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2 // Não precisa atribuir pois a variável já está declarada
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculosMatematicos(10, 20)
	fmt.Println(soma, subtracao)
}
