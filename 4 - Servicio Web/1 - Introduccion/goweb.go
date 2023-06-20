package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler
func Hola(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hola Mundo")
}

func main() {

	//Router
	http.HandleFunc("/", Hola)

	//Creacion de Servidor
	fmt.Println("Servidor corriendo en 8080")
	fmt.Println("Run Server: http://localhost:8080/")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
