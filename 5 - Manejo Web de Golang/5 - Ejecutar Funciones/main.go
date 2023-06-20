package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Saludar(nombre string) string {
	return "Hola " + nombre + " desde una funcion"
}

func Index(rw http.ResponseWriter, r *http.Request) {

	funciones := template.FuncMap{
		"saludar": Saludar,
	}

	template, error := template.New("Index.html").Funcs(funciones).ParseFiles("index.html")

	if error != nil {
		panic(error)
	} else {
		template.Execute(rw, nil)
	}
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
