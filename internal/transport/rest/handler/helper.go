package handler

import (
	"encoding/json"
	"net/http"

	"github.com/CleysonPH/reading-tracker/internal/transport/rest/dto"
)

func sendJSON(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		sendInternalServerError(w, "Failed to encode response", err)
	}
}

func sendInternalServerError(w http.ResponseWriter, message string, err error) {
	sendJSON(w, http.StatusInternalServerError, dto.NewErrorResponse(http.StatusInternalServerError, message, err))
}

func sendBadRequest(w http.ResponseWriter, message string, err error) {
	sendJSON(w, http.StatusBadRequest, dto.NewErrorResponse(http.StatusBadRequest, message, err))
}

func sendNotFound(w http.ResponseWriter, message string, err error) {
	sendJSON(w, http.StatusNotFound, dto.NewErrorResponse(http.StatusNotFound, message, err))
}
