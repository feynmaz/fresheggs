package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/feynmaz/fresheggs/config"
	"github.com/feynmaz/fresheggs/internal/product/adapters"
	"github.com/feynmaz/fresheggs/internal/product/app"
	v1 "github.com/feynmaz/fresheggs/internal/product/ports/http/v1"

	"github.com/feynmaz/fresheggs/migrations"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg, err := config.GetDefault()
	if err != nil {
		return err
	}
	fmt.Println(cfg)

	if cfg.DoMigrations {
		if err := migrations.Run(cfg.PgDsn); err != nil {
			return err
		}
	}

	// product
	postgresRepo, err := adapters.NewProductPgRepository(cfg.PgDsn)
	if err != nil {
		return err
	}

	productService := app.NewProductService(postgresRepo)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Buy fresh eggs here!"))
	})

	productHandlerV1 := v1.NewProductHandler(productService)
	productHandlerV1.Register(router)

	log.Println("server is starting")
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router)
	return nil
}
