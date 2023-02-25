package main

import (
	"github.com/joho/godotenv"
	"github.com/samfelgar/finances-go/database"
	"log"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.Connect()
	database.Migrate(db)

}
