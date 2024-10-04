package tools

import (
	"net/http"

	"github.com/feynmaz/fresheggs/internal/types"
)

func GetRequestID(r *http.Request) string {
	reqID := r.Context().Value(types.RequestIDKeyValue)
	if s, ok := reqID.(string); ok {
		return s
	}

	return ""
}
