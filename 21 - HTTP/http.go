package main

import (
	"log"
	"net/http"
)

/*
As funções de rotas geralmente são escritas dessa forma, fora da função main()

Dessa forma, o que precisa ser feito é no http.HandleFunc(), passar os parâmetros com o nome da página e a função que a representa. Não importa o nome da função, desde que seja chamada no handler

O que importa nesse caso são os parâmetros W e R
r - request / w - response
*/
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar Página de Usuários"))
}

func main() {
	/*
		HTTP É UM PROTOCOLO DE COMUNICAÇÃO - BASE DA COMUNICAÇÃO WEB

		CLIENTE - SERVIDOR
		Request - Reponse
			- Cliente faz uma requisição ao servidor
			- Servidor processa a requisição e devolve uma resposta

		Essa troca é feita por mensagens

		Outro aspecto importante são as Rotas
		ROTAS
			- URI -> Identificador do Recurso
			- MÉTODO -> GET, POST, PUT, DELETE
				Get: buscar dados
				Post: cadastrar dados
				Put: atualizar dados
				Delete: deletar dados
	*/

	// Essa função faz handler do URI - Request e Reponse no /home
	http.HandleFunc("/home", home)
	// dessa forma que funciona o roteamento das páginas
	http.HandleFunc("/users", users)

	// Criação do servidor e subindo na porta 5000
	log.Fatal(http.ListenAndServe(":5000", nil))
}
