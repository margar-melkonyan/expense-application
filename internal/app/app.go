package app

import (
	"context"
	"expense-application/internal/db"
	"expense-application/internal/handler"
	"expense-application/internal/repository"
	"expense-application/internal/seeder"
	"expense-application/internal/service"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Panic("Error loading .env file")
	}
}

func Run() {
	postgresDB, err := db.NewPostgresDB()

	if err != nil {
		log.Fatalln(err)
	}

	repos := repository.NewRepository(postgresDB)
	seeders := seeder.NewSeeder(postgresDB)
	services := service.NewService(repos)
	seeders.Seed()

	switch os.Getenv("SERVICE") {
	case "api":
		handlers := handler.NewHandler(services)
		srv := new(Server)

		if err := srv.run(os.Getenv("SERVER_PORT"), handlers); err != nil {
			slog.Error(fmt.Sprintf("Error occured while running http server: %v", err.Error()))
		}

		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
		<-sc

		slog.Info("Expense application Shutting Down")

		if err := srv.shutdown(context.Background()); err != nil {
			slog.Info("error occured on server shutting down: %s", err.Error())
		}
	case "tg-bot":
		runBot(repos, services)
	}
}
