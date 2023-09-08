package main

import "fmt"

/*
	A função init será executada antes da função main, independente da ordem
	É permitido uma função init por arquivo
	Usada geralmente para realizar setups ou iniciar variáveis que serão utilizadas
*/

var n int

func main() {
	fmt.Println("Função Main sendo executada")

	fmt.Println(n)
}

func init() {
	fmt.Println("Executando a função init")

	n = 10
}
