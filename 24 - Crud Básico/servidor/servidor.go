package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	// Ler o corpo da requisição
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	// Jogar essa requisição dentro de um struct de usuário
	// Instância da struct
	var usuario usuario

	// Recebe o struct e decompõe dentro da struct
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	// Abre a conexão com o banco de dados e depois fecha
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	// PREPARE STATEMENT - evitando sql injection
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()

	// Execução do statement, passando as variáveis que foram inseridas no struct
	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	// Retorno do id inserido e tratamento de erro correspondente
	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter id inserido"))
		return
	}

	/*
		STATUS CODES

		O Go tem os status codes na standard library, buscando por http.Status... é possível visualizar todos os códigos

		A cada inserção bem sucedida, retornará essa mensagem
	*/
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))

}

// BuscarUsuarios busca todos os usuários salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
	}
	defer db.Close()

	// SELECT * FROM usuarios
	linhas, erro := db.Query("SELECT * FROM usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários"))
	}
	defer linhas.Close()

	var usuarios []usuario
	// Iterar cada linha que for retornada
	for linhas.Next() {
		// Pega um usuário e popula o slice de usuários
		var usuario usuario

		// Realiza o scan em cada linha e insere no usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	// Transformação do slice de usuários em JSON
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}

}

// BuscarUsuario busca um usuário específico no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	// Ler o parâmetro que está sendo passado na rota
	// Essa função vai retornar todos os parâmetros
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	// Nesse caso é melhor abrir a conexão depois que confirmar a passagem do parâmetro corretamente
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
	}
	defer db.Close()

	// SELECT * FROM usuarios where id = ? -> um pouco diferente do prepare statement
	linha, erro := db.Query("SELECT * FROM usuarios WHERE id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuário"))
		return
	}
	defer linha.Close()

	// Criação da variável usuário e armazenamento do resultado nessa struct
	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear usuário"))
			return
		}
	}

	// Conversão do usuário para JSON
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuario para JSON"))
		return
	}
}

// AtualizarUsuario altera os dados de um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

// DeletarUsuario deleta o usuário do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	// Converter o parâmetro passado na URL para inteiro de base 10 e 32 bits
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	// Conecta na base de dados e depois fecha a conexão
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	// Cria o statement e depois fecha a conexão
	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	// Executa o statement, ignorando o primeiro parâmetro do SQL result
	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
