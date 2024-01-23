package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/feynmaz/fresheggs/config"
	"github.com/feynmaz/fresheggs/internal/adapter"
	"github.com/feynmaz/fresheggs/internal/app"
	v1 "github.com/feynmaz/fresheggs/internal/ports/http/v1"
	"github.com/feynmaz/fresheggs/migrations"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	router.Use(middleware.Logger)

	productRepo, err := adapter.NewProductPgRepository(cfg.PgDsn)
	if err != nil {
		return err
	}
	productService := app.NewProductService(productRepo)
	routeHandlerV1 := v1.NewRouteHandler(productService)
	routeHandlerV1.Register(router)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router); err != nil {
		return err
	}
	return nil
}
