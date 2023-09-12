package main

import (
	"database/sql"
	"fmt"
	"log"

	// Import implícito. Não será esse arquivo que vai usar o driver, mas sim o pacote SQL
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/*
		Realizar o download do driver do mysql para o go:
		github.com/go-sql-driver/mysql

		Cada banco de dados possui um driver específico

		Estrutura básica de uma string de conexão
		stringConexao := "usuario:senha@/banco" -> dessa forma já funcionaria, porém existem outros parâmetros que podem ser passados
	*/

	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	// Pacote SQL do Go -> ("driver name", string de conexão)
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}
	// Realizar um DEFER para fechar a conexão
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta!")

	// Realizando a query
	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}
