package middleware

import (
	"net/http"
	"time"

	"github.com/CleysonPH/reading-tracker/internal/service"
)

func NewLoggerMiddleware(loggerService service.LoggerService) LoggerMiddleware {
	return &loggerMiddleware{
		loggerService: loggerService,
	}
}

type loggerMiddleware struct {
	loggerService service.LoggerService
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	r.body = b
	return r.ResponseWriter.Write(b)
}

// Use implements LoggerMiddleware
func (m *loggerMiddleware) Use(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &responseRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(rec, r)

		elapsed := time.Since(start)

		logger := m.loggerService.Info
		if rec.statusCode >= http.StatusInternalServerError {
			logger = m.loggerService.Error
		}

		logger(
			"method=%s path=%s status=%d elapsed=%s",
			r.Method,
			r.URL.Path,
			rec.statusCode,
			elapsed,
		)
	})
}
