package service

import (
	"bytes"
	"expense-application/internal/model"
	"expense-application/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/johnfercher/maroto/v2/pkg/core"
)

type Category interface {
	IndexCategories() ([]model.Category, error)
	GetCategoryBySlug(slug string) (model.Category, error)
	GetIncomeByCategory(category model.Category) ([]model.Budget, error)
	Store(category model.Category) (uint, error)
	Update(slug string, category model.Category) (uint, error)
	Delete(slug string) (uint, error)
}

type Budget interface {
	GetBudget(id uint) (*model.Budget, error)
	GetUserBudgets(userID uint) ([]model.Budget, error)
	Store(budget model.Budget, category model.Category) error
	Update(budget model.Budget) (uint, error)
	Delete(id uint) (uint, error)
}

type Tg interface {
	CommandHandler(bot *tgbotapi.BotAPI, message tgbotapi.Update) error
	SendMessage(bot *tgbotapi.BotAPI, message tgbotapi.MessageConfig, update tgbotapi.Update) error
	CreateKeyboard(commands []string, commandsPerRow int) [][]tgbotapi.KeyboardButton
}

type PDF interface {
	GenDayReport(typeBudget string, userId uint) core.Document
	GenWeekReport(typeBudget string, userId uint) core.Document
	GenMonthReport(typeBudget string, userId uint) core.Document
}

type XLSX interface {
	GenDayReport(typeBudget string, userId uint) *bytes.Buffer
	GenWeekReport(typeBudget string, userId uint) *bytes.Buffer
	GenMonthReport(typeBudget string, userId uint) *bytes.Buffer
}

type Service struct {
	Category
	Budget
	Tg
	PDF
	XLSX
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Category: NewCategoryService(repos.Category),
		Budget:   NewBudgetService(repos.Budget),
		Tg: NewTgService(
			repos.Category,
			repos.Budget,
			repos.User,
			NewPdfService(repos.Budget),
			NewXLSXService(repos.Budget),
		),
		XLSX: NewXLSXService(repos.Budget),
		PDF:  NewPdfService(repos.Budget),
	}
}
