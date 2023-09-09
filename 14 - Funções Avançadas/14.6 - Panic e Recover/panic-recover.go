package main

import "fmt"

/*
	O recover recupera a execução do programa setando dessa forma:
*/
func recuperarExecuca() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

/*
	O panic termina com a execução do programa, chamando todas as funções que tem defer
*/

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecuca()

	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É EXATAMENTE 6!")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução!")
}
