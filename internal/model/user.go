package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint           `json:"id,omitempty" gorm:"primaryKey;autoIncrement" binding:"omitempty,numeric,gt=0"`
	TgId      int64          `json:"tg_id,omitempty" gorm:"tg_id:bigint;not null" binding:"omitempty,numeric,gt=0"`
	Name      string         `json:"name" gorm:"name:varchar(255);not null" binding:"omitempty,alphanum,min=2,max=255"`
	Email     string         `json:"email" gorm:"email:varchar(255);not null" binding:"required,email"`
	Password  string         `json:"password" gorm:"password:varchar(255);not null" binding:"required,ascii,gte=8"`
	CreatedAt time.Time      `json:"-" gorm:"created_at:timestamp;not null"`
	UpdatedAt time.Time      `json:"-" gorm:"updated_at:timestamp;not null"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"deleted_at:timestamp;default:null"`
	Budgets   []Budget       `json:"budgets,omitempty" gorm:"-"`
	Roles     []Role         `json:"-" gorm:"many2many:user_roles"`

	// fields only for request
	RefreshToken         []byte `json:"refresh_token,omitempty" gorm:"refresh_token:jsonb;default:null"`
	PasswordConfirmation string `json:"password_confirmation" gorm:"-" binding:"omitempty,ascii,gte=8,eqfield=Password"`
}

type UserResponse struct {
	Id    uint   `json:"id,omitempty"`
	Name  string `json:"name"`
	TgId  uint64 `json:"tg_id,omitempty"`
	Email string `json:"email"`
	Role  Role   `json:"role"`
}
