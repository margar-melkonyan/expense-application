package helper

import (
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Message string        `json:"message"`
	Errors  []ErrorDetail `json:"errors"`
}

type ErrorDetail struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func FormatValidationError(err interface{}) ErrorResponse {
	var errorResponse ErrorResponse
	errorResponse.Message = "Check your entered data!"

	for _, ver := range err.(validator.ValidationErrors) {
		var errorDetail ErrorDetail

		errorDetail.Field = ver.StructField()
		errorDetail.Error = ver.ActualTag()
		errorResponse.Errors = append(errorResponse.Errors, errorDetail)
	}

	return errorResponse
}
