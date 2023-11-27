package main

import (
	"context"
	"fmt"

	"github.com/feynmaz/fresheggs/internal/adapters/db/memory"
	"github.com/feynmaz/fresheggs/internal/domain/service"
	"github.com/feynmaz/fresheggs/internal/domain/usecase/product"
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

	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Buy fresh eggs here!"))
	// })
	// http.ListenAndServe(":8080", r)
}
