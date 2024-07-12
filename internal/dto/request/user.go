package request

type User struct {
	Name                 string `form:"name" json:"name" binding:"required"`
	TgID                 int64  `from:"tg_id" json:"tg_id" binding:"required"`
	Email                string `form:"email" json:"email"`
	Password             string `form:"password" json:"password" binding:"required"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation" binding:"required"`
}
