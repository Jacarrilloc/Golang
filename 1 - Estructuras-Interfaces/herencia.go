package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
}

type Empleado struct {
	Persona
	sueldo float64
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

	EMPLEADO2 := Empleado{
		sueldo: 5000,
	}

	EMPLEADO2.nombre = "Juan"
	EMPLEADO2.edad = 80

	EMPLEADO2.imprimir()
	fmt.Println(EMPLEADO2)

	empleado1 := Empleado{
		sueldo: 500,
	}
	empleado1.nombre = "Pedro"
	empleado1.edad = 30
	empleado1.imprimir()
	fmt.Println(empleado1)
}
