package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	
	dns := os.Getenv("sqlSerConnectionString")
	if dns == "" {
		panic("sqlSerConnectionString not found.")
	}

	config := config{
		addr: ":3000",
		db:   dns,
	}

	app := applicaiton{config: config}
	if err := app.run(app.mount()); err != nil {
		log.Panicf("Server has failed to start, err: %s.", err.Error())
	}
}
