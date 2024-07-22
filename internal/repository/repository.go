package repository

import (
	"expense-application/internal/model"
	"gorm.io/gorm"
)

type User interface {
	Get(id uint) (model.User, error)
	GetByEmail(email string) (model.User, error)
	CurrentTgUser(tgId int64) (model.User, error)
	CreateByTg(user *model.User) error
	Create(user *model.User) (uint, error)
	Update(user *model.User, id uint) error
}

type Role interface {
	Role(id uint) (model.Role, error)
	Roles() (*[]model.Role, error)
	StoreRole(role *model.Role) error
	UpdateRole(role *model.Role, id uint) error
	DeleteRole(id uint) error
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
	GetBudget(id uint) (*model.Budget, error)
	GetUserBudget(userId uint) ([]model.Budget, error)
	Store(budget *model.Budget, category *model.Category) error
	GetBudgetByCategoryAndPeriod(budgetType string, userId uint, period string) ([]model.Category, error)
	Update(budget *model.Budget) (uint, error)
	Delete(id uint) (uint, error)
}

type Repository struct {
	Category
	Budget
	User
	Role
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Category: NewCategoryRepository(db),
		User:     NewUserRepository(db),
		Budget:   NewBudgetRepository(db),
		Role:     NewRoleRepository(db),
	}
}
