package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()

	fmt.Println(db.ExistTable("users"))
	db.CreateTable(models.UserSchema, "users")

	//db.Ping()

	db.Close()
}
