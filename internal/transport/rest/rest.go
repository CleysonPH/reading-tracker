package rest

import "net/http"

type BookHandler interface {
	GetBooks(w http.ResponseWriter, r *http.Request)
}
