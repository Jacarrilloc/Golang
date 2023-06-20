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

	//Router
	http.HandleFunc("/", Hola)
	http.HandleFunc("/page", PaginaNoFunciona)
	http.HandleFunc("/error", Error)
	http.HandleFunc("/saludar", Saludar)

	//Creacion de Servidor
	fmt.Println("Servidor corriendo en 8080")
	fmt.Println("Run Server: http://localhost:3000/")
	fmt.Println("\n")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
