package rest

import (
	"net/http"

	"github.com/CleysonPH/reading-tracker/internal/transport/rest/handler"
	"github.com/go-chi/chi/v5"
)

func NewRouter(bookHandler handler.BookHandler) http.Handler {
	router := chi.NewRouter()
	router.Get("/api/v1/books", bookHandler.GetBooks)
	return router
}
