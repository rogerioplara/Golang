package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

/*
Esta aplicação busca ips públicos e servidores de hospedagem de sites que desejar

Como rodar ex:
go run main.go ip --host amazon.com.br
go run main.go servidores --host amazon.com.br

Como já está compilado num executável, só precisa utilizar o nome do arquivo ao invés de "go run main.go..."
*/
func main() {
	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal()
	}
}
