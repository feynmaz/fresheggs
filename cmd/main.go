package main

import (
	"context"
	"os"
	"os/signal"

	v1 "github.com/feynmaz/fresheggs/internal/api/v1"
	"github.com/feynmaz/fresheggs/internal/config"
	"github.com/feynmaz/fresheggs/internal/logger"
	"github.com/feynmaz/fresheggs/internal/server"
	"github.com/feynmaz/fresheggs/internal/storage"
	"github.com/feynmaz/fresheggs/internal/telemetry"
	"go.opentelemetry.io/otel"
)

func main() {
	log := logger.New()

	cfg, err := config.GetDefault()
	if err != nil {
		log.Err(err).Msg("failed to get config")
	}

	log.SetLevel(cfg.LogLevel)
	log.Debug().Msgf("config: %#v", cfg)

	storage := storage.NewMemory()

	v1 := v1.New(log, storage)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	shutdown, err := telemetry.InitTracer(cfg.TelemetryEndpoint, cfg.AppName)
	if err != nil {
		log.Err(err).Msg("failed to init tracer")
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Err(err).Msg("failed to shutdown tracer")
		}
	}()

	tracer := otel.GetTracerProvider().Tracer("server")

	s := server.New(cfg, log, tracer, v1)

	go func() {
		err = s.Run(ctx)
		if err != nil {
			log.Fatal().Err(err).Msg("server error")
		}
	}()

	<-ctx.Done()
	s.Shutdown()
}
