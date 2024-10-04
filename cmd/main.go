package main

import (
	"context"
	"os"
	"os/signal"

	v1 "github.com/feynmaz/fresheggs/internal/api/v1"
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

	v1 := v1.New(log)
	s := server.New(cfg, log, v1)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	go func() {
		err = s.Run(ctx)
		if err != nil {
			log.Fatal().Err(err).Msg("server error")
		}
	}()

	<-ctx.Done()
	s.Shutdown()
}
