package handler

import (
	"apirest/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {

	if users, error := models.ListUsers(); error != nil {
		models.SendNotFound(rw)
	} else {
		models.SendData(rw, users)
	}
}

func GetUser(rw http.ResponseWriter, r *http.Request) {

	if user, error := getUserByRequest(r); error != nil {
		models.SendNotFound(rw)
	} else {
		models.SendData(rw, user)
	}
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	//Obtener Registro
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if error := decoder.Decode(&user); error != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		user.Save()
		models.SendData(rw, user)
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {

	//Obtener Registro
	var userId int64

	if user, error := getUserByRequest(r); error != nil {
		models.SendNotFound(rw)
	} else {
		userId = user.Id
	}

	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if error := decoder.Decode(&user); error != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		user.Id = userId
		user.Save()
		models.SendData(rw, user)
	}
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	if user, error := getUserByRequest(r); error != nil {
		models.SendNotFound(rw)
	} else {
		user.Delete()
		models.SendData(rw, user)
	}
}

func getUserByRequest(r *http.Request) (models.User, error) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	if user, error := models.GetUser(userId); error != nil {
		return *user, error
	} else {
		return *user, nil
	}
}
