package service

import repository "expense-application/internal/repository"

type Categories interface {
}

type Service struct {
	Categories
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
