package repository

import (
	"encoding/json"
	"expense-application/internal/model"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (repository *RoleRepository) Role(id uint) (model.Role, error) {
	var role model.Role
	err := repository.db.First(&role).Where("id = ?", id).Error
	_ = json.Unmarshal(role.Permissions, &role.PermissionsUnmarshalled)

	return role, err
}

func (repository *RoleRepository) Roles() (*[]model.Role, error) {
	var roles []model.Role
	err := repository.db.Find(&roles).Error

	if err != nil {
		return nil, err
	}

	for i, role := range roles {
		if err := json.Unmarshal(role.Permissions, &roles[i].PermissionsUnmarshalled); err != nil {
			return nil, err
		}
	}

	return &roles, nil
}

func (repository *RoleRepository) StoreRole(role *model.Role) error {
	role.Title = slug.Make(role.DisplayTitle)
	role.Permissions, _ = json.Marshal(role.PermissionsUnmarshalled)

	return repository.db.Save(role).Error
}

func (repository *RoleRepository) UpdateRole(role *model.Role, id uint) error {
	role.Title = slug.Make(role.DisplayTitle)
	role.Permissions, _ = json.Marshal(role.PermissionsUnmarshalled)

	return repository.db.Table("roles").Where("id = ?", id).Updates(&role).Error
}

func (repository *RoleRepository) DeleteRole(id uint) error {
	return repository.db.Delete(&model.Role{}, id).Error
}

func (repository *RoleRepository) AssignRole(userRole model.UserRole) error {
	var currentUserRole model.UserRole
	currentUserRole.UserID = userRole.UserID

	repository.db.Table("user_roles").
		Where("user_id = ?", userRole.UserID).
		First(&currentUserRole).
		Delete(&currentUserRole)

	return repository.db.Save(&userRole).Error
}
