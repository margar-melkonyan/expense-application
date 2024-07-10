package service

import (
	"expense-application/internal/model"
	"expense-application/internal/repository"
	"github.com/gosimple/slug"
)

type CategoryService struct {
	repository repository.Category
}

func NewCategoryService(repository repository.Category) *CategoryService {
	return &CategoryService{
		repository: repository,
	}
}

func (s *CategoryService) GetIncomeByCategory(category model.Category) ([]model.Budget, error) {
	return nil, nil
}

func (s *CategoryService) Store(category model.Category) (int, error) {
	category.Slug = slug.Make(category.Name)

	id, err := s.repository.Store(&category)

	if err != nil {
		return 0, err
	}
	return id, nil
}
