package main

import "fmt"

func main() {
	// ARITMÉTICOS
	/*
		+ - / * %

		Operadores padrão
	*/

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	resto := 10 % 3

	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)

	/*
		Não é permitido comparar ou realizar operações com tipos de dados diferentes
		var numero1 int16 = 10
		var numero2 int32 = 25
		soma := numero1 + numero2
	*/

	// ATRIBUIÇÃO
	// = ou := (inferência)

	// RELACIONAIS
	/*
		== < > <= >= !=

		Operadores padrão
	*/

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	// LÓGICOS
	/*
		&& and
		|| or
		! not
	*/
	fmt.Println("-----------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// UNÁRIOS
	fmt.Println("-----------------")
	numero := 10
	numero++     // 11
	numero += 15 // 26
	numero--     // 25
	numero -= 20 // 5
	numero *= 3  // 15
	numero /= 2  // 7
	numero %= 2  // 1
	fmt.Println(numero)
	// não existe pré e pós fixado no Go

	// não existe operador ternário em Go

}
