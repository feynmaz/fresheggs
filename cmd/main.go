package main

import (
	"fmt"
	"log"
	"net/http"

	v1 "github.com/feynmaz/fresheggs/internal/ports/http/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	routeHandlerV1 := v1.NewRouteHandler()
	routeHandlerV1.Register(router)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", 8080), router); err != nil {
		return err
	}
	return nil
}
