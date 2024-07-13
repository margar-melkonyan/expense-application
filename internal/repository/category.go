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

func (repository CategoryRepository) GetCategories() ([]model.Category, error) {
	var categories []model.Category

	err := repository.db.Model(&model.Category{}).Find(&categories).Error

	return categories, err
}

func (repository CategoryRepository) GetBySlug(slug string) (model.Category, error) {
	var category model.Category
	err := repository.db.Model(model.Category{}).Select("*").Where("slug = ?", slug).Find(&category).Error

	return category, err
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

func (repository CategoryRepository) Store(category *model.Category) (uint, error) {
	err := repository.db.Create(&category).Error

	return category.Id, err
}

func (repository CategoryRepository) Update(category *model.Category) (uint, error) {
	err := repository.db.Save(&category).Error

	return category.Id, err
}

func (repository CategoryRepository) Delete(category *model.Category) (uint, error) {
	err := repository.db.Delete(&category).Error

	return category.Id, err
}
