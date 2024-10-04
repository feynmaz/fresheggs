package v1

import (
	"net/http"

	"github.com/feynmaz/fresheggs/internal/logger"
)

type API struct {
	logger *logger.Logger

	storage Storage
}

func New(logger *logger.Logger, storage Storage) *API {
	return &API{
		logger:  logger,
		storage: storage,
	}
}

func (api *API) GetHandler() http.Handler {
	return Handler(api)
}
