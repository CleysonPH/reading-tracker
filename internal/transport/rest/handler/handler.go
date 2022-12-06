package handler

import "net/http"

type BookHandler interface {
	GetBooks(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
}
