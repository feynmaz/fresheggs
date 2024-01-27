package adapter

import (
	"context"
	"sync"

	"github.com/feynmaz/fresheggs/domain/product"
)

type MemoryProductRepo struct {
	mu       sync.Mutex
	products map[string]product.Product
}

func NewMemoryProductRepo() *MemoryProductRepo {
	return &MemoryProductRepo{
		products: make(map[string]product.Product),
	}
}

func (r *MemoryProductRepo) GetProduct(ctx context.Context, productId string) (product.Product, error) {
	r.mu.Lock()
	p, ok := r.products[productId]
	r.mu.Unlock()
	if !ok {
		return product.Product{}, product.ErrProductNotFound
	}
	return p, nil
}

func (r *MemoryProductRepo) CreateProduct(ctx context.Context, product product.Product) error {
	r.mu.Lock()
	r.products[product.ProductId] = product
	r.mu.Unlock()
	return nil
}

func (r *MemoryProductRepo) DeleteProduct(ctx context.Context, productId string) error {
	r.mu.Lock()
	delete(r.products, productId)
	r.mu.Unlock()
	return nil
}

func (r *MemoryProductRepo) UpdateProduct(ctx context.Context, product product.Product) error {
	r.mu.Lock()
	r.products[product.ProductId] = product
	r.mu.Unlock()
	return nil
}

func (r *MemoryProductRepo) GetProducts(ctx context.Context) ([]product.Product, error) {
	products := make([]product.Product, 0, len(r.products))
	for _, p := range r.products {
		products = append(products, p)
	}
	return products, nil
}
