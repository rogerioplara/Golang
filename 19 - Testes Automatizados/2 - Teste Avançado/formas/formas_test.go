package formas

import (
	"math"
	"testing"
)

/*
TDD - Test Drive Development
Essa abordagem busca desenvolver o teste antes de desenvolver a função que será testada
*/

// TestArea vai testar os métodos das duas áreas
func TestArea(t *testing.T) {

	// t.Run é um subteste da função principal
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			// Fatalf para a execução dos testes quando encontrar algum erro
			t.Fatalf("Area recebida %.2f é diferete da esperada %.2f", areaRecebida, areaEsperada)
		}
	})

	t.Run("Circulo", func(t *testing.T) {
		circ := Circulo{10}
		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			// Fatalf para a execução dos testes quando encontrar algum erro
			t.Fatalf("Area recebida %.2f é diferete da esperada %.2f", areaRecebida, areaEsperada)
		}
	})
}
