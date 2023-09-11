package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Os canais podem enviar ou receber dados -> canal de comunicação
		Esses canais são utilizados para sincronizar as goroutines -> é o método mais utilizado

		Os canais podem fazer envio ou podem receber dados
		Essas operações bloqueiam a execução do programa
	*/

	// Criação do canal: make(chan tipo-de-dado) -> só poderão trafegar informações desse tipo neste canal
	canal := make(chan string)

	go escrever("Olá Mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	// for {
	// 	// Essa segunda variável "aberto" é uma propriedade dos canais que são passados juntos
	// 	// A primeira é o canal em si e a segunda é se o canal está ou não aberto
	// 	mensagem, aberto := <-canal // recebendo o valor no canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	// Mesma função anterior, porém simplificada
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

// A função também vai receber o canal do tipo string
func escrever(texto string, canal chan string) {
	// time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- texto // passando o valor para dentro do canal
		time.Sleep(time.Second)
	}

	// Fecha o canal para evitar deadlock
	close(canal)
}
