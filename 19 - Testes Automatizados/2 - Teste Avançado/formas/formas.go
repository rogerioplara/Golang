package formas

import (
	"math"
)

type Retangulo struct {
	Altura  float64
	Largura float64
}

// Implementação da interface na struct
func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	Raio float64
}

// Implementação da interface na struct
func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

// Declaração da interface -> só podem conter assinaturas de métodos
type Forma interface {
	Area() float64
}
