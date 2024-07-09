package service

import (
	"expense-application/internal/model"
	"expense-application/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Category interface {
	GetIncomeByCategory(category model.Category) ([]model.Budget, error)
	Store(category model.Category) (int, error)
}

type Tg interface {
	CommandHandler(bot *tgbotapi.BotAPI, message tgbotapi.Update) error
	SendMessage(bot *tgbotapi.BotAPI, message tgbotapi.MessageConfig, update tgbotapi.Update) error
	CreateKeyboard(commands []string, commandsPerRow int) [][]tgbotapi.KeyboardButton
}

type PDF interface {
	GenDayReport(typeBudget string)
	GenWeekReport(typeBudget string)
	GenMonthReport(typeBudget string)
	GenYearReport(typeBudget string)
}

type Service struct {
	Category
	Tg
	PDF
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Category: NewCategoryService(repos.Category),
		Tg:       NewTgService(repos.Category, repos.Budget, repos.User),
		PDF:      NewPdfService(),
	}
}
