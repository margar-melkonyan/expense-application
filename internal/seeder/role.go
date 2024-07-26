package seeder

import (
	"encoding/json"
	"expense-application/internal/model"
	"gorm.io/gorm"
)

type RoleSeeder struct {
	db *gorm.DB
}

func NewRoleSeeder(db *gorm.DB) *RoleSeeder {
	return &RoleSeeder{
		db: db,
	}
}

func (s *RoleSeeder) Seed() {
	s.createDefaultRoles()
}

func (s *RoleSeeder) createDefaultRoles() {
	var roles []model.Role
	userPermissions, _ := json.Marshal([]string{
		"categories_read",
		"budgets_create",
		"budgets_read",
		"budgets_update",
		"budgets_delete",
		"users_read",
		"users_update",
	})

	adminPermissions, _ := json.Marshal([]string{
		"categories_create",
		"categories_read",
		"categories_update",
		"categories_delete",
		"budgets_create",
		"budgets_read",
		"budgets_update",
		"budgets_delete",
		"users_read",
		"users_update",
		"roles_create",
		"roles_read",
		"roles_update",
		"roles_delete",
	})

	roles = append(roles, model.Role{
		Title:        "user",
		DisplayTitle: "User",
		Permissions:  userPermissions,
	})

	roles = append(roles, model.Role{
		Title:        "admin",
		DisplayTitle: "Admin",
		Permissions:  adminPermissions,
	})

	for _, role := range roles {
		if s.db.
			Table("roles").
			Where("display_title = ?", role.DisplayTitle).
			First(&role).Error != nil {
			s.db.Create(&role)
		}
	}
}
