package main

import "fmt"

type Calcular interface {
	area() float64
	perimetro()
}

type Circulo struct {
	radio int
}

type Cuadrado struct {
	lado int
}

func (cuadrado *Cuadrado) area() float64 {
	return float64(cuadrado.lado * cuadrado.lado)
}

func medidas(g Calcular) {
	fmt.Println(g)
	fmt.Println(g.area())
}

func main() {

	cuadrado := Cuadrado{
		ancho: 5,
		alto:  9,
	}

	medidas(&cuadrado)
}
