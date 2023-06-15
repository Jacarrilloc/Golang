package main

type Calcular interface {
	area() float64
	perimetro()
}

type Circulo struct {
	radio int
}

type Cuadrado struct {
	ancho int
	alto  int
}

func (*Cuadrado) area() float64 {
	resultado := cuadrado.alto * cuadrado.ancho
	return float64(resultado)
}

func main() {

	cuadrado := Cuadrado{
		ancho: 5,
		alto:  9,
	}

	resultado := area(&cuadrado)

}
