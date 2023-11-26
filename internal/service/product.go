package service

import (
	"context"

	"github.com/feynmaz/fresheggs/internal/model"
)

type ProductService interface {
	GetProduct(ctx context.Context, id string) (*model.Product, error)
	GetProducts(ctx context.Context, limit, offset int) ([]*model.Product, error)
	CreateProduct(ctx context.Context, product model.Product) (string, error)
	UpdateProduct(ctx context.Context, product model.Product) (*model.Product, error)
	DeleteAnswer(ctx context.Context, productId string) error
}

type answerService struct {
	dao repository.DAO
}
