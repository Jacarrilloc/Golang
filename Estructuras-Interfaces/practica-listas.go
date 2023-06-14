package main

import "fmt"

type TaskList struct {
	tasks []*Task
}

func (tl *TaskList) appendTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) removeTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

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

	lista := TaskList{}
	lista.appendTask(&t1)
	lista.appendTask(&t2)

	fmt.Println(lista)

	t3 := Task{
		name:        "Info3",
		descripcion: "Este es un elemento del item",
		complete:    false,
	}

	lista.appendTask(&t3)
	fmt.Println(lista)

	lista.removeTask(1)

	for i, task := range lista.tasks {
		fmt.Println(i, task.name)
	}

	//t1.toPrint()
	//t2.toPrint()
}
