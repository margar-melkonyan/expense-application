package model

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	Id        uint           `json:"id"  gorm:"primaryKey;autoIncrement" binding:"omitempty,numeric,gt=0"`
	Name      string         `json:"name" gorm:"type:varchar(255);not null" binding:"required,alpha,min=6,max=2048"`
	Slug      string         `json:"-" gorm:"name:varchar(255);not null;unique"`
	Type      string         `json:"type" gorm:"type:varchar(255);not null" binding:"required,alpha,oneof=income expense,min=4,max=255"`
	CreatedAt time.Time      `json:"-" gorm:"created_at:timestamp;not null"`
	UpdatedAt time.Time      `json:"-" gorm:"updated_at:timestamp;not null"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"deleted_at:timestamp;default:null"`
	Budgets   []*Budget      `json:"-" gorm:"many2many:budget_categories"`
}
