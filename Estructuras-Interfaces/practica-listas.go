package main

import "fmt"

type Task struct {
	name        string
	descripcion string
	complete    bool
}

func (t *Task) toPrint() {
	fmt.Printf("Nombre: %s\nDescripcion: %s\nCompletado: %t\n\n", t.name, t.descripcion, t.complete)
}

func (t *Task) markComplete() {
	t.complete = true
}

func main() {

	t1 := Task{
		name:        "Prueba",
		descripcion: "Prueba para la verificación de los Datos",
		complete:    false,
	}

	t2 := Task{
		name:        "Información",
		descripcion: "Prueba para la verificación de los Datos en segunda parte",
		complete:    false,
	}

	t1.toPrint()
	t2.toPrint()
}
