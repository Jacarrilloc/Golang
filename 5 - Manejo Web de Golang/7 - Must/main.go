package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Usuarios struct {
	UserName string
	Edad     int
}

func Saludar(nombre string) string {
	return "Hola " + nombre + " desde una funcion"
}

func Index(rw http.ResponseWriter, r *http.Request) {

	template := template.Must(template.New("index.html").ParseFiles("index.html", "base.html"))

	Usuario := Usuarios{"Julian", 23}

	template.Execute(rw, Usuario)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	fmt.Println("Servidor corriendo en 3000")
	fmt.Println("Run Server: http://localhost:3000/")
	log.Fatal(server.ListenAndServe())
}
