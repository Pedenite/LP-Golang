package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUsuario := "root"
	dbSenha := "senha"
	dbBanco := "LP"
	db, err := sql.Open("mysql", dbUsuario+":"+dbSenha+"@/"+dbBanco)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var a string
	usuario := db.QueryRow("SELECT original FROM palavras WHERE id = 1")
	usuario.Scan(&a)
	fmt.Println(a)

}
