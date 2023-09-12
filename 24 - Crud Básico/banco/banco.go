package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o MySQL
)

// Conectar não tem nenhum parâmetro passado porém retorna o ponteiro para o banco e o erro
// Aqui deve ser passada a string de conexão
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		// Caso ocorra erro, retorna nil para o ponteiro e o erro
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		// Caso ocorra erro de credencial ou outro erro não visto anteriormente, retora esse if
		return nil, erro
	}

	// Se tudo correr bem, retorna o db e o erro como nil
	return db, nil
}
