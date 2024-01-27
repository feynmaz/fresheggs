package product

import (
	"context"
	"errors"
)

var (
	ErrProductNotFound = errors.New("product not found by id")
)

type Repository interface {
	GetProduct(ctx context.Context, productId string) (Product, error)
	CreateProduct(ctx context.Context, product Product) error
	DeleteProduct(ctx context.Context, productId string) error
	UpdateProduct(ctx context.Context, product Product) error
	GetProducts(ctx context.Context) ([]Product, error)
}
