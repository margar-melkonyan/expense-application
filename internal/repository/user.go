package repository

import (
	"encoding/json"
	"errors"
	"expense-application/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func addDefaultUserRole(repository *UserRepository, user model.User) error {
	var role model.Role
	repository.db.Table("roles").Where("title = ?", "user").First(&role)
	return repository.db.Create(&model.UserRole{
		UserID: user.Id,
		RoleID: role.Id,
	}).Error
}

func (repository UserRepository) Get(id uint) (model.User, error) {
	var user model.User
	err := repository.db.Preload("Roles").Find(&user, "id = ?", id).Error
	_ = json.Unmarshal(user.Roles[0].Permissions, &user.Roles[0].PermissionsUnmarshalled)

	return user, err
}

func (repository UserRepository) GetByEmail(email string) (model.User, error) {
	var user model.User
	err := repository.db.Table("users").Where("email = ?", email).First(&user).Error
	return user, err
}

func (repository UserRepository) CurrentTgUser(tgId int64) (model.User, error) {
	var user model.User
	err := repository.db.Model(model.User{}).Where("tg_id = ?", tgId).First(&user).Error
	return user, err
}

func (repository UserRepository) CreateByTg(user *model.User) error {
	repository.db.Create(&user)
	return addDefaultUserRole(&repository, *user)
}

func (repository UserRepository) Create(user *model.User) (uint, error) {
	var foundedUser model.User
	repository.db.Table("users").Where("email", user.Email).First(&foundedUser)

	if foundedUser.Email != "" {
		return 0, errors.New("test")
	}

	err := repository.db.Create(&user).Error
	err = addDefaultUserRole(&repository, *user)
	return user.Id, err
}

func (repository UserRepository) Update(user *model.User, id uint) error {
	if user.Password != "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hash)
	}
	return repository.db.Table("users").Where("id = ?", id).Updates(&user).Error
}
