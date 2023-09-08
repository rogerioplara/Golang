package main

import "fmt"

// Sintaxe básica de uma struct
type usuario struct {
	nome     string
	idade    uint8
	endereco endereco // inserindo o endereço que é uma outra struct
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	// Maneiras de criar uma struct
	var u usuario
	u.nome = "Rogerio"
	u.idade = 30

	fmt.Println(u)

	enderecoPriscilla := endereco{"Rua Tiradentes", 2}
	// Utilizando inferência de tipos
	u2 := usuario{"Priscilla", 29, enderecoPriscilla}
	fmt.Println(u2)

	// É possível criar a struct sem ter todos os dados
	u3 := usuario{nome: "Tico"}
	fmt.Println(u3)
}
