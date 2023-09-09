package main

import "fmt"

func main() {
	// é possível declarar variáveis explicitamente tipadas ou implicitamente

	//declaração explícita
	var variavel1 string = "Variável 1"

	//declaração implícita -> inferência de tipo
	variavel2 := "Variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//também é possível declarar dessa forma:
	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	variavel7, variavel8, variavel9 := "Variável 7", 8, 9.1
	fmt.Println(variavel7, variavel8, variavel9)

	//no GO, o código não é compilado caso tenha variáveis ou importações não utilizadas.

	// CONSTANTES

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	// troca de valores de variáveis em Go, não precisa de variável auxiliar

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
