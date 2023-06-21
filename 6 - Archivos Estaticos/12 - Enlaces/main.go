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

var templates = template.Must(template.New("T").ParseGlob("templates/**/*.html"))
var errorTemplate = template.Must(template.ParseFiles("templates/error/error.html"))

// Handler Error
func handlerError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	errorTemplate.Execute(rw, nil)
}

// Funcion de Render Template
func renderTemplate(rw http.ResponseWriter, name string, data interface{}) {

	error := templates.ExecuteTemplate(rw, name, data)

	if error != nil {
		handlerError(rw, http.StatusInternalServerError)
	}
}

// Handler
func Index(rw http.ResponseWriter, r *http.Request) {

	Usuario := Usuarios{"Julian", 23}

	renderTemplate(rw, "index.html", Usuario)
}

func Registro(rw http.ResponseWriter, r *http.Request) {
	renderTemplate(rw, "registro.html", nil)
}

func main() {
	//Configuración de Archivos Estaticos
	staticfile := http.FileServer(http.Dir("static"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)

	//Mux de Static Files
	mux.Handle("/static/", http.StripPrefix("/static/", staticfile))

	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}

	fmt.Println("Servidor corriendo en 3000")
	fmt.Println("Run Server: http://localhost:3000/")
	log.Fatal(server.ListenAndServe())
}
