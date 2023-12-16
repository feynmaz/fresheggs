package postgres

import (
	"context"
	"fmt"

	"github.com/feynmaz/fresheggs/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type productStorage struct {
	db *sqlx.DB
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
	err := p.db.Get(&product, selectProductById)
	if err != nil {
		return &product, err
	}

	return &product, nil
}

func (p productStorage) GetAll(ctx context.Context, limit, offset int) ([]*entity.Product, error) {
	selectAllProducts := `
	select 
		p.product_id ,
		p."name" ,
		i.price 
	from products p 
	inner join items i 
	on p.product_id = i.product_id 
	where 1=1
	`

	products := []*entity.Product{}
	err := p.db.Select(&products, selectAllProducts)
	if err != nil {
		return products, err
	}

	return products, nil
}

func (p productStorage) Create(ctx context.Context, product entity.Product) (string, error) {
	createProduct := `insert into products (product_id, name, description) values (:product_id, :name, :description)`
	createItem := `insert into items (item_id, product_id, price, stock_quantity) values (:item_id, :product_id, :price, :stock_quantity) 
	on conflict on constraint items_un_product_id 
	do update set stock_quantity = items.stock_quantity + 1`

	_, err := p.db.NamedExec(createProduct, map[string]interface{}{
		"product_id":  product.ProductId,
		"name":        product.Name,
		"description": product.Description,
	})
	if err != nil {
		return "", err
	}

	itemId := uuid.New().String()
	_, err = p.db.NamedExec(createItem, map[string]interface{}{
		"item_id":        itemId,
		"product_id":     product.ProductId,
		"price":          product.Price,
		"stock_quantity": 1,
	})
	if err != nil {
		return "", err
	}

	return product.ProductId, nil
}

func (p productStorage) Update(ctx context.Context, id string, product entity.Product) (*entity.Product, error) {
	updateProduct := `
	update products 
	set "name"=:name, description=:description
	where product_id = :product_id
	`
	_, err := p.db.NamedExec(updateProduct, map[string]interface{}{
		"product_id":  product.ProductId,
		"name":        product.Name,
		"description": product.Description,
	})
	if err != nil {
		return nil, err
	}

	updateItem := `
	update items
	set price=:price, stock_quantity=:stock_quantity
	where product_id = :product_id
	`
	_, err = p.db.NamedExec(updateItem, map[string]interface{}{
		"product_id":     product.ProductId,
		"price":          product.Price,
		"stock_quantity": 0,
	})
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p productStorage) Delete(ctx context.Context, productId string) error {
	return nil
}
