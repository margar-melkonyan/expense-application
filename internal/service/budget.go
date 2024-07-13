package service

import (
	"expense-application/internal/model"
	"expense-application/internal/repository"
)

type BudgetService struct {
	budgetRepository repository.Budget
}

func NewBudgetService(budgetRepository repository.Budget) *BudgetService {
	return &BudgetService{
		budgetRepository: budgetRepository,
	}
}

func (s BudgetService) GetUserBudgets(userID uint) ([]model.Budget, error) {
	return s.budgetRepository.GetUserBudget(userID)
}

func (s BudgetService) Store(budget model.Budget, category model.Category) error {
	return s.budgetRepository.Store(&budget, &category)
}

func (s *BudgetService) Delete(userId uint) (uint, error) {
	return 0, nil
}
