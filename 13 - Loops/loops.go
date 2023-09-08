package main

import (
	"fmt"
)

func main() {
	i := 0

	// Em Go, não existem loops além do For
	// Entretanto, o for pode ser usado como um while, ou como um for normal
	for i < 10 {
		i++
		fmt.Printf("Incrementando i -> %d\n", i)
	}
	fmt.Println("---------")

	for j := 0; j < 10; j += 2 {
		fmt.Printf("Incrementando j -> %d\n", j)
		// time.Sleep(time.Second)
	}
	fmt.Println("---------")

	// Cláusula range -> utilizada para iteração, assemelha-se a um foreach
	// O primeiro parâmetro sempre vai ser o índice, para ignorar é possível utilizar o _
	nomes := [3]string{"João", "Davi", "Lucas"}

	for _, value := range nomes {
		fmt.Println(value)
	}
	fmt.Println("---------")

	// Nesse caso vai a letra vai retornar os códigos da tabela ascii
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra)) // Para trazer as letras, necessário fazer o casting
	}
	fmt.Println("---------")

	// Range de um map
	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
	fmt.Println("---------")
}
