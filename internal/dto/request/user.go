package request

type User struct {
	Name                 string `json:"name" binding:"required,min=2,max=255"`
	TgID                 int64  `json:"tg_id" binding:""`
	Email                string `json:"email"`
	Password             string `json:"password" binding:"required,min=8,max=255"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required,min=8,max=255"`
}
