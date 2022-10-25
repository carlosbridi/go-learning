package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	stringConexao := "root:admin@/devbook?charset=utf8"
	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		log.Fatal(erro)
	}

	linhas, erro := db.Query("SELECT * FROM USUARIOS")

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(linhas)
}
