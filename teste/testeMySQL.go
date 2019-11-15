package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUsuario := "root"
	dbSenha := "senha"
	dbTabela := "login"
	db, err := sql.Open("mysql", dbUsuario+":"+dbSenha+"@/"+dbTabela)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var a string
	usuario := db.QueryRow("SELECT * FROM palavras WHERE id = 1")
	usuario.Scan(&a)
	fmt.Println(a)

}
