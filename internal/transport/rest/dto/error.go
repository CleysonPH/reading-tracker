package dto

import "net/http"

type ErrorResponse struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Cause   string `json:"cause"`
}

func NewErrorResponse(status int, message string, err error) *ErrorResponse {
	return &ErrorResponse{
		Status:  status,
		Error:   http.StatusText(status),
		Message: message,
		Cause:   err.Error(),
	}
}

type ValidationErrorResponse struct {
	*ErrorResponse
	Errors map[string][]string `json:"errors"`
}

func NewValidationErrorResponse(message string, err error, errors map[string][]string) *ValidationErrorResponse {
	return &ValidationErrorResponse{
		ErrorResponse: NewErrorResponse(http.StatusBadRequest, message, err),
		Errors:        errors,
	}
}
