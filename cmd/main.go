package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/feynmaz/fresheggs/internal/adapters/db/memory"
	v1 "github.com/feynmaz/fresheggs/internal/controller/http/v1"
	"github.com/feynmaz/fresheggs/internal/domain/service"
	"github.com/feynmaz/fresheggs/internal/domain/usecase/product"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	serviceStorage := memory.NewProductStorage()
	productService := service.NewProductService(serviceStorage)
	productUsecase := product.NewProductUsecase(productService)

	ctx := context.Background()
	productUsecase.CreateProduct(ctx, "chicken egg", "egg of a chicken", 0.1)
	productUsecase.CreateProduct(ctx, "quail egg", "egg of a quail", 0.08)

	products, _ := productUsecase.GetProducts(ctx, 10, 0)
	fmt.Println(products)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Buy fresh eggs here!"))
	})

	productHandlerV1 := v1.NewProductHandler(productUsecase)
	productHandlerV1.Register(router)

	http.ListenAndServe(":8080", router)
}
