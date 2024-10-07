package tools

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/feynmaz/fresheggs/internal/types"
)

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) Status() int {
	return rw.statusCode
}

//

func WriteJSON(w http.ResponseWriter, r *http.Request, response any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		WriteError(w, r, err)
	}
}

func WriteError(w http.ResponseWriter, r *http.Request, err error) {
	resp := types.ErrorResponse{
		RequestID: GetRequestID(r),
		Message:   err.Error(),
	}
	w.Header().Set("Content-Type", "application/json")

	switch err.(type) {
	case types.ErrNotFound:
		resp.StatusCode = http.StatusNotFound

	default:
		resp.StatusCode = http.StatusInternalServerError
	}

	w.WriteHeader(resp.StatusCode)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		// TODO: what logger use here?
		// log.Err(err).Send()
		log.Fatal(err)
	}
}
