package memory

import (
	"context"
	"sync"

	"github.com/feynmaz/fresheggs/internal/domain/entity"
	"github.com/google/uuid"
)

type productsStorage struct {
	// TODO: use sync.Pool
	mu       sync.Mutex
	products map[string]*entity.Product
}

func NewProductStorage() *productsStorage {
	return &productsStorage{
		products: make(map[string]*entity.Product),
	}
}

func (p *productsStorage) GetOne(ctx context.Context, id string) (*entity.Product, error) {
	for productId, product := range p.products {
		if productId == id {
			return product, nil
		}
	}

	return nil, nil
}

func (p *productsStorage) GetAll(ctx context.Context, limit, offset int) ([]*entity.Product, error) {
	products := make([]*entity.Product, 0, len(p.products))
	for _, product := range p.products {
		products = append(products, product)
	}

	return products, nil
}

func (p *productsStorage) Create(ctx context.Context, product entity.Product) (string, error) {
	// TODO: pass product without ID
	productId := uuid.New().String()
	p.mu.Lock()
	p.products[productId] = &product
	p.mu.Unlock()
	return productId, nil
}

func (p *productsStorage) Update(ctx context.Context, productId string, product entity.Product) (*entity.Product, error) {
	p.mu.Lock()
	p.products[productId] = &product
	p.mu.Unlock()
	return &product, nil
}

func (p *productsStorage) Delete(ctx context.Context, productId string) error {
	p.mu.Lock()
	delete(p.products, productId)
	p.mu.Unlock()
	return nil
}
