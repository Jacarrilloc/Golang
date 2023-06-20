package main

import "fmt"

type User struct {
	nombre string
	email  string
	estado bool
}

// Relación uno a uno
type Student struct {
	user   User
	codigo string
}

func main() {
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
}
