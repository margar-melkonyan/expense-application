package app

import (
	"expense-application/pkg/config"
)

func Run(config *config.Config) {
	//repos := repository.NewRepository()
	//services := service.NewService(repos)
	//handlers := handler.NewHandler(services)

	//srv := new(Server)
	//
	//if err := srv.run(config.HttpServer.Port, handlers); err != nil {
	//	slog.Error(fmt.Sprintf("Error occured while running http server: %v", err.Error()))
	//}

	runBot()
}
