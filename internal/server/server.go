package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"

	v1 "github.com/feynmaz/fresheggs/internal/api/v1"
	"github.com/feynmaz/fresheggs/internal/config"
	"github.com/feynmaz/fresheggs/internal/logger"
	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/otel/trace"
)

type Server struct {
	cfg    *config.Config
	logger *logger.Logger

	tracer trace.Tracer

	v1 *v1.API
}

func New(cfg *config.Config, logger *logger.Logger, tracer trace.Tracer, v1 *v1.API) *Server {
	s := &Server{
		cfg:    cfg,
		logger: logger,
		tracer: tracer,
		v1:     v1,
	}
	return s
}

func (s *Server) Run(ctx context.Context) error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%v", s.cfg.ServerPort),
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		Handler:      s.getRouter(),
		ReadTimeout:  s.cfg.ServerReadTimeout,
		WriteTimeout: s.cfg.ServerWriteTimeout,
	}

	s.logger.Info().Msgf("server started on port %d", s.cfg.ServerPort)
	return srv.ListenAndServe()
}

func (s *Server) getRouter() *chi.Mux {
	router := chi.NewMux()

	// Middleware
	router.Use(s.RequestID)
	router.Use(s.TelemetryMiddleware)

	// Profiler
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))

	// API
	router.Mount("/api/v1", s.v1.GetHandler())

	return router
}

func (s *Server) Shutdown() {
	s.logger.Info().Msg("graceful server shutdown")
}
