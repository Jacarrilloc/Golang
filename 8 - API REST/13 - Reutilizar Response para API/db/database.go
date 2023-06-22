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
	rows, error := Query(sql)
	if error != nil {
		fmt.Println("Error: ", error)
	}
	return rows.Next()
}

// Crea una Tabla
func CreateTable(schema string, name string) {
	if !ExistTable(name) {
		_, err := Exec(schema)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// Reiniciar el Registro de una Tabla
func TruncadeTable(tablename string) {
	sql := fmt.Sprintf("TRUNCATE %s", tablename)
	Exec(sql)
}

// Polimorfizar el Execute
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, error := db.Exec(query, args...)
	Close()
	if error != nil {
		fmt.Println(error)
	}

	return result, error
}

// Polimorfismo de Query
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, error := db.Query(query, args...)
	Close()
	if error != nil {
		fmt.Println(error)
	}

	return rows, error
}
