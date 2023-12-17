package product

import (
	"context"
	"fmt"
)

type NotFoundError struct {
	ProductId string
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("product %s not found", e.ProductId)
}

type Repository interface {
	CreateProduct(ctx context.Context, product *Product) error
	GetProduct(ctx context.Context, productId string) (*Product, error)
	UpdateProduct(ctx context.Context, productId string, product *Product) (*Product, error)
	DeleteProduct(ctx context.Context, productId string) error
	GetAllProducts(ctx context.Context, limit, offset int) ([]*Product, error)
}
