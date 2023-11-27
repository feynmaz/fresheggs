package product

import (
	"context"

	"github.com/feynmaz/fresheggs/internal/domain/entity"
)

type Service interface {
	GetProduct(ctx context.Context, productId string) (*entity.Product, error)
	GetProducts(ctx context.Context, limit, offset int) ([]*entity.Product, error)
	CreateProduct(ctx context.Context, product entity.Product) (string, error)
	UpdateProduct(ctx context.Context, productId string, product entity.Product) (*entity.Product, error)
	DeleteProduct(ctx context.Context, productId string) error
}

type productUsecase struct {
	productService Service
}

func NewProductUsecase(productService Service) *productUsecase {
	return &productUsecase{
		productService: productService,
	}
}

func (u *productUsecase) CreateProduct(ctx context.Context, name, description string, price float32) (string, error) {
	product := entity.Product{
		Name:        name,
		Description: description,
		Price:       price,
	}
	return u.productService.CreateProduct(ctx, product)
}

func (u *productUsecase) GetProducts(ctx context.Context, limit, offset int) ([]*entity.Product, error) {
	return u.productService.GetProducts(ctx, limit, offset)
}

func (u *productUsecase) GetProduct(ctx context.Context, id string) (*entity.Product, error) {
	return u.productService.GetProduct(ctx, id)
}
