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
	Activo   bool
	Admin    bool
	Cursos   []Curso
}

type Curso struct {
	Nombre string
}

func Index(rw http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(rw, "Hola Mundo")
	template, error := template.ParseFiles("Index.html")

	c1 := Curso{"Go"}
	c2 := Curso{"Python"}
	c3 := Curso{"Java"}
	c4 := Curso{"C++"}

	cursos := []Curso{c1, c2, c3, c4}

	usuario := Usuarios{"Julian", 20, true, false, cursos}

	if error != nil {
		panic(error)
	} else {
		template.Execute(rw, usuario)
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
