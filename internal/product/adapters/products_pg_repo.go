package adapters

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/feynmaz/fresheggs/internal/product/adapters/postgres"
	"github.com/feynmaz/fresheggs/internal/product/domain/product"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type productPgRepository struct {
	db *sqlx.DB
}

func NewProductPgRepository(pgDsn string) (*productPgRepository, error) {
	db, err := postgres.GetDb(pgDsn)
	if err != nil {
		return nil, err
	}
	return &productPgRepository{
		db: db,
	}, nil
}

func (p productPgRepository) CreateProduct(ctx context.Context, product *product.Product) error {
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
		return err
	}

	itemId := uuid.New().String()
	_, err = p.db.NamedExec(createItem, map[string]interface{}{
		"item_id":        itemId,
		"product_id":     product.ProductId,
		"price":          product.Price,
		"stock_quantity": product.StockQuantity,
	})
	if err != nil {
		return err
	}

	return nil
}

func (p productPgRepository) GetProduct(ctx context.Context, productId string) (*product.Product, error) {
	selectProductById := fmt.Sprintf(`
	select 
		p.product_id ,
		p."name" ,
		p.description ,
		i.price ,
		i.stock_quantity
	from products p 
	inner join items i 
	on p.product_id = i.product_id 
	where 1=1
	and p.product_id = '%s'`, productId)

	res := struct {
		Id            string  `db:"product_id"`
		Name          string  `db:"name"`
		Description   string  `db:"description"`
		Price         float32 `db:"price"`
		StockQuantity int     `db:"stock_quantity"`
	}{}
	err := p.db.Get(&res, selectProductById)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, product.NotFoundError{
				ProductId: productId,
			}
		}
		return nil, err
	}
	foundProduct := &product.Product{
		ProductId:     res.Id,
		Name:          res.Name,
		Description:   res.Description,
		Price:         res.Price,
		StockQuantity: res.StockQuantity,
	}
	return foundProduct, nil
}

func (p productPgRepository) UpdateProduct(ctx context.Context, productId string, productIn *product.Product) (*product.Product, error) {
	// TODO: implement
	// updateProduct := `
	// update products
	// set "name"=:name, description=:description
	// where product_id = :product_id
	// `
	// res, err := p.db.NamedExec(updateProduct, map[string]interface{}{
	// 	"product_id":  productId,
	// 	"name":        productIn.Name,
	// 	"description": productIn.Description,
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// rowsAffected, err := res.RowsAffected()
	// if err != nil {
	// 	return nil, err
	// }
	// if rowsAffected == 0 {
	// 	err = product.NotFoundError{
	// 		ProductId: productId,
	// 	}
	// }

	// updateItem := `
	// update items
	// set price=:price, stock_quantity=:stock_quantity
	// where product_id = :product_id
	// `
	// _, err = p.db.NamedExec(updateItem, map[string]interface{}{
	// 	"product_id":     productId,
	// 	"price":          productIn.Price,
	// 	"stock_quantity": productIn.StockQuantity,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (p productPgRepository) DeleteProduct(ctx context.Context, productId string) error {
	deleteProduct := `
	delete from products 
	where product_id = :product_id 
	`
	res, err := p.db.NamedExec(deleteProduct, map[string]interface{}{
		"product_id": productId,
	})
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return product.NotFoundError{ProductId: productId}
	}

	return nil
}

func (p productPgRepository) GetAllProducts(ctx context.Context, limit, offset int) ([]*product.Product, error) {
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
	productsRes := []struct {
		Id    string  `db:"product_id"`
		Name  string  `db:"name"`
		Price float32 `db:"price"`
	}{}

	err := p.db.Select(&productsRes, selectAllProducts)
	if err != nil {
		return nil, err
	}
	products := make([]*product.Product, 0, len(productsRes))
	for _, res := range productsRes {
		products = append(products, &product.Product{
			ProductId: res.Id,
			Name:      res.Name,
			Price:     res.Price,
		})
	}

	return products, nil
}
