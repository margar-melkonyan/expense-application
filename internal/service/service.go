package service

import (
	"expense-application/internal/model"
	"expense-application/internal/repository"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Auth interface {
	SignIn(user *model.User) (map[string]string, error)
	SignUp(user *model.User) (map[string]string, error)
	RefreshToken(ctx *gin.Context)
}

type User interface {
	Get(user model.User) (model.User, error)
	Update(user *model.User, id uint) error
}

type Role interface {
	Role(roleID uint) (model.Role, error)
	Roles() (*[]model.Role, error)
	StoreRole(role *model.Role) error
	UpdateRole(role *model.Role, roleID uint) error
	DeleteRole(roleID uint) error
	AssignRole(userRole model.UserRole) error
	Permissions() map[string][]string
}

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
	UpdateHandler(bot *tgbotapi.BotAPI, message tgbotapi.Update) error
	SendMessage(bot *tgbotapi.BotAPI, message tgbotapi.MessageConfig, update tgbotapi.Update) error
	CreateKeyboard(commands []string, commandsPerRow int) [][]tgbotapi.KeyboardButton
}

type Report interface {
	GenDayReport(typeBudget string, userId uint) []byte
	GenWeekReport(typeBudget string, userId uint) []byte
	GenMonthReport(typeBudget string, userId uint) []byte
}

type Service struct {
	Auth
	User
	Role
	Budget
	Category
	Tg
	Reports map[string]Report
}

func NewService(repos *repository.Repository) *Service {
	reportServices := make(map[string]Report)
	reportServices["pdf"] = NewPdfService(repos.Budget)
	reportServices["xlsx"] = NewXLSXService(repos.Budget)

	return &Service{
		Auth:     NewAuthService(repos.User),
		User:     NewUserService(repos.User),
		Role:     NewRoleService(repos.Role),
		Budget:   NewBudgetService(repos.Budget),
		Category: NewCategoryService(repos.Category),
		Tg: NewTgService(
			repos.Category,
			repos.Budget,
			repos.User,
			reportServices,
		),
		Reports: reportServices,
	}
}
