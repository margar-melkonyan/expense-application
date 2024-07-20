package app

import (
	"expense-application/internal/db"
	"expense-application/internal/handler"
	"expense-application/internal/repository"
	"expense-application/internal/service"
	"expense-application/pkg/config"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func Run(config *config.Config) {
	postgresDB, err := db.NewPostgresDB(config)

	if err != nil {
		log.Fatalln(err)
	}

	repos := repository.NewRepository(postgresDB)
	services := service.NewService(repos)

	switch os.Getenv("SERVICE") {
	case "api":
		handlers := handler.NewHandler(services)
		srv := new(Server)

		if err := srv.run(config.HttpServer.Port, handlers); err != nil {
			slog.Error(fmt.Sprintf("Error occured while running http server: %v", err.Error()))
		}
	case "tg-bot":
		runBot(repos, services)
	}
}
