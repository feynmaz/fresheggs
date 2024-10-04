package server

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type RequestIDKeyType string

const RequestIDKeyValue RequestIDKeyType = "X-Request-ID"

func (s *Server) RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var requestID string
		if r.Header.Get(string(RequestIDKeyValue)) != "" {
			requestID = r.Header.Get(string(RequestIDKeyValue))
		} else {
			requestID = uuid.New().String()
		}
		ctx := context.WithValue(r.Context(), RequestIDKeyValue, requestID)
		r = r.Clone(ctx)

		next.ServeHTTP(w, r)
	})
}

func GetRequestID(r *http.Request) string {
	reqID := r.Context().Value(RequestIDKeyValue)
	if s, ok := reqID.(string); ok {
		return s
	}

	return ""
}

func (s *Server) LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.logger.Info().Msgf("%s: %s %s", GetRequestID(r), r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
