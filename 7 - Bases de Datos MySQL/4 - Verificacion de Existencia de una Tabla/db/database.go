package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(localhost:3308)/database

const url = "root:admin@tcp(localhost:3308)/goweb_db"

// Guarda la Conexion
var db *sql.DB

// Funcion que abre la conexion con la base de datos
func Connect() {
	conection, error := sql.Open("mysql", url)
	if error != nil {
		panic(error)
	}

	fmt.Println("Conexion Exitosa ")
	db = conection
}

// Funcion que cierra la conexion
func Close() {
	db.Close()
}

// Funcion que verifica que este abierta la conexion
func Ping() {
	if error := db.Ping(); error != nil {
		panic(error)
	}
}

// Verificar la existencia de la tabla
func ExistTable(tablename string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tablename)
	rows, error := db.Query(sql)
	if error != nil {
		fmt.Println("Error: ", error)
	}
	return rows.Next()
}

// Crea una Tabla
func CreateTable(schema string, name string) {
	if !ExistTable(name) {
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}
}
