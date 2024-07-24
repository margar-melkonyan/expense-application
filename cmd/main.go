package main

import (
	"expense-application/internal/app"
	"expense-application/pkg/config"
)

// @title Expense Application
// @version 1.0
// @description API for income and expense applications that allows you to receive a report for a certain period

// @host localhost:8080/api
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	app.Run(config.MustLoad())
}
