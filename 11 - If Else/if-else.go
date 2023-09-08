package main

import "fmt"

func main() {
	numero := -11

	// If - Else
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// IF Initi - inicia a variável dentro do if -> variável só é válida dentro do escopo do IF
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if numero < -10 {
		fmt.Println("Número é menor que -10")
	} else {
		fmt.Println("entre 0 e -10")
	}

}
