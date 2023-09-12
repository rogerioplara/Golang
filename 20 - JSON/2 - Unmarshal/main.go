package main

import (
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
	// Dessa forma é possível ler o json
	cachorroEmJSON := `{"nome":"Rex","raca":"Dálmata","idade":3}`

	// Instância da struct
	var c cachorro

	// Passar os dados para dentro da struct
	// É necessário fazer um casting de string para um slice de bytes
	// A segunda variável é o endereço da instância da struct
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	cachorro2EmJSON := `{"nome":"Toby", "raca":"Poodle"}`

	c2 := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)

}
