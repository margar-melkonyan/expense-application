package main

import (
	"expense-application/internal/app"
	"expense-application/pkg/config"
)

func main() {
	app.Run(config.MustLoad())
}
