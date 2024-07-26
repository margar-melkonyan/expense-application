package app

import (
	"context"
	"expense-application/internal/db"
	"expense-application/internal/handler"
	"expense-application/internal/repository"
	"expense-application/internal/seeder"
	"expense-application/internal/service"
	"expense-application/pkg/config"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func Run(config *config.Config) {
	postgresDB, err := db.NewPostgresDB(config)

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

		if err := srv.run(config.HttpServer.Port, handlers); err != nil {
			slog.Error(fmt.Sprintf("Error occured while running http server: %v", err.Error()))
		}

		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
		<-sc

		slog.Info("TodoApp Shutting Down")

		if err := srv.shutdown(context.Background()); err != nil {
			slog.Info("error occured on server shutting down: %s", err.Error())
		}
	case "tg-bot":
		runBot(repos, services)
	}
}
