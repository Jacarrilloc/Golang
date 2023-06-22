package handler

import (
	"apirest/db"
	"apirest/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")

	db.Connect()
	users := models.ListUsers()
	db.Close()
	output, _ := json.Marshal(users)
	fmt.Fprintln(rw, string(output))
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Obtener un Usuario")
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Crear un Usuario")
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Editar un Usuario")
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Eliminar un Usuario")
}
