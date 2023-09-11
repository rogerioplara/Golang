package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARALELISMO

	/*
		Paralelismo acontece quando existem 2 ou mais tarefas ocorrendo ao mesmo tempo em núcleos diferentes do processador. As tarefas são distribuídas entre os núcleos.

		Concorrência não necessáriamente ocorrem ao mesmo tempo. É possível aplicar concorrência em processadores de núcleo único. As tarefas "revezam" o tempo quando isso acontece.
	*/

	go escrever("Olá Mundo!") // goroutine -> executa a função mas não espera o fim dela para seguir com o programa
	escrever("Programando em Go!")
}

// Função exemplo que fica escrevendo o que for passado infinitamente todo segundo
func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
