package handler

import (
	"fmt"
	"net/http"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Listar Todos los Usuarios")
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
