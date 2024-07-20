package repository

import (
	"errors"
	"expense-application/internal/model"
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

func (repository UserRepository) Get(id uint) (model.User, error) {
	var user model.User
	err := repository.db.Find(&user, "id = ?", id).Error
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
	err := repository.db.Create(&user).Error
	return err
}

func (repository UserRepository) Create(user *model.User) (uint, error) {
	var foundedUser model.User
	repository.db.Table("users").Where("email", user.Email).First(&foundedUser)

	if foundedUser.Email != "" {
		return 0, errors.New("test")
	}

	err := repository.db.Create(&user).Error
	return user.Id, err
}

func (repository UserRepository) Update(user *model.User) error {
	return repository.db.Model(&user).Updates(&user).Error
}
