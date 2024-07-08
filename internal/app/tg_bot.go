package app

import (
	"expense-application/internal/repository"
	"expense-application/internal/service"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"log/slog"
	"os"
)

func runBot(repositories *repository.Repository) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	slog.Info(fmt.Sprintf("Authorized on account %s", bot.Self.UserName))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if err := service.NewTgService(
			repositories.Category,
			repositories.Budget,
			repositories.User,
		).CommandHandler(bot, update); err != nil {
			slog.Error(err.Error())
		}
	}
}
