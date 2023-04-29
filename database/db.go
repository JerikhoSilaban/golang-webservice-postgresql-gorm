package database

import (
	"DTSGolang/Kelas2/Assignment8/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "1234"
	dbPort   = 5432
	dbName   = "assignment-gorm"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database: ", err.Error())
	}

	db.Debug().AutoMigrate(models.Author{}, models.Book{})
}

func GetDB() *gorm.DB {
	return db
}
