package app

import (
	"context"
	"errors"

	"github.com/feynmaz/fresheggs/domain/product"
	"github.com/google/uuid"
)

type ProductService struct {
	repo product.Repository
}

func NewProductService(repo product.Repository) ProductService {
	return ProductService{repo: repo}
}

func (s ProductService) GetProduct(ctx context.Context, productId string) (product.Product, error) {
	return s.repo.GetProduct(ctx, productId)
}

func (s ProductService) CreateProduct(ctx context.Context, productCreate ProductCreate) (product.Product, error) {
	if productCreate.GetName() == "" {
		return product.Product{}, errors.New("product must have name")
	}
	p := product.Product{
		Name:        productCreate.GetName(),
		Description: productCreate.GetDescription(),
		Price:       productCreate.GetPrice(),
		Quantity:    productCreate.GetQuantity(),
		ProductId:   uuid.New().String(),
	}
	s.repo.CreateProduct(ctx, p)
	return p, nil
}

func (s ProductService) DeleteProduct(ctx context.Context, productId string) error {
	return s.repo.DeleteProduct(ctx, productId)
}

func (s ProductService) UpdateProduct(ctx context.Context, productId string, productUpdate ProductCreate) (product.Product, error) {
	p, err := s.repo.GetProduct(ctx, productId)
	if err != nil {
		return p, err
	}
	if productUpdate.GetDescription() != "" && p.Description != productUpdate.GetDescription() {
		p.Description = productUpdate.GetDescription()
	}
	if productUpdate.GetName() != "" && p.Name != productUpdate.GetName() {
		p.Name = productUpdate.GetName()
	}
	if productUpdate.GetPrice() != 0 && p.Price != productUpdate.GetPrice() {
		p.Price = productUpdate.GetPrice()
	}
	if productUpdate.GetQuantity() != 0 && p.Quantity != productUpdate.GetQuantity() {
		p.Quantity = productUpdate.GetQuantity()
	}

	err = s.repo.UpdateProduct(ctx, p)
	return p, err
}

func (s ProductService) GetProducts(ctx context.Context) ([]product.Product, error) {
	return s.repo.GetProducts(ctx)
}
