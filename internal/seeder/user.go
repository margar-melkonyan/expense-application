package seeder

import (
	"expense-application/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

type UserSeeder struct {
	db *gorm.DB
}

func NewUserSeeder(db *gorm.DB) *UserSeeder {
	return &UserSeeder{
		db: db,
	}
}

func (s *UserSeeder) Seed() {
	s.CreateDefaultAdmin()
}

func (s *UserSeeder) CreateDefaultAdmin() {
	var admin model.User
	var role model.Role
	var userRole model.UserRole
	err := s.db.Table("users").Where("email = ?", os.Getenv("ADMIN_EMAIL")).First(&admin).Error

	if err != nil {
		slog.Error(err.Error())

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASSWORD")), bcrypt.DefaultCost)
		admin.Name = "Admin"
		admin.Email = os.Getenv("ADMIN_EMAIL")
		admin.Password = string(hashedPassword)
		err = s.db.Create(&admin).Error
		if err != nil {
			slog.Error(err.Error())
		}

		err = s.db.Table("roles").Where("title = ?", "admin").First(&role).Error
		userRole.UserID = admin.Id
		userRole.RoleID = role.Id

		err = s.db.Create(&userRole).Error

		if err != nil {
			slog.Error(err.Error())
		}
	}
}
