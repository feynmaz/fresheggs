package postgres

import (
	"context"

	"github.com/feynmaz/fresheggs/internal/domain/entity"
)

type productStorage struct {
	db interface{}
}

func NewProductStorage(pgDsn string) (*productStorage, error) {
	db, err := GetDb(pgDsn)
	if err != nil {
		return nil, err
	}
	return &productStorage{
		db: db,
	}, nil
}

func (p productStorage) GetOne(ctx context.Context, id string) (*entity.Product, error) {
	return nil, nil
}

func (p productStorage) GetAll(ctx context.Context, limit, offset int) ([]*entity.Product, error) {
	return nil, nil
}

func (p productStorage) Create(ctx context.Context, product entity.Product) (string, error) {
	return "", nil
}

func (p productStorage) Update(ctx context.Context, id string, product entity.Product) (*entity.Product, error) {
	return nil, nil
}

func (p productStorage) Delete(ctx context.Context, productId string) error {
	return nil
}
