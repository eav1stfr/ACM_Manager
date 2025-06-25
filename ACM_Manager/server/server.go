package main

import (
	"acmmanager/internal/models"
	"acmmanager/internal/sqlconnect"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db, err := sqlconnect.ConnectDb()

	if err != nil {
		panic(err)
	}
	defer db.Close()

	var members []models.Member
	err = db.Select(&members, "SELECT * FROM members")
	if err != nil {
		panic(err)
	}
	fmt.Println(members)
}
