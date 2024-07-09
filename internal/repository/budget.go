package repository

import (
	"expense-application/internal/model"
	"fmt"
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

func (repository BudgetRepository) GetBudgetByCategoryAndPeriod(budgetType string, userId int, period string) ([]model.Category, error) {
	var budgetsByCategory []model.Category

	err := repository.db.Model(&model.Category{}).
		Where("type = ?", budgetType).
		Preload("Budgets", func(db *gorm.DB) *gorm.DB {
			return db.
				Order("budgets.created_at").
				Where(fmt.Sprintf("budgets.user_id = ? AND DATE_TRUNC('%s', budgets.created_at) = DATE_TRUNC('%s', CURRENT_DATE)", period, period), userId)
		}).
		Find(&budgetsByCategory).Error

	return budgetsByCategory, err
}

func (repository BudgetRepository) Create(budget *model.Budget, category *model.Category) error {
	repository.db.Create(&budget)

	return repository.db.Create(&model.BudgetCategory{
		BudgetID:   budget.Id,
		CategoryID: category.Id,
	}).Error
}
