package tools

import (
	"net/http"
)

type RequestIDKeyType string

const RequestIDKeyValue RequestIDKeyType = "X-Request-ID"

func GetRequestID(r *http.Request) string {
	reqID := r.Context().Value(RequestIDKeyValue)
	if s, ok := reqID.(string); ok {
		return s
	}

	return ""
}
