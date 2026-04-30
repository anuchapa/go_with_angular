package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gen"
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

	g := gen.NewGenerator(gen.Config{
		OutPath:      "./internal/adapter/db/query",
		ModelPkgPath: "./model",
		Mode: gen.WithoutContext|gen.WithDefaultQuery|gen.WithQueryInterface,
	})

	g.UseDB(db)
	g.ApplyBasic(
        g.GenerateModel("products"),
    )
	
	g.Execute()

	fmt.Println("Generated is successful.")
}
