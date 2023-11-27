package product

import (
	"context"

	"github.com/feynmaz/fresheggs/internal/domain/entity"
)

type Service interface {
	GetProduct(ctx context.Context, id string) (*entity.Product, error)
	GetProducts(ctx context.Context, limit, offset int) ([]*entity.Product, error)
	CreateProduct(ctx context.Context, product entity.Product) (string, error)
	UpdateProduct(ctx context.Context, product entity.Product) (*entity.Product, error)
	DeleteProduct(ctx context.Context, productId string) error
}

type productUsecase struct {
	productService Service
}

func (u *productUsecase) GetProducts(ctx context.Context, limit, offset int) ([]*entity.Product, error) {
	return u.productService.GetProducts(ctx, limit, offset)
}

func (u *productUsecase) GetProduct(ctx context.Context, id string) (*entity.Product, error) {
	return u.productService.GetProduct(ctx, id)
}
