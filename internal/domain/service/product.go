package service

import (
	"context"

	"github.com/feynmaz/fresheggs/internal/domain/entity"
)

type ProductStorage interface {
	GetOne(ctx context.Context, id string) (*entity.Product, error)
	GetAll(ctx context.Context, limit, offset int) ([]*entity.Product, error)
	Create(ctx context.Context, product entity.Product) (string, error)
	Update(ctx context.Context, product entity.Product) (*entity.Product, error)
	Delete(ctx context.Context, productId string) error
}

type productService struct {
	storage ProductStorage
}

func NewProductService(storage ProductStorage) *productService {
	return &productService{
		storage: storage,
	}
}

func (s *productService) GetProduct(ctx context.Context, id string) (*entity.Product, error) {
	return s.storage.GetOne(ctx, id)
}

func (s *productService) GetProducts(ctx context.Context, limit, offset int) ([]*entity.Product, error) {
	return s.storage.GetAll(ctx, limit, offset)
}

func (s *productService) CreateProduct(ctx context.Context, product entity.Product) (string, error) {
	return s.storage.Create(ctx, product)
}

func (s *productService) UpdateProduct(ctx context.Context, product entity.Product) (*entity.Product, error) {
	return s.storage.Update(ctx, product)
}

func (s *productService) DeleteProduct(ctx context.Context, productId string) error {
	return s.storage.Delete(ctx, productId)
}
