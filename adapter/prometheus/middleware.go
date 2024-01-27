package prometheus

import (
	"net/http"
	"time"
)

func (m Metrics) GetMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			method := r.Method

			start := time.Now()
			next.ServeHTTP(w, r)
			duration := time.Since(start)

			m.HttpRequestDuration.WithLabelValues(path, method).Observe(duration.Seconds())
			m.HttpRequestsTotal.WithLabelValues(path, method).Inc()
		})
	}
}
