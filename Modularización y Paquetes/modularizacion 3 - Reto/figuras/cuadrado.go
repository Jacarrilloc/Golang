package figuras

type Cuadrado struct {
	Lado float64
}

func (cuadrado *Cuadrado) area() float64 {
	return cuadrado.Lado * cuadrado.Lado
}

func (cuadrado *Cuadrado) perimetro() float64 {
	return 4 * cuadrado.Lado
}
