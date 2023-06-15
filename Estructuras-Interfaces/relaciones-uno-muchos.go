package main

import "fmt"

type User struct {
	nombre string
	email  string
	estado bool
}

// Relación uno a muchos
type Curso struct {
	titulo string
	videos []Video
}

type Video struct {
	titulo string
	curso  Curso
}

// Relación uno a uno
type Student struct {
	user   User
	codigo string
}

func main() {
	//Relacion uno a uno
	/*
		alex := User{
			nombre: "alex",
			email:  "Alex@gmail.com",
			estado: true,
		}

		roel := User{
			nombre: "Roel",
			email:  "roel@hotmail.com",
			estado: true,
		}

		alexStudent := Student{
			user:   alex,
			codigo: "0000111",
		}

		fmt.Println(roel)
		fmt.Println(alexStudent)
	*/

	//Relacion uno a muchos

	video1 := Video{
		titulo: "Introducción",
	}

	video2 := Video{
		titulo: "Prueba de Video 2",
	}

	curso := Curso{
		titulo: "Curso de Golang",
		videos: []Video{video1, video2},
	}

	video1.curso = curso
	video2.curso = curso

	fmt.Println(video1.curso.titulo)

	for , video := range curso.videos {
		fmt.Println(video.titulo)
	}
}
