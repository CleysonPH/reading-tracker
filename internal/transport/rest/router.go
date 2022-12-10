package rest

import (
	"net/http"

	"github.com/CleysonPH/reading-tracker/internal/transport/rest/handler"
	"github.com/go-chi/chi/v5"
)

func NewRouter(
	bookHandler handler.BookHandler,
	readingSessionHandler handler.ReadingSessionHandler,
) http.Handler {
	router := chi.NewRouter()

	// Books
	router.Get("/api/v1/books", bookHandler.GetBooks)
	router.Post("/api/v1/books", bookHandler.CreateBook)
	router.Get("/api/v1/books/{bookId}", bookHandler.GetBook)
	router.Put("/api/v1/books/{bookId}", bookHandler.UpdateBook)
	router.Delete("/api/v1/books/{bookId}", bookHandler.DeleteBook)

	// Reading sessions
	router.Post("/api/v1/books/{bookId}/reading-sessions", readingSessionHandler.CreateReadingSession)

	return router
}
