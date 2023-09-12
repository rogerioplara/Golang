package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// CRUD - CREATE, READ, UPDATE, DELETE

	/*
		CREATE - POST
		READ - GET
		UPDATE - PUT
		DELETE - DELETE

		Será utilizado o mux para gerenciar as funções
		Variável que vai conter todas as rotas (ROUTER)
	*/
	router := mux.NewRouter()

	// Quando minha rota /usuarios receber um post, será executada essa função
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	// Para buscar um usuário específico, é necessário passar o id pela URL (nesse caso)
	// Inserido entre chaves para denotar que é um valor dinâmico
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", servidor.DeletarUsuario).Methods(http.MethodDelete)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
