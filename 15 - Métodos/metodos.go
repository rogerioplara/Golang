package main

import "fmt"

/*
	A diferença de Métodos para funções é que os métodos são associadas a algo, seja uma struct ou interface
*/

type usuario struct {
	nome  string
	idade uint8
}

/*
	Declaração do método:

	func('alias' 'struct') nomeMetodo('parametros') retorno {
		Código do método
	}

*/
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// Método de incrementar um atributo da struct utilizando ponteiros
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 20}
	usuario1.salvar()

	usuario2 := usuario{"Davi", 30}
	usuario2.salvar()

	usuario3 := usuario{"Enzo", 12}
	usuario3.salvar()

	maiorDeIdade := usuario3.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario() // idade de 30 para 31
	fmt.Println(usuario2.idade)
}
