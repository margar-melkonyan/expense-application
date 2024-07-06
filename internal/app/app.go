package app

import (
	"expense-application/internal/handler"
	"expense-application/internal/repository"
	"expense-application/internal/service"
	"expense-application/pkg/config"
	"fmt"
	"log/slog"
)

func Run(config *config.Config) {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(Server)

	if err := srv.run(config.HttpServer.Port, handlers); err != nil {
		slog.Error(fmt.Sprintf("Error occured while running http server: %v", err.Error()))
	}
}
