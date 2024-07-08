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

func (repository CategoryRepository) Store(category model.Category) (int, error) {
	err := repository.db.Create(&category).Error

	return category.Id, err
}

func (repository CategoryRepository) GetAll() []model.Category {
	var categories []model.Category
	repository.db.Model(model.Category{}).Select("*").Find(&categories)

	return categories
}

func (repository CategoryRepository) GetCategoriesName() []string {
	var categoriesName []string
	repository.db.Model(model.Category{}).Select("name").Find(&categoriesName)

	return categoriesName
}
