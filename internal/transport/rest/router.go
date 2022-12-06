package rest

import "net/http"

func NewRouter(bookHandler BookHandler) http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/api/v1/books", bookHandler.GetBooks)
	return router
}
