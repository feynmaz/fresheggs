package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/feynmaz/fresheggs/internal/config"
	"github.com/feynmaz/fresheggs/internal/logger"
	"github.com/feynmaz/fresheggs/internal/server"
)

func main() {
	log := logger.New()

	cfg, err := config.GetDefault()
	if err != nil {
		log.Err(err).Msg("failed to get config")
	}

	log.SetLevel(cfg.LogLevel)
	log.Debug().Msgf("config: %#v", cfg)

	s, err := server.New(cfg, log)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create server")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	go func() {
		err = s.Run(ctx)
		if err != nil {
			log.Fatal().Err(err).Msg("failed running server")
		}
	}()

	<-ctx.Done()
	s.Shutdown()
}
