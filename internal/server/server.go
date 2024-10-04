package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	v1 "github.com/feynmaz/fresheggs/internal/api/v1"
	"github.com/feynmaz/fresheggs/internal/config"
	"github.com/feynmaz/fresheggs/internal/logger"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	cfg    *config.Config
	logger *logger.Logger

	v1 *v1.API
}

func New(cfg *config.Config, logger *logger.Logger, v1 *v1.API) *Server {
	s := &Server{
		cfg:    cfg,
		logger: logger,
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

	return srv.ListenAndServe()
}

func (s *Server) getRouter() *chi.Mux {
	router := chi.NewMux()

	router.Use(s.RequestID)
	router.Use(s.LoggerMiddleware)

	router.Mount("/api/v1", s.v1.GetHandler())

	return router
}

func (s *Server) Shutdown() {
	s.logger.Info().Msg("graceful server shutdown")
}
