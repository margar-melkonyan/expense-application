package repository

import (
	"expense-application/internal/model"
	"gorm.io/gorm"
)

type BudgetRepository struct {
	db *gorm.DB
}

func NewBudgetRepository(db *gorm.DB) *BudgetRepository {
	return &BudgetRepository{
		db: db,
	}
}

func (repository BudgetRepository) GetToday(userId int) ([]model.Budget, error) {
	var budgets []model.Budget
	err := repository.db.Model(&model.Budget{}).Where("user_id = ?", userId).Scan(&budgets).Error

	return budgets, err
}

func (repository BudgetRepository) Create(budget *model.Budget, category *model.Category) error {
	repository.db.Create(&budget)

	return repository.db.Create(&model.BudgetCategory{
		BudgetID:   budget.Id,
		CategoryID: category.Id,
	}).Error
}
