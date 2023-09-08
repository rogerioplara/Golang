package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	// Declarando um array
	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"P1", "P2", "P3", "P4", "P5"}
	fmt.Println(array2)

	// Arrays não são flexíveis, o tipo e a quantidade de posições são fixos

	// É possível declarar também dessa forma, porém será fixado baseado no que for passado -> não torna o array dinâmico
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// Geralmente, utiliza-se o Slice
	// Declarando um slice
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17} // não especificar o tamanho
	fmt.Println(slice)

	// Arrays e Slices são de tipos diferentes, portanto não é possível realizar operações entre eles
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	// append
	slice = append(slice, 18)
	fmt.Println(slice)

	// O slice é uma "fatia" de um array, ele referencia o array
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)

	fmt.Println("---------------")
	fmt.Println("Arrays Internos")
	/*
		Arrays Internos
	*/
	// make(tipo da variável, tamanho, capacidade)
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho do slice
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 5)

	// Quando estouramos a capacidade do slice, o Go cria um outro array e dobra sua capacidade
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // tamanho do slice
	fmt.Println(cap(slice3)) // capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
