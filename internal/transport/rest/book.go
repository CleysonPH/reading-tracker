package rest

import (
	"net/http"

	"github.com/CleysonPH/reading-tracker/internal/repository"
	"github.com/CleysonPH/reading-tracker/internal/transport/rest/dto"
)

func NewBookHandler(bookRepository repository.BookRepository) BookHandler {
	return &bookHandler{bookRepository: bookRepository}
}

type bookHandler struct {
	bookRepository repository.BookRepository
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
