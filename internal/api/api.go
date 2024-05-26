package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/feynmaz/fresheggs/internal/api/middleware"
	"github.com/feynmaz/fresheggs/internal/config"
	"github.com/feynmaz/fresheggs/internal/logger"
	_ "github.com/feynmaz/fresheggs/openapi" // OpenAPI docs.
)

type API struct {
	cfg config.Config
	log logger.Logger

	router *chi.Mux
}

func New(
	cfg config.Config,
	log logger.Logger,
) *API {
	api := API{
		cfg:    cfg,
		log:    log,
		router: chi.NewRouter(),
	}

	api.router.Use(middleware.RequestID)
	api.router.Use(middleware.Metrics)

	api.registerEndpoints()

	return &api
}

func (api *API) Serve(ctx context.Context) error {
	api.log.Info().Msgf("HTTP server listen on port: %v", api.cfg.Server.ListenPort)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%v", api.cfg.Server.ListenPort),
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		Handler:      api.router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return srv.ListenAndServe()
}

// @title Fresheggs API
// @version 1.0
// @description Eggs shop on web3 technology
// @contact.name Nikolai Mazein
// @contact.email feynmaz@gmail.com
// @BasePath /
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

	// Healthcheck.
	api.router.HandleFunc("/healthcheck", api.HealthCheck)

	// Swagger.
	api.router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(
			fmt.Sprintf("http://localhost:%v/swagger/doc.json", api.cfg.Server.ListenPort),
		),
	))

	// Application handlers.
}
