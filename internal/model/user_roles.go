package model

type UserRole struct {
	UserID uint `json:"-" gorm:"primary_key"`
	RoleID uint `json:"role_id,omitempty" gorm:"primary_key"`
	User   User `gorm:"foreignkey:UserID"`
	Role   Role `gorm:"foreignkey:RoleID"`
}
