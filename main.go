package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samfelgar/finances-go/database"
	"github.com/samfelgar/finances-go/http"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.DB()
	database.Migrate(db)

	router := gin.Default()

	// Provide the database connection to the controllers within goroutines
	router.Use(func(context *gin.Context) {
		context.Set("database", db)
		context.Next()
	})
	router.Use(errorHandler)

	http.Routes(router)

	router.Run(fmt.Sprintf("%s:%s", os.Getenv("APP_URL"), os.Getenv("APP_PORT")))
}

func errorHandler(c *gin.Context) {
	c.Next()

	errors := c.Errors

	if len(errors) == 0 {
		return
	}

	var messages []string

	for _, err := range errors {
		messages = append(messages, err.Error())
	}

	c.JSON(-1, gin.H{
		"messages": messages,
	})
}
