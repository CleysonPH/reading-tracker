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

// GetReadingSessionsByBookID implements ReadingSessionHandler
func (h *readingSessionHandler) GetReadingSessionsByBookID(w http.ResponseWriter, r *http.Request) {
	bookId, err := strconv.ParseInt(chi.URLParam(r, "bookId"), 10, 64)
	if err != nil {
		sendBadRequest(w, "Invalid book id", err)
		return
	}

	if !h.bookRepository.Exists(bookId) {
		sendNotFound(w, fmt.Sprintf("Book with id %d not found", bookId), errors.New("book not found"))
		return
	}

	readingSessions, err := h.readingSessionRepository.AllByBookID(bookId)
	if err != nil {
		sendInternalServerError(w, "Failed to get reading sessions", err)
		return
	}

	readingSessionsResponse := make([]dto.ReadingSessionResponse, len(readingSessions))
	for i, readingSession := range readingSessions {
		readingSessionsResponse[i].FromReadingSession(readingSession)
	}

	sendOk(w, readingSessionsResponse)
}

// DeleteReadingSession implements ReadingSessionHandler
func (h *readingSessionHandler) DeleteReadingSession(w http.ResponseWriter, r *http.Request) {
	readingSessionId, err := strconv.ParseInt(chi.URLParam(r, "readingSessionId"), 10, 64)
	if err != nil {
		sendBadRequest(w, "Invalid reading session id", err)
		return
	}

	readingSession, err := h.readingSessionRepository.Get(readingSessionId)
	if err != nil {
		if errors.Is(err, repository.ErrReadingSessionNotFound) {
			sendNotFound(w, fmt.Sprintf("Reading session with id %d not found", readingSessionId), err)
			return
		}
		sendInternalServerError(w, "Failed to get reading session", err)
		return
	}

	bookId, err := strconv.ParseInt(chi.URLParam(r, "bookId"), 10, 64)
	if err != nil {
		sendBadRequest(w, "Invalid book id", err)
		return
	}
	book, err := h.bookRepository.Get(bookId)
	if err != nil {
		if errors.Is(err, repository.ErrBookNotFound) {
			sendNotFound(w, fmt.Sprintf("Book with id %d not found", readingSession.BookID), err)
			return
		}
	}

	if bookId != readingSession.BookID {
		sendBadRequest(w, "Reading session does not belong to book", errors.New("reading session does not belong to book"))
		return
	}

	if err := h.readingSessionRepository.Delete(readingSessionId); err != nil {
		sendInternalServerError(w, "Failed to delete reading session", err)
		return
	}

	readPages := book.ReadPages - readingSession.ReadPages
	var readingStatus string
	if readPages == 0 {
		readingStatus = "to-read"
	} else if readPages == book.Pages {
		readingStatus = "read"
	} else {
		readingStatus = "reading"
	}
	err = h.bookRepository.UpdateReadPagesAndReadingStatus(bookId, readPages, readingStatus)
	if err != nil {
		sendInternalServerError(w, "Failed to update book read pages", err)
		return
	}

	sendNoContent(w)
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
