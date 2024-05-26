package middleware

import (
	"net/http"

	"github.com/feynmaz/fresheggs/internal/metrics"
)

// Wrapper for http.ResponseWriter that preserves HTTP status code.
type statusRecorder struct {
	http.ResponseWriter
	code int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.code = code
	rec.ResponseWriter.WriteHeader(code)
}

func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := statusRecorder{w, 200}

		next.ServeHTTP(&rec, r)

		metrics.RequestsTotal.Inc()
		metrics.StatusCodes.WithLabelValues(http.StatusText(rec.code)).Inc()
	})
}
