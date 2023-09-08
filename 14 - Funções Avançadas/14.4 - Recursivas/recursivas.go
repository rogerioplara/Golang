package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	/*
		Funçõe recursivas chamam elas mesmas
		Ex: Sequência de fibonacci -> o próximo número é sempre a soma dos dois números anteriores
		1 1 2 3 5 8 13
	*/

	posicao := uint(12)
	for i := uint(1); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
