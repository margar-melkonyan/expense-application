package repository

import (
	"expense-application/internal/model"
	"gorm.io/gorm"
)

type Category interface {
	GetAll() []model.Category
	Store(category model.Category) (int, error)
}

type Repository struct {
	Category
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Category: NewCategoryRepository(db),
	}
}
