package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()

	//fmt.Println(db.ExistTable("users"))
	//db.CreateTable(models.UserSchema, "users")

	//fmt.Println(user)

	user := models.CreateUser("Alex", "Alex", "Alex@gmail.com")
	fmt.Println(user)

	//db.TruncadeTable("users")

	fmt.Println(models.ListUsers())

	//db.Ping()

	db.Close()
}
