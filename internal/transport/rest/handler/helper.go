package handler

import (
	"encoding/json"
	"net/http"

	"github.com/CleysonPH/reading-tracker/internal/transport/rest/dto"
	"github.com/CleysonPH/reading-tracker/internal/transport/rest/validator"
)

func sendJSON(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		sendInternalServerError(w, "Failed to encode response", err)
	}
}

func sendNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func sendInternalServerError(w http.ResponseWriter, message string, err error) {
	sendJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(http.StatusInternalServerError, message, err))
}

func sendBadRequest(w http.ResponseWriter, message string, err error) {
	ve, ok := err.(*validator.ValidationError)
	if ok {
		sendJSON(w, http.StatusBadRequest, dto.NewValidationErrorResponse(message, err, ve.Errors))
		return
	}
	sendJSON(w, http.StatusBadRequest, dto.NewErrorResponse(http.StatusBadRequest, message, err))
}

func sendNotFound(w http.ResponseWriter, message string, err error) {
	sendJSON(w, http.StatusNotFound, dto.NewErrorResponse(http.StatusNotFound, message, err))
}
