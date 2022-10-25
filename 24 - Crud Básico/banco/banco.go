package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Conexão com banco de dados
func Conectar() (*sql.DB, error) {

	stringConexao := "root:admin@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil

}
