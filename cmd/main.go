package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/feynmaz/fresheggs/internal/config"
	"github.com/feynmaz/fresheggs/internal/logger"
	"github.com/feynmaz/fresheggs/internal/metrics"
	"github.com/feynmaz/fresheggs/internal/server"
	"github.com/spf13/pflag"
)

func main() {
	log := logger.New()
	configPath := pflag.StringP("config", "c", "", "path to config file")
	pflag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}
	logger.SetGlobalLevel(cfg.LogLevel)
	metrics.InitMetrics(cfg.AppName + "_")

	s, err := server.New(cfg, log)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create server")
	}

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
