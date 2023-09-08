package main

import (
	"errors"
	"fmt"
)

func main() {
	// Tipos básicos

	/*
		Existem 4 tipos de números inteiros no Go
		int8, int16, int32, int64
		A diferença é a quantidade de bits que os números comportam
		pode ser declarado como int somente, nesse caso será utilizado os bits da máquina, 64 ou 32
	*/

	numero := 10000000000
	fmt.Printf("%T\n", numero) // forma de verificar o tipo de uma variável

	// uint -> unsigned int
	var numero2 uint32 = 2020
	fmt.Println(numero2)

	// alias
	// int32 = rune - utilizado quando estamos utilizando a tabela ascii
	var numero3 rune = 12345
	fmt.Println(numero3)

	//byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	// FLOAT

	/*
		Tem somente 2 tipos:
		float32, float64
	*/

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 112300000.23
	fmt.Println(numeroReal2)

	numeroReal3 := 321.789
	fmt.Println(numeroReal3)

	// STRING
	var str string = "Texto"
	fmt.Println(str)

	// No go não existe o tipo CHAR, nesse caso será convertido para número
	str2 := "Texto2"
	fmt.Println(str2)

	// CHAR - o mais perto que existe no go
	char := 'B' // converte o char para o número da tabela ascii
	fmt.Println(char)

	// ZERO - valor atribuído às variáveis não inicializadas
	var texto string
	fmt.Println(texto)

	var num int16
	fmt.Println(num)

	// BOOLEAN
	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool // se não iniciada, é falsa
	fmt.Println(booleano2)

	//ERRO
	var erro1 error
	fmt.Println(erro1)
	// retorna <nil> -> tipo de dado que serve para muitas coisas em Go, como se fosse um zero

	// para criar um erro, deve ser utilizado o pacote errors
	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)
}
