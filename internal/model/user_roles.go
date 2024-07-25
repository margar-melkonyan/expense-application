package model

type UserRole struct {
	UserID uint `json:"-" gorm:"primary_key"`
	RoleID uint `json:"role_id,omitempty" gorm:"primary_key" binding:"required,numeric,gt=0"`
	User   User `gorm:"foreignkey:UserID"`
	Role   Role `gorm:"foreignkey:RoleID"`
}
