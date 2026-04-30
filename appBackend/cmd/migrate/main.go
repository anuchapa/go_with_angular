package main

import (
	"fmt"
	"goBackend/cmd/migrate/migration"
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
	masterDb, err := gorm.Open(sqlserver.Open(os.Getenv("sqlSerMasterConnectionString")), &gorm.Config{})
	if err != nil {
		panic("Master database can't contected.")
	}

	masterDb.Exec("IF NOT EXISTS (SELECT * FROM sys.databases WHERE name = 'myapp') CREATE DATABASE myapp")

	db, err := gorm.Open(sqlserver.Open(os.Getenv("sqlSerConnectionString")), &gorm.Config{})

	if err != nil {
		panic("Database can't contected.")
	}

	db.AutoMigrate(&migration.Product{})
	fmt.Println("Migration is successful.")
}
