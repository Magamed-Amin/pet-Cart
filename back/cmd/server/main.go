package main

import (
	"fmt"
	"log"

	"awesomeProject/internal/database"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Connected to database")
}
