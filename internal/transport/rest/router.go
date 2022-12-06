package rest

import (
	"net/http"

	"github.com/CleysonPH/reading-tracker/internal/transport/rest/handler"
)

func NewRouter(bookHandler handler.BookHandler) http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/api/v1/books", bookHandler.GetBooks)
	return router
}
