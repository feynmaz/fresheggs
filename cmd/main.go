package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/feynmaz/fresheggs/config"
	"github.com/feynmaz/fresheggs/internal/adapters/db/memory"
	v1 "github.com/feynmaz/fresheggs/internal/controller/http/v1"
	"github.com/feynmaz/fresheggs/internal/domain/service"
	"github.com/feynmaz/fresheggs/internal/domain/usecase/product"
	"github.com/feynmaz/fresheggs/migrations"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg, err := config.GetDefault()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)

	if cfg.DoMigrations {
		if err := migrations.Run(cfg.PgDsn); err != nil {
			log.Fatal(err)
		}
	}

	serviceStorage := memory.NewProductStorage()
	productService := service.NewProductService(serviceStorage)
	productUsecase := product.NewProductUsecase(productService)

	ctx := context.Background()

	products, _ := productUsecase.GetProducts(ctx, 10, 0)
	fmt.Println(products)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Buy fresh eggs here!"))
	})

	productHandlerV1 := v1.NewProductHandler(productUsecase)
	productHandlerV1.Register(router)

	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router)
}
