package repository

import (
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

func (repository UserRepository) CurrentUser() (model.User, error) {
	return model.User{}, nil
}

func (repository UserRepository) CurrentTgUser(tgId int64) (model.User, error) {
	var user model.User
	err := repository.db.Model(model.User{}).Where("tg_id = ?", tgId).First(&user).Error

	return user, err
}

func (repository UserRepository) SignUpByTg(user model.User) error {
	repository.db.Create(&user)

	return nil
}

func (repository UserRepository) SignUp(user model.User) (model.User, error) {
	return user, nil
}

func (repository UserRepository) SignIn(user model.User) (model.User, error) {
	return user, nil
}
