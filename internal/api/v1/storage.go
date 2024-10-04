package v1

import "github.com/feynmaz/fresheggs/internal/types"

type Storage interface {
	GetProduct(productID string) (types.Product, error)
}
