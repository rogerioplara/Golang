package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	/*
		Tipos genéricos
		É possível utilizar funções  genéricas que não vão ter restrição de tipos
	*/

	generica("string")
	generica(1)
	generica(true)

	fmt.Println()

	// Tipos genéricos permitem estruturas como essa:
	mapa := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "string",
		true:         float64(12),
	}
	fmt.Println(mapa)

}
