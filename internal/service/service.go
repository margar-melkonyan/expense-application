package service

import (
	"expense-application/internal/model"
	"expense-application/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Category interface {
	GetIncomeByCategory(category model.Category) ([]model.Budget, error)
	GetAll() []model.Category
	Store(category model.Category) (int, error)
}

type Tg interface {
	CommandHandler(bot *tgbotapi.BotAPI, message tgbotapi.Update) error
	SendMessage(bot *tgbotapi.BotAPI, message tgbotapi.MessageConfig, update tgbotapi.Update) error
}

type Service struct {
	Category
	Tg
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Category: NewCategoryService(repos.Category),
	}
}
