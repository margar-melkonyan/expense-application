package service

import (
	"bytes"
	"expense-application/internal/repository"
)

type XLSXService struct {
	budgetRepository repository.Budget
}

func NewXLSXService(budgetRepository repository.Budget) *XLSXService {
	return &XLSXService{
		budgetRepository: budgetRepository,
	}
}

func (s XLSXService) GenDayReport(typeBudget string, userId int) *bytes.Buffer {
	return &bytes.Buffer{}
}
func (s XLSXService) GenWeekReport(typeBudget string, userId int) *bytes.Buffer {
	return &bytes.Buffer{}
}
func (s XLSXService) GenMonthReport(typeBudget string, userId int) *bytes.Buffer {
	return &bytes.Buffer{}
}
