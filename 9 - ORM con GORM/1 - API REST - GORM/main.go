package main

import (
	"log"
	"net/http"

	"gorm/handler"

	"github.com/gorilla/mux"
)

func main() {
	//Rutas

	mux := mux.NewRouter()

	//Endpoints

	mux.HandleFunc("/api/user/", handler.GetUsers).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handler.GetUser).Methods("GET")
	mux.HandleFunc("/api/user/", handler.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handler.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handler.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", mux))
}
