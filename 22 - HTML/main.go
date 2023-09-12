package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Passando o template para ser executado no caminho home
	// Parâmetros: writer, template, dados
	u := usuario{
		"João",
		"joao.pedro@gmail.com",
	}
	templates.ExecuteTemplate(w, "home.html", u) // struct 'u' passada na página
}

type usuario struct {
	Nome  string
	Email string
}

// Utilização da biblioteca html/template do Go, instanciada fora da func main
// Convenção
var templates *template.Template

func main() {
	// Jogando na variável templates todos os templates que serão criados
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	fmt.Println("Executando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
