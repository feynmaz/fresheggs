package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/feynmaz/fresheggs/adapter"
	promtools "github.com/feynmaz/fresheggs/adapter/prometheus"
	"github.com/feynmaz/fresheggs/app"
	"github.com/feynmaz/fresheggs/config"
	"github.com/feynmaz/fresheggs/migrations"
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

	if cfg.DoMigrations {
		if err := migrations.Run(cfg.PgDsn); err != nil {
			return err
		}
	}

	router := chi.NewRouter()
	reg := prometheus.NewRegistry()
	reg.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	metrics := promtools.NewMetrics(reg)
	router.Use(metrics.GetMiddleware())
	router.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))

	productRepo := adapter.NewMemoryProductRepo()
	// if err != nil {
	// 	return err
	// }
	productService := app.NewProductService(productRepo)
	routeHandlerV1 := v1.NewRouteHandler(productService)
	routeHandlerV1.Register(router.With(middleware.Logger))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router); err != nil {
		return err
	}
	return nil
}
