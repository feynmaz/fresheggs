package postgres

import (
	"context"
	"fmt"

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
	//  If a function name includes Query, it is designed to ask a question of the database, and will return a set of rows
	rows, err := db.Query(`
select 
	p.product_id ,
	p."name" ,
	p.description ,
	i.price 
from products p 
inner join items i 
on p.product_id = i.product_id 
where 1=1
and p.product_id = $1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var product entity.Product
	count := 0
	for rows.Next() {
		if err = rows.Scan(&product); err != nil {
			return nil, err
		}
		count++
	}
	if count == 0 {
		return nil, fmt.Errorf("product with id=%s not found", id)
	}
	if err = rows.Err(); err != nil {
		return &product, err
	}
	return &product, nil
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
