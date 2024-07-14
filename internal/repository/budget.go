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

func (repository BudgetRepository) GetBudget(id uint) (*model.Budget, error) {
	var budget model.Budget
	err := repository.db.Where("id = ?", id).Find(&budget).Error

	return &budget, err
}

func (repository BudgetRepository) GetUserBudget(userId uint) ([]model.Budget, error) {
	var budgets []model.Budget
	err := repository.db.Where("user_id = ?", userId).Find(&budgets).Error

	return budgets, err
}

func (repository BudgetRepository) GetBudgetByCategoryAndPeriod(budgetType string, userId uint, period string) ([]model.Category, error) {
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

func (repository BudgetRepository) Store(budget *model.Budget, category *model.Category) error {
	repository.db.Create(&budget)

	return repository.db.Create(&model.BudgetCategory{
		BudgetID:   budget.Id,
		CategoryID: category.Id,
	}).Error
}

func (repository BudgetRepository) Update(budget *model.Budget) (uint, error) {
	err := repository.db.Save(&budget).Error
	return budget.Id, err
}

func (repository BudgetRepository) Delete(id uint) (uint, error) {
	var budget model.Budget
	repository.db.Find(&model.Budget{}).Where("id = ?", id).First(&budget)

	return id, repository.db.Delete(&budget).Error
}
