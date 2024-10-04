package v1

import (
	"net/http"

	"github.com/feynmaz/fresheggs/internal/logger"
)

type API struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *API {
	return &API{
		logger: logger,
	}
}

func (api *API) GetHandler() http.Handler {
	return Handler(api)
}
