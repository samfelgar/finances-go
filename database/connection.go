package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

var db *gorm.DB

func Connect() *gorm.DB {
	if db != nil {
		return db
	}

	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_DATABASE"),
	)
	newConnection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err.Error())
		panic("error while connecting to database")
	}

	log.Println("connected to the database")

	db = newConnection
	return newConnection
}
