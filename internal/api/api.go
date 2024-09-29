package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"

	"github.com/feynmaz/fresheggs/internal/config"
	"github.com/feynmaz/fresheggs/internal/logger"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type API struct {
	cfg *config.Config
	log *logger.Logger

	router *chi.Mux
}

func New(cfg *config.Config, log *logger.Logger) (*API, error) {
	api := &API{
		cfg: cfg,
		log: log,

		router: chi.NewRouter(),
	}

	api.registerEndpoints()

	return api, nil
}

func (api *API) registerEndpoints() {
	// Profiler.
	api.router.HandleFunc("/debug/pprof/", pprof.Index)
	api.router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	api.router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	api.router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	api.router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	api.router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	api.router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	api.router.Handle("/debug/pprof/block", pprof.Handler("block"))

	// Promethrus metrics.
	api.router.Handle("/metrics", promhttp.Handler())

	// Application handlers.

}

func (api *API) Serve(ctx context.Context) error {
	api.log.Info().Msgf("HTTP server started on port: %v", api.cfg.ServerPort)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%v", api.cfg.ServerPort),
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		Handler:      api.router,
		ReadTimeout:  api.cfg.ServerReadTimeout,
		WriteTimeout: api.cfg.ServerWriteTimeout,
	}

	return srv.ListenAndServe()
}
