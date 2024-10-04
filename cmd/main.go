package main

import (
	"github.com/feynmaz/fresheggs/internal/config"
	"github.com/feynmaz/fresheggs/internal/logger"
)

func main() {
	log := logger.New()

	cfg, err := config.GetDefault()
	if err != nil {
		log.Err(err).Msg("failed to get config")
	}

	log.SetLevel(cfg.LogLevel)
	log.Debug().Msgf("config: %#v", cfg)
}
