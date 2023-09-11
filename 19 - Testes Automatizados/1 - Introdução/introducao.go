package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(tipoEndereco)
	tipoEndereco2 := enderecos.TipoDeEndereco("Rua dos Bobos")
	fmt.Println(tipoEndereco2)
	tipoEndereco3 := enderecos.TipoDeEndereco("Preça dos Expedicionários")
	fmt.Println(tipoEndereco3)
}
