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

func (repository CategoryRepository) Store(category *model.Category) (int, error) {
	err := repository.db.Create(&category).Error

	return category.Id, err
}

func (repository CategoryRepository) GetByType(budgetType string) []model.Category {
	var categories []model.Category
	repository.db.Model(model.Category{}).Select("*").Where("type = ?", budgetType).Find(&categories)

	return categories
}

func (repository CategoryRepository) GetCategoriesName(budgetType string) []string {
	var categoriesName []string
	repository.db.Model(model.Category{}).Select("name").Where("type = ?", budgetType).Find(&categoriesName)

	return categoriesName
}
