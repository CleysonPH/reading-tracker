package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/CleysonPH/reading-tracker/internal/repository"
	"github.com/CleysonPH/reading-tracker/internal/transport/rest/dto"
	"github.com/CleysonPH/reading-tracker/internal/transport/rest/validator"
	"github.com/go-chi/chi/v5"
)

func NewReadingSessionHandler(
	bookRepository repository.BookRepository,
	readingSessionValidator validator.ReadingSessionValidator,
	readingSessionRepository repository.ReadingSessionRepository,
) ReadingSessionHandler {
	return &readingSessionHandler{
		bookRepository:           bookRepository,
		readingSessionValidator:  readingSessionValidator,
		readingSessionRepository: readingSessionRepository,
	}
}

type readingSessionHandler struct {
	bookRepository           repository.BookRepository
	readingSessionValidator  validator.ReadingSessionValidator
	readingSessionRepository repository.ReadingSessionRepository
}

// CreateReadingSession implements ReadingSessionHandler
func (h *readingSessionHandler) CreateReadingSession(w http.ResponseWriter, r *http.Request) {
	bookId, err := strconv.ParseInt(chi.URLParam(r, "bookId"), 10, 64)
	if err != nil {
		sendBadRequest(w, "Invalid book id", err)
		return
	}

	if !h.bookRepository.Exists(bookId) {
		sendNotFound(w, fmt.Sprintf("Book with id %d not found", bookId), errors.New("book not found"))
		return
	}

	readingSessionRequest := dto.ReadingSessionRequest{}
	if err := json.NewDecoder(r.Body).Decode(&readingSessionRequest); err != nil {
		sendBadRequest(w, "Invalid request payload", err)
		return
	}

	readingSessionRequest.BookID = bookId
	if err := h.readingSessionValidator.ValidateReadingSessionCreate(&readingSessionRequest); err != nil {
		sendBadRequest(w, "Validation failed", err)
		return
	}

	readingSession := readingSessionRequest.ToReadingSession()
	readingSession, err = h.readingSessionRepository.Create(readingSession)
	if err != nil {
		sendInternalServerError(w, "Failed to create reading session", err)
		return
	}

	sendCreated(w, readingSession)
}
