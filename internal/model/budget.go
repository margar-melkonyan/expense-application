package model

import (
	"gorm.io/gorm"
	"time"
)

type Budget struct {
	Id         uint           `json:"id,omitempty" gorm:"primaryKey;autoIncrement" binding:"omitempty,numeric,min=1"`
	Title      string         `json:"title" gorm:"size:4096;not null" binding:"required,alphanum,min=6,max=4096"`
	Type       string         `json:"type" gorm:"type:varchar(255);not null" binding:"required,alpha,oneof=income expense"`
	Amount     float64        `json:"amount" gorm:"amount:bigint;not null" binding:"required,numeric,min=1,max=1000000"`
	CreatedAt  time.Time      `json:"-" gorm:"created_at:timestamp;not null"`
	UpdatedAt  time.Time      `json:"-" gorm:"updated_at:timestamp;not null"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"deleted_at:timestamp;default:null"`
	UserID     uint           `json:"user_id" gorm:"foreignkey:User REFERENCES users(id)" binding:"required,min=1"`
	User       User           `json:"-" gorm:"association_foreignkey:ID"`
	Categories []*Category    `json:"-" gorm:"many2many:budget_categories"`

	// fields only for request
	CategorySlug string `json:"category_slug,omitempty" gorm:"-"`
}
