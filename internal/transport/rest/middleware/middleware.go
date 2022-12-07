package middleware

import "net/http"

type LoggerMiddleware interface {
	Use(next http.Handler) http.Handler
}
