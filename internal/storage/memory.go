package storage

import (
	"sync"

	"github.com/feynmaz/fresheggs/internal/types"
)

type Memory struct {
	products map[string]types.Product
	mu       sync.Mutex
}

func NewMemory() *Memory {
	products := map[string]types.Product{
		"0b87c363-7bcd-4707-9f30-41c773c3d28f": {
			ProductID:   "0b87c363-7bcd-4707-9f30-41c773c3d28f",
			Name:        "chicken egg",
			Description: "egg of a chicken",
			Price:       0.1,
			Quantity:    100,
		},
		"eff8868d-db6d-4765-a21a-2bce3dce3972": {
			ProductID:   "eff8868d-db6d-4765-a21a-2bce3dce3972",
			Name:        "quail egg",
			Description: "egg of a quail",
			Price:       0.14,
			Quantity:    400,
		},
	}

	return &Memory{
		products: products,
	}
}

func (m *Memory) GetProduct(productID string) (types.Product, error) {
	m.mu.Lock()
	product, ok := m.products[productID]
	m.mu.Unlock()

	if ok {
		return product, nil
	}

	return types.Product{}, nil
}
