package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar retorna a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	// Instanciação da biblioteca
	app := cli.NewApp()
	// Nome da aplicação
	app.Name = "Aplicação de Linha de Comando"
	// Uso da aplicação
	app.Usage = "Busca IPs e Nomes de Serviores na internet"
	// Esses são parâmetros básicos da aplicação

	// Declaração das flags fora do comando e depois passado como variável
	flags := []cli.Flag{
		cli.StringFlag{
			// Parâmetro -> inserido após a indicação da função --host
			Name: "host",
			// Valor -> esse é o valor padrão, deve ser inserido o desejado após a flag/parâmetro
			Value: "devbook.com.br",
		},
	}

	// Lista de flags para análise
	app.Commands = []cli.Command{
		{
			// Primeira função
			Name:  "ip",
			Usage: "Busca IPs de endereços na internet",
			Flags: flags,
			// Ação que o comando executará
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app
}

// Função da ação que será executada para buscar ips
func buscarIps(c *cli.Context) {
	host := c.String("host")

	// Pacote net
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

// Função da ação que será executada para buscar servidor
func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host) // name server
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
