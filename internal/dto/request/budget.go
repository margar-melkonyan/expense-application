package request

type Budget struct {
	Title  string  `json:"title" binding:"required,min=6,max=4096"`
	Type   string  `json:"type" binding:"required,min=6,max=255"`
	Amount float32 `json:"amount" binding:"required,min=1,max=400000"`
	UserID string  `json:"user_id" binding:"required,integer"`
}
