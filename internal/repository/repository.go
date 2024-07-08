package repository

import (
	"expense-application/internal/model"
	"gorm.io/gorm"
)

type User interface {
	CurrentUser() (model.User, error)
	CurrentTgUser(tgId int64) (model.User, error)
	SignUpByTg(user model.User) error
	SignUp(user model.User) (model.User, error)
	SignIn(user model.User) (model.User, error)
}

type Category interface {
	GetAll() []model.Category
	GetCategoriesName() []string
	Store(category model.Category) (int, error)
}

type Repository struct {
	Category
	User
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Category: NewCategoryRepository(db),
		User:     NewUserRepository(db),
	}
}
