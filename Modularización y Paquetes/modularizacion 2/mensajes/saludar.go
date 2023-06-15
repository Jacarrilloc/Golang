package mensajes

import "fmt"

func Saludar() {
	fmt.Println("Hola")
}

const mensaje = "Hola desde la constante"

func funcionPrivada() {
	fmt.Println("Hola desde la funcion Privada")
}

func Imprimir() {
	fmt.Println(mensaje)
	funcionPrivada()
}
