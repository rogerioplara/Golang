package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		O Select é como se fosse um switch mas específico para utilziar com concorrência
	*/

	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		/*
			O select tem a mesma estrutura do switch
			Trata os canais como independentes e executa conforme cada canal esteja pronto para ser executado. No exemplo abaixo, o canal 1 executa 4 vezes enquanto o 2 executa 1
		*/
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}

}
