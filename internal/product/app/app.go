package app

import (
	"context"

	"github.com/feynmaz/fresheggs/internal/product/domain/product"
	"github.com/google/uuid"
)

type ProductService struct {
	repo product.Repository
}

func NewProductService(repo product.Repository) ProductService {
	return ProductService{
		repo: repo,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, productCreate ProductCreate) (*product.Product, error) {
	product := &product.Product{
		ProductId:     uuid.New().String(),
		Name:          productCreate.GetName(),
		Description:   productCreate.GetDescription(),
		Price:         productCreate.GetPrice(),
		StockQuantity: productCreate.GetStockQuantity(),
	}

	err := s.repo.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) GetProduct(ctx context.Context, productId string) (*product.Product, error) {
	return s.repo.GetProduct(ctx, productId)
}

func (s *ProductService) UpdateProduct(ctx context.Context, productId string, product *product.Product) (*product.Product, error) {
	// TODO: implement
	// return s.repo.UpdateProduct(ctx, productId, product)
	return nil, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, productId string) error {
	return s.repo.DeleteProduct(ctx, productId)
}

func (s *ProductService) GetProducts(ctx context.Context, limit, offset int) ([]*product.Product, error) {
	return s.repo.GetAllProducts(ctx, limit, offset)
}
