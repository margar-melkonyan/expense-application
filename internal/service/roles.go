package service

import (
	"expense-application/internal/model"
	"expense-application/internal/repository"
)

type RoleService struct {
	repository repository.Role
}

func (s *RoleService) AssignRole(userRole model.UserRole) error {
	return s.repository.AssignRole(userRole)
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

func (s *RoleService) Permissions() map[string][]string {
	return map[string][]string{
		"roles": {
			"roles_create",
			"roles_read",
			"roles_update",
			"roles_delete",
		},
		"budgets": {
			"budgets_create",
			"budgets_read",
			"budgets_update",
			"budgets_delete",
		},
		"categories": {
			"categories_create",
			"categories_read",
			"categories_update",
			"categories_delete",
		},
		"users": {
			"users_read",
			"users_update",
		},
	}
}
