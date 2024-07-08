package dto

type User struct {
	Name     string `json:"name"`
	TgID     int64  `json:"tg_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
