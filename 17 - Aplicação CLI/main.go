package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

/*
Esta aplicação busca ips públicos de sites que desejar

Como rodar ex:
go run main.go ip --host amazon.com.br
*/
func main() {
	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal()
	}

}
