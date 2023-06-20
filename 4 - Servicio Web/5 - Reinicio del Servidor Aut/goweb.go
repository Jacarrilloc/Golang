package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler
func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("El metodo es + " + r.Method)
	fmt.Fprintln(rw, "Hola Mundo")
}

func PaginaNoFunciona(rw http.ResponseWriter, r *http.Request) {
	http.NotFound(rw, r)
}

func Error(rw http.ResponseWriter, r *http.Request) {
	http.Error(rw, "Este es un Error", http.StatusNotFound)
	log.Println("Se ha enviado una respuesta de error 404")
}

func Saludar(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.URL.Query())

	name := r.URL.Query().Get("name")
	age := r.URL.Query().Get("age")

	fmt.Fprintf(rw, "Hola, %s tu edad es %s !!", name, age)
}

func main() {

	//Mux
	mux := http.NewServeMux()

	mux.HandleFunc("/", Hola)
	mux.HandleFunc("/page", PaginaNoFunciona)
	mux.HandleFunc("/error", Error)
	mux.HandleFunc("/saludar", Saludar)

	//Creacion de Servidor

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	fmt.Println("Servidor corriendo en 3000")
	fmt.Println("Run Server: http://localhost:3000/")
	log.Fatal(server.ListenAndServe())
}
