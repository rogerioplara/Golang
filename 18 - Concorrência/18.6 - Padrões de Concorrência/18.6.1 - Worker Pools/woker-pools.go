package main

import "fmt"

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

/*
É possível especificar as funções dos canais, no caso abaixo o canal TAREFAS somente recebe dados e o de RESULTADOS somente envia dados.

A worker-pools é utilizada quando existe uma fila de tarefas a ser executadas e que seja possível separar essas tarefas recursivamente.

No exemplo, a função worker é chamada diversas vezes, dividindo assim o processo em partes menores, agilizando a execução do programa
*/
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
