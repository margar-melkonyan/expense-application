package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint           `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	TgId      int64          `json:"tg_id,omitempty" gorm:"tg_id:bigint;not null"`
	Name      string         `json:"name" gorm:"name:varchar(255);not null"`
	Email     string         `json:"email" gorm:"email:varchar(255);not null"`
	Password  string         `json:"password" gorm:"password:varchar(255);not null"`
	CreatedAt time.Time      `json:"-" gorm:"created_at:timestamp;not null"`
	UpdatedAt time.Time      `json:"-" gorm:"updated_at:timestamp;not null"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"deleted_at:timestamp;default:null"`
	Budgets   []Budget       `json:"budgets,omitempty" gorm:"-"`

	// fields only for request
	PasswordConfirmation string `json:"password_confirmation" gorm:"-"`
}
