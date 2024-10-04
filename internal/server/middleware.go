package server

import (
	"context"
	"net/http"

	"github.com/feynmaz/fresheggs/internal/tools"
	"github.com/feynmaz/fresheggs/internal/types"
	"github.com/google/uuid"
)

func (s *Server) RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var requestID string
		if r.Header.Get(string(types.RequestIDKeyValue)) != "" {
			requestID = r.Header.Get(string(types.RequestIDKeyValue))
		} else {
			requestID = uuid.New().String()
		}
		ctx := context.WithValue(r.Context(), types.RequestIDKeyValue, requestID)
		r = r.Clone(ctx)

		next.ServeHTTP(w, r)
	})
}

func (s *Server) LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := tools.NewResponseWriter(w)
		next.ServeHTTP(rw, r)

		event := s.logger.Info()
		if rw.Status() != http.StatusOK {
			event = s.logger.Warn()
		}
		event.Msgf("%s %s | %d [RequestID: %s]", r.Method, r.RequestURI, rw.Status(), tools.GetRequestID(r))
	})
}
