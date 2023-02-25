package database

import (
	"fmt"
	baseMysqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"time"
)

var db *gorm.DB

func DB() *gorm.DB {
	if db == nil {
		log.Println("database connection already exists, returning")
		db = Connect()
	}

	return db
}

func Connect() *gorm.DB {
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))

	config := baseMysqlDriver.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", os.Getenv("DB_HOST"), port),
		DBName:               os.Getenv("DB_DATABASE"),
		ParseTime:            true,
		Loc:                  time.Local,
		AllowNativePasswords: true,
	}

	newConnection, err := gorm.Open(mysql.Open(config.FormatDSN()), &gorm.Config{})

	if err != nil {
		log.Println(err.Error())
		panic("error while connecting to database")
	}

	log.Println("connected to the database")

	db = newConnection
	return newConnection
}
