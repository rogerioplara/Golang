package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	// Dentro da crase é possível inserir a chave que esse atributo será dentro do json
	// Atributo tipo `json:"chaveJson"`
	// Serve como mapeamento do json
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	/*
		Converter um struct ou map para json
		json.Marshal()

		Converter json para struct ou map
		json.Unmmarshal()
	*/

	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)
	fmt.Println()

	// A função json.Marshal devolve o json e um erro
	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}
	// Retorna um slice de bytes []uint8
	fmt.Println(cachorroEmJSON)
	// Converter para algo legível -> função bytes.NewBuffer converte o slice de bytes em caracteres legíveis
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))
	fmt.Println()

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJSON)
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}
