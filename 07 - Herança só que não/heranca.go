package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}

type estudante struct {
	// colocando a struct dentro da que recebe a "herança", a struct "herda" as características da struct referenciada
	pessoa
	curso     string
	faculdade string
}

func main() {

	p1 := pessoa{"João", "Pedro", 20, 1.75}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1)

	// Ao fazer dessa forma, podemos acessar os atributos de pessoa pelo "objeto" estudante
	fmt.Println(e1.nome)
	fmt.Println(e1.sobrenome)
	fmt.Println(e1.altura)

	// Bem como acessar as propriedades do estudante
	fmt.Println(e1.faculdade)
	fmt.Println(e1.curso)

	// Também é possível acessar dessa forma
	fmt.Println(e1.pessoa.idade)

}
