package service

import (
	"expense-application/internal/model"
	"expense-application/internal/repository"
)

type RoleService struct {
	repository repository.Role
}

func NewRoleService(repository repository.Role) *RoleService {
	return &RoleService{
		repository: repository,
	}
}

func (s *RoleService) Role(roleID uint) (model.Role, error) {
	return s.repository.Role(roleID)
}

func (s *RoleService) Roles() (*[]model.Role, error) {
	return s.repository.Roles()
}

func (s *RoleService) StoreRole(role *model.Role) error {
	return s.repository.StoreRole(role)
}

func (s *RoleService) UpdateRole(role *model.Role, roleID uint) error {
	return s.repository.UpdateRole(role, roleID)
}

func (s *RoleService) DeleteRole(roleID uint) error {
	return s.repository.DeleteRole(roleID)
}
