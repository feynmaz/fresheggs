package v1

import (
	"net/http"
)

func (api *API) LoggerMiddleware(next http.Handler) http.Handler {
	// TODO: get request ID
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.logger.Info().Msgf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
