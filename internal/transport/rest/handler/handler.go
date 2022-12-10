package handler

import "net/http"

type BookHandler interface {
	GetBooks(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
	CreateBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
}

type ReadingSessionHandler interface {
	CreateReadingSession(w http.ResponseWriter, r *http.Request)
	DeleteReadingSession(w http.ResponseWriter, r *http.Request)
}
