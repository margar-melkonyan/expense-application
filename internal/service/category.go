package service

import (
	"expense-application/internal/model"
	"expense-application/internal/repository"
	sluger "github.com/gosimple/slug"
)

type CategoryService struct {
	repository repository.Category
}

func NewCategoryService(repository repository.Category) *CategoryService {
	return &CategoryService{
		repository: repository,
	}
}

func (s *CategoryService) IndexCategories() ([]model.Category, error) {
	return s.repository.GetCategories()
}

func (s *CategoryService) GetCategoryBySlug(slug string) (model.Category, error) {
	return s.repository.GetBySlug(slug)
}

func (s *CategoryService) GetIncomeByCategory(category model.Category) ([]model.Budget, error) {
	return nil, nil
}

func (s *CategoryService) Store(category model.Category) (int, error) {
	category.Slug = sluger.Make(category.Name)

	id, err := s.repository.Store(&category)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *CategoryService) Update(slug string, category model.Category) (int, error) {
	oldCategory, _ := s.repository.GetBySlug(slug)
	oldCategory.Type = category.Type
	oldCategory.Name = category.Name
	oldCategory.Slug = sluger.Make(category.Name)

	return s.repository.Update(&oldCategory)
}

func (s *CategoryService) Delete(slug string) (int, error) {
	category, err := s.repository.GetBySlug(slug)

	if err != nil {
		return 0, err
	}

	return s.repository.Delete(&category)
}
