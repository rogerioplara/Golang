package main

import "fmt"

func main() {
	/*
		Buffer -> especoficar a capacidade do canal
		Só será bloqueado quando atingir a capacidade máxima do canal
	*/

	canal := make(chan string, 2) // Buffer

	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
