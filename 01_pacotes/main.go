package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("rogerioo.lara@hotmail.com")
	fmt.Println(erro)

	//auxiliar2.escrever2() -> não é possível importar ou utilizar essa função devido ao nome da função ter o nome com letra minúscula
}
