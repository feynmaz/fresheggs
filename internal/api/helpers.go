package api

import (
	"encoding/json"
	"net/http"

	"github.com/feynmaz/fresheggs/internal/api/middleware"
	"github.com/feynmaz/fresheggs/internal/errs"
)

type ErrorResponse struct {
	RequestID  string `json:"requestId"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// WriteError sends error message and status code.
func (api *API) WriteError(w http.ResponseWriter, r *http.Request, err error) {
	resp := ErrorResponse{
		RequestID: middleware.GetRequestID(r),
		Message:   err.Error(),
	}
	w.Header().Set("Content-Type", "application/json")

	api.log.Err(err).Send()
	switch err.(type) {
	case errs.ErrNoData, errs.ErrBadRequest:
		resp.StatusCode = http.StatusBadRequest
	default:
		resp.StatusCode = http.StatusInternalServerError
	}

	w.WriteHeader(resp.StatusCode)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		api.log.Err(err).Send()
	}
}

func (api *API) WriteJSON(w http.ResponseWriter, r *http.Request, response any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		api.WriteError(w, r, err)
	}
}
