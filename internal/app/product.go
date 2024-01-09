package app

import (
	"errors"

	"github.com/feynmaz/fresheggs/internal/domain/product"
)

type ProductService struct {
}

func NewProductService() ProductService {
	return ProductService{}
}

func (s ProductService) GetProduct(productId string) product.Product {
	product := product.Product{
		Name:        "Chicken egg",
		Price:       0.2,
		Quantity:    100,
		Description: "Egg of a chicken",
	}
	return product
}

func (s ProductService) CreateProduct(productCreate ProductCreate) (product.Product, error) {
	if productCreate.GetName() == "" {
		return product.Product{}, errors.New("product must have name")
	}
	product := product.Product{
		Name:        productCreate.GetName(),
		Description: productCreate.GetDescription(),
		Price:       productCreate.GetPrice(),
		Quantity:    productCreate.GetQuantity(),
	}
	return product, nil
}

func (s ProductService) DeleteProduct(productId string) error {
	return nil
}

func (s ProductService) UpdateProduct(productUpdate ProductCreate) (product.Product, error) {
	return product.Product{}, nil
}

func (s ProductService) GetProducts() ([]product.Product, error) {
	products := []product.Product{{
		Name:        "Chicken egg",
		Price:       0.2,
		Quantity:    100,
		Description: "Egg of a chicken",
	}, {
		Name:        "Quail egg",
		Price:       0.015,
		Quantity:    800,
		Description: "Egg of a quail",
	}}
	return products, nil
}
