package dto

type Budget struct {
	ID     int     `json:"-"`
	Type   string  `json:"type"`
	Amount float32 `json:"amount"`
}
