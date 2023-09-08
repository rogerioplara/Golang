package main

import "fmt"

func main() {
	/*
		Declaração de um Map

		Maps só armazenam um tipo de dados e deve ser especificado na sua declaração

		varmap := map[tipo chave] tipo valor {}
	*/

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)

	// Para acessar uma das chaves
	fmt.Println(usuario["nome"], usuario["sobrenome"])

	// Aninhamento de maps
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Maringá",
		},
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2["nome"]["ultimo"])

	// Para apagar uma das chaves do map
	delete(usuario2, "curso")
	fmt.Println(usuario2)

	// Para adicionar outra chave no map
	usuario2["signo"] = map[string]string{
		"mes":   "Maio",
		"signo": "Gêmeos",
	}
	fmt.Println(usuario2)

}
