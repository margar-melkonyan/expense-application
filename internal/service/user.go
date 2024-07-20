package service

import (
	"expense-application/internal/model"
	"expense-application/internal/repository"
)

type UserService struct {
	repository repository.User
}

func NewUserService(repository repository.User) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) Get(user model.User) (model.User, error) {
	return s.repository.Get(user.Id)
}

func (s *UserService) Update(user *model.User, id uint) error {
	return s.repository.Update(user, id)
}
