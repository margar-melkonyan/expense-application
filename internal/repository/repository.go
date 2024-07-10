package repository

import (
	"expense-application/internal/model"
	"gorm.io/gorm"
)

type User interface {
	CurrentUser() (model.User, error)
	CurrentTgUser(tgId int64) (model.User, error)
	SignUpByTg(user *model.User) error
	SignUp(user *model.User) (model.User, error)
	SignIn(user *model.User) (model.User, error)
}

type Category interface {
	GetByType(budgetType string) []model.Category
	GetCategoriesName(budgetType string) []string
	GetByName(categoryName string) (model.Category, error)
	Store(category *model.Category) (int, error)
}

type Budget interface {
	Create(budget *model.Budget, category *model.Category) error
	GetBudgetByCategoryAndPeriod(budgetType string, userId int, period string) ([]model.Category, error)
}

type Repository struct {
	Category
	User
	Budget
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Category: NewCategoryRepository(db),
		User:     NewUserRepository(db),
		Budget:   NewBudgetRepository(db),
	}
}
