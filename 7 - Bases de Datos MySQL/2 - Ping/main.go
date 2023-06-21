package main

import (
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()

	db.CreateTable(models.UserSchema)

	//db.Ping()

	db.Close()
}
