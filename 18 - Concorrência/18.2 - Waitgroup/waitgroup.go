package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	/*
		Sincronizando as goroutines com WaitGroup
	*/
	// Criação da variável waitGroup e "instanciando" um sync.WaitGroup
	var waitGroup sync.WaitGroup

	// Definição de quantas goroutines serão adicionadas ao grupo
	waitGroup.Add(4)

	// Criação da função anônima para que execute a função do waitgroup
	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done() // Esse método tira um do contador adicionado no Add
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Goroutine 3!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Goroutine 4!")
		waitGroup.Done()
	}()

	waitGroup.Wait() // Força o programa a esperar as goroutines finalizar e o contador chegar em 0
}

// Função exemplo que fica escrevendo o que for passado infinitamente todo segundo
func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
