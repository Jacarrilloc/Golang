package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
}

func (p *Persona) editNombre(nombre string) {
	p.nombre = nombre
}

func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s\nEdad: %d\n", p.nombre, p.edad)
}

func main() {
	p1 := Persona{"Alex", 26}

	//fmt.Println(p1)
	//p1.nombre = "Juan"
	p1.editNombre("Julian")
	p1.imprimir()

	//fmt.Println(p1)

	p2 := Persona{
		nombre: "Roel",
		edad:   27,
	}
	p2.imprimir()
	p2.editNombre("Carrillo")

	//fmt.Println(p2)
}
