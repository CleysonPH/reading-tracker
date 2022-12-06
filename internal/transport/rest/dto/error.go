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
