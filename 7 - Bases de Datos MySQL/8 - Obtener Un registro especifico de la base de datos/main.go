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

	//db.TruncadeTable("users")

	//user := models.CreateUser("Pedro", "Pedro", "Pedro@gmail.com")
	//fmt.Println(user)

	users := models.GetUser(1)
	fmt.Println(users)

	//db.Ping()

	db.Close()
}
