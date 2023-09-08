package main

import "fmt"

// É possível declarar o switch de uma forma tradicional
func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Insira um número de 1 a 7"
	}
}

// Ou declarar dessa forma, testando as condições dentro dos cases. Essa forma é bem útil para realizar testes de origens diferentes
func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	case numero == 3:
		return "Terça-Feira"
	case numero == 4:
		return "Quarta-Feira"
	case numero == 5:
		return "Quinta-Feira"
	case numero == 6:
		return "Sexta-Feira"
	case numero == 7:
		return "Sábado"
	default:
		return "Insira um número de 1 a 7"
	}
}

func main() {
	dia := diaDaSemana(20)
	fmt.Println(dia)

	fmt.Println("----------")

	dia2 := diaDaSemana2(4)
	fmt.Println(dia2)
}
