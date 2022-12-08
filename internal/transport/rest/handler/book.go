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

func NewBookHandler(
	bookRepository repository.BookRepository,
	bookValidator validator.BookValidator,
) BookHandler {
	return &bookHandler{
		bookRepository: bookRepository,
		bookValidator:  bookValidator,
	}
}

type bookHandler struct {
	bookRepository repository.BookRepository
	bookValidator  validator.BookValidator
}

// UpdateBook implements BookHandler
func (h *bookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookId, err := strconv.ParseInt(chi.URLParam(r, "bookId"), 10, 64)
	if err != nil {
		sendBadRequest(w, "Invalid book id", err)
		return
	}

	if !h.bookRepository.Exists(bookId) {
		message := fmt.Sprintf("Book with id %d not found", bookId)
		sendBadRequest(w, message, errors.New("book not found"))
		return
	}

	bookRequest := dto.BookRequest{}
	if err := json.NewDecoder(r.Body).Decode(&bookRequest); err != nil {
		sendBadRequest(w, "Invalid request payload", err)
		return
	}

	if err := h.bookValidator.ValidateBookUpdate(bookId, &bookRequest); err != nil {
		sendBadRequest(w, "Validation failed", err)
		return
	}

	book := bookRequest.ToBook()
	book.ID = bookId
	book, err = h.bookRepository.Update(book)
	if err != nil {
		sendInternalServerError(w, "Failed to update book", err)
		return
	}

	bookResponse := dto.BookResponse{}
	bookResponse.FromBook(book)

	sendJSON(w, http.StatusOK, bookResponse)
}

// CreateBook implements BookHandler
func (h *bookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	bookRequest := dto.BookRequest{}
	if err := json.NewDecoder(r.Body).Decode(&bookRequest); err != nil {
		sendBadRequest(w, "Invalid request payload", err)
		return
	}

	if err := h.bookValidator.ValidateBookCreate(&bookRequest); err != nil {
		sendBadRequest(w, "Validation failed", err)
		return
	}

	book := bookRequest.ToBook()
	bookId, err := h.bookRepository.Create(book)
	if err != nil {
		sendInternalServerError(w, "Failed to create book", err)
		return
	}

	bookResponse := dto.BookResponse{}
	bookResponse.FromBook(book)
	bookResponse.ID = bookId

	sendJSON(w, http.StatusCreated, bookResponse)
}

// DeleteBook implements BookHandler
func (h *bookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId, err := strconv.ParseInt(chi.URLParam(r, "bookId"), 10, 64)
	if err != nil {
		sendBadRequest(w, "Invalid book id", err)
		return
	}

	err = h.bookRepository.Delete(bookId)
	if err != nil {
		sendInternalServerError(w, "Failed to delete book", err)
		return
	}

	sendNoContent(w)
}

// GetBook implements BookHandler
func (h *bookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	bookId, err := strconv.ParseInt(chi.URLParam(r, "bookId"), 10, 64)
	if err != nil {
		sendBadRequest(w, "Invalid book id", err)
		return
	}

	book, err := h.bookRepository.Get(bookId)
	if err != nil {
		if errors.Is(err, repository.ErrBookNotFound) {
			sendNotFound(w, "Book not found", err)
			return
		}
		sendInternalServerError(w, "Failed to get book", err)
		return
	}

	bookResponse := dto.BookResponse{}
	bookResponse.FromBook(book)

	sendJSON(w, http.StatusOK, bookResponse)
}

// GetBooks implements BookHandler
func (h *bookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	books, err := h.bookRepository.All(q)
	if err != nil {
		sendInternalServerError(w, "Failed to get books", err)
		return
	}

	bookSummaryResponses := make([]dto.BookSummaryResponse, len(books))
	for i, book := range books {
		bookSummaryResponses[i].FromBook(book)
	}

	sendJSON(w, http.StatusOK, bookSummaryResponses)
}
