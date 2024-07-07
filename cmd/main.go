package main

import (
	"expense-application/internal/app"
	"expense-application/pkg/config"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Run(config.MustLoad())
}
