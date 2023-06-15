package main

import "paquetes/figuras"

func main() {
	cuadrado := figuras.Cuadrado{Lado: 8}
	figuras.Medidas(&cuadrado)

	circulo := figuras.Circulo{Radio: 15}
	figuras.Medidas(&circulo)
}
