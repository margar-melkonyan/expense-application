package handler

type ErrorResponse struct {
	Message string `example:"Something went wrong try again later!" json:"message"`
}

type StatusResponse struct {
	Status string `example:"Success created!" json:"status"`
}
