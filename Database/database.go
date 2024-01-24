package database

import (
	"fmt"
	"log"
	"pinaca/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseConnection() {

	dsn := "host=localhost user=postgres password=postgres dbname=book port=54323"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	DB.AutoMigrate(&model.Book{})
	fmt.Println("Database connection successful...")
}
