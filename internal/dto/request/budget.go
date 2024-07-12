package request

type Budget struct {
	Title  string  `json:"title" binding:"required"`
	Type   string  `json:"type" binding:"required"`
	Amount float32 `json:"amount" binding:"required"`
	UserID string  `json:"user_id" binding:"required"`
}
