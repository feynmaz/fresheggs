package server

import (
	"context"
	"fmt"

	"github.com/feynmaz/fresheggs/internal/api"
	"github.com/feynmaz/fresheggs/internal/config"
	"github.com/feynmaz/fresheggs/internal/eth"
	"github.com/feynmaz/fresheggs/internal/logger"
)

type Server struct {
	cfg *config.Config
	log logger.Logger
	api *api.API
}

func New(cfg *config.Config, log logger.Logger) (*Server, error) {
	s := Server{
		cfg: cfg,
		log: log,
	}

	ethClient, err := eth.NewClient(cfg.EthClient, log)
	if err != nil {
		return nil, fmt.Errorf("failed to create eth client: %w", err)
	}

	s.api = api.New(*s.cfg, s.log, ethClient)

	return &s, nil
}

func (s *Server) Run(ctx context.Context) error {
	return s.api.Serve(ctx)
}

func (s *Server) Shutdown() {
	s.log.Info().Msg("graceful server shutdown")
}
