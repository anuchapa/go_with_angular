package main

import (
	"goBackend/internal/adapter/db/model"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := gorm.Open(sqlserver.Open(os.Getenv("sqlSerConnectionString")), &gorm.Config{})

	if err != nil {
		panic("Database can't contected.")
	}

	db.AutoMigrate(&model.Product{})
}
