package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) string {
	defer fmt.Println("Média calculada. Resultado será retornado")

	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2
	if media >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {

	/*
		Defer = Adiar

		Adia a execução da função até o último momento possível, imediatamente antes do return
	*/

	// defer funcao1()
	// funcao2()

	fmt.Println(alunoEstaAprovado(3, 4))
}
