package main

import (
	"fmt"
	"math"
)

type retangulo struct {
	altura  float64
	largura float64
}

// Implementação da interface na struct
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

// Implementação da interface na struct
func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

// Declaração da interface -> só podem conter assinaturas de métodos
type forma interface {
	area() float64
}

// Função que implementa a interface área
func escreverArea(f forma) {
	fmt.Printf("A área da forma é %.2f\n", f.area())
}

func main() {
	/*
		Interfaces são muito utilizadas quando é necessário ter flexibilidade com os tipos
	*/

	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}
