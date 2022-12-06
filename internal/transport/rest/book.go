package rest

import (
	"encoding/json"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bookSummaryResponses := make([]dto.BookSummaryResponse, len(books))
	for i, book := range books {
		bookSummaryResponses[i].FromBook(book)
	}

	body, err := json.Marshal(bookSummaryResponses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
