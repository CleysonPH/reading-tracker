package rest

import (
	"net/http"

	"github.com/CleysonPH/reading-tracker/internal/transport/rest/handler"
	"github.com/go-chi/chi/v5"
)

func NewRouter(bookHandler handler.BookHandler) http.Handler {
	router := chi.NewRouter()

	router.Get("/api/v1/books", bookHandler.GetBooks)
	router.Get("/api/v1/books/{bookId}", bookHandler.GetBook)
	router.Delete("/api/v1/books/{bookId}", bookHandler.DeleteBook)
	router.Post("/api/v1/books", bookHandler.CreateBook)

	return router
}
