package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/feynmaz/fresheggs/adapter"
	promtools "github.com/feynmaz/fresheggs/adapter/prometheus"
	"github.com/feynmaz/fresheggs/app"
	"github.com/feynmaz/fresheggs/config"
	v1 "github.com/feynmaz/fresheggs/ports/http/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg := config.GetDefault()

	router := chi.NewRouter()
	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	metrics := promtools.NewMetrics(reg)
	router.Use(metrics.GetMiddleware())
	router.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))

	ctx := context.Background()
	productRepo, err := adapter.NewMongoProductRepo(ctx, cfg.MongoURI)
	if err != nil {
		return err
	}
	productService := app.NewProductService(productRepo)
	routeHandlerV1 := v1.NewRouteHandler(productService)
	routeHandlerV1.Register(router.With(middleware.Logger))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router); err != nil {
		return err
	}
	return nil
}
