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
	GetCategories() ([]model.Category, error)
	GetBySlug(slug string) (model.Category, error)
	GetByType(budgetType string) []model.Category
	GetCategoriesName(budgetType string) []string
	Store(category *model.Category) (uint, error)
	Update(category *model.Category) (uint, error)
	Delete(category *model.Category) (uint, error)
}

type Budget interface {
	Store(budget *model.Budget, category *model.Category) error
	GetBudgetByCategoryAndPeriod(budgetType string, userId uint, period string) ([]model.Category, error)
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
