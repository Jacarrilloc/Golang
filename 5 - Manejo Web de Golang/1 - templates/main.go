package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Index(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Hola Mundo")
	template, error := template.ParseFiles("Index.html")

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
