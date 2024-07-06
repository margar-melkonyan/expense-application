package repository

import (
	"expense-application/internal/model"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repository CategoryRepository) GetAll() []model.Category {
	var categories []model.Category
	repository.db.Select("*").Find(&model.Category{}).Scan(&categories)

	return categories
}
