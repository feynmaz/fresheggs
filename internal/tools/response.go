package tools

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/feynmaz/fresheggs/internal/types"
)

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
	case types.ErrNoData, types.ErrBadRequest:
		resp.StatusCode = http.StatusBadRequest

	case types.ErrUnauthorized:
		resp.StatusCode = http.StatusUnauthorized

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
