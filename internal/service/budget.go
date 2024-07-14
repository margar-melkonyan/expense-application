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

func (s *BudgetService) GetBudget(id uint) (*model.Budget, error) {
	return s.budgetRepository.GetBudget(id)
}

func (s *BudgetService) GetUserBudgets(userID uint) ([]model.Budget, error) {
	rawBudgets, err := s.budgetRepository.GetUserBudget(userID)
	var budgets []model.Budget

	for _, rawBudget := range rawBudgets {
		rawBudget.Amount /= 100
		budgets = append(budgets, rawBudget)
	}

	return budgets, err
}

func (s *BudgetService) Store(budget model.Budget, category model.Category) error {
	budget.Amount *= 100
	return s.budgetRepository.Store(&budget, &category)
}

func (s *BudgetService) Update(budget model.Budget) (uint, error) {
	budget.Amount *= 100
	return s.budgetRepository.Update(&budget)
}

func (s *BudgetService) Delete(id uint) (uint, error) {
	return s.budgetRepository.Delete(id)
}
