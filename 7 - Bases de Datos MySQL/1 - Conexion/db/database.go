package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(localhost:3308)/database

const url = "root:admin@tcp(localhost:3308)/goweb_db"

var db *sql.DB

func Connect() {
	conection, error := sql.Open("mysql", url)
	if error != nil {
		panic(error)
	}

	fmt.Println("Conexion Exitosa ")
	db = conection
}

func Close() {
	db.Close()
}
