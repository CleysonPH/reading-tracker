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

	book, err := h.bookRepository.Get(bookId)
	if err != nil {
		if errors.Is(err, repository.ErrBookNotFound) {
			sendNotFound(w, fmt.Sprintf("Book with id %d not found", bookId), err)
			return
		}
		sendInternalServerError(w, "Failed to get book", err)
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

	readPages := book.ReadPages + readingSessionRequest.ReadPages
	var readingStatus string
	if readPages == book.Pages {
		readingStatus = "read"
	} else {
		readingStatus = "reading"
	}
	if err := h.bookRepository.UpdateReadPagesAndReadingStatus(bookId, readPages, readingStatus); err != nil {
		sendInternalServerError(w, "Failed to update book read pages", err)
		return
	}

	readingSessionResponse := dto.ReadingSessionResponse{}
	readingSessionResponse.FromReadingSession(readingSession)
	sendCreated(w, readingSessionResponse)
}
