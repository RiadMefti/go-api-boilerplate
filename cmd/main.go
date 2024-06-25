package main

import (
	"log"
	"os"

	"github.com/RiadMefti/go-api-boilerplate/cmd/api"
	"github.com/RiadMefti/go-api-boilerplate/config"
	"github.com/RiadMefti/go-api-boilerplate/db"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	config := config.NewConfig(
		os.Getenv("DbHost"),
		os.Getenv("DbPort"),
		os.Getenv("DbUser"),
		os.Getenv("DbPassword"),
		os.Getenv("DbName"),
		"3000",
	)

	db, err := db.NewPostgresStore(config)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Init()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewApiServer(":3000", db)
	api.Run(server)
}
