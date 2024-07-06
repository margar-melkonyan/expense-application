package repository

import (
	"expense-application/internal/model"
	"gorm.io/gorm"
)

type Categories interface {
	GetAll() []model.Category
}

type Repository struct {
	Categories
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Categories: NewCategoryRepository(db),
	}
}
