package handler

import (
	"apirest/db"
	"apirest/models"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v3"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {

	db.Connect()
	users := models.ListUsers()
	db.Close()
	output, _ := yaml.Marshal(users)
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
