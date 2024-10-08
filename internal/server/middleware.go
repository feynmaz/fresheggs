package server

import (
	"context"
	"net/http"

	"github.com/feynmaz/fresheggs/internal/tools"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/attribute"
)

var (
	skipLoggingPaths = []string{"favicon", "debug", "metrics"}
)

func (s *Server) RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var requestID string
		if r.Header.Get(string(tools.RequestIDKeyValue)) != "" {
			requestID = r.Header.Get(string(tools.RequestIDKeyValue))
		} else {
			requestID = uuid.New().String()
		}
		ctx := context.WithValue(r.Context(), tools.RequestIDKeyValue, requestID)
		r = r.Clone(ctx)

		next.ServeHTTP(w, r)
	})
}

func (s *Server) TelemetryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, span := s.tracer.Start(
			r.Context(),
			r.URL.Path,
		)
		defer span.End()

		rw := tools.NewResponseWriter(w)
		r = r.Clone(ctx)

		next.ServeHTTP(rw, r)

		requestID := tools.GetRequestID(r)

		span.SetAttributes(
			attribute.Int("http.status_code", rw.Status()),
			attribute.String("http.method", r.Method),
			attribute.String("http.url", r.URL.String()),
			attribute.String("http.request_id", requestID),
		)

		// Logging
		if !tools.ContainsAny(r.URL.Path, skipLoggingPaths) {
			event := s.logger.Info()
			if rw.Status() != http.StatusOK {
				event = s.logger.Warn()
			}
			event.Msgf("%s %s | %d [RequestID: %s]", r.Method, r.RequestURI, rw.Status(), requestID)
		}
	})
}
