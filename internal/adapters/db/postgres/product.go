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
	selectProductById := fmt.Sprintf(`
	select 
		p.product_id ,
		p."name" ,
		p.description ,
		i.price 
	from products p 
	inner join items i 
	on p.product_id = i.product_id 
	where 1=1
	and p.product_id = '%s'`, id)

	product := entity.Product{}
	err := db.Get(&product, selectProductById)
	if err != nil {
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
