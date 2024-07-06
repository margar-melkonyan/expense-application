package service

import (
	"expense-application/internal/model"
	"expense-application/internal/repository"
)

type CategoryService struct {
	repository repository.Categories
}

func NewCategoryService(repository repository.Categories) *CategoryService {
	return &CategoryService{
		repository: repository,
	}
}

func (s *CategoryService) GetIncomeByCategory(category model.Category) ([]model.Budget, error) {

	return nil, nil
}

func (s *CategoryService) GetAll() ([]model.Category, error) {
	return s.repository.GetAll(), nil
}
