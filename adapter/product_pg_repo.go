package adapter

// import (
// 	"context"
// 	"errors"
// 	"fmt"

// 	sq "github.com/Masterminds/squirrel"
// 	"github.com/feynmaz/fresheggs/adapter/postgres"
// 	"github.com/feynmaz/fresheggs/domain/product"
// 	"github.com/jackc/pgx/v5"
// )

// var (
// 	ErrGetDb = errors.New("failed to get db")
// )

// type productPgRepository struct {
// 	conn         *pgx.Conn
// 	pgSqlBuilder sq.StatementBuilderType
// }

// func NewProductPgRepository(pgDsn string) (*productPgRepository, error) {
// 	ctx := context.Background()
// 	conn, err := postgres.GetDb(ctx, pgDsn)
// 	if err != nil {
// 		return nil, fmt.Errorf("%w: %w", ErrGetDb, err)
// 	}
// 	return &productPgRepository{
// 		conn:         conn,
// 		pgSqlBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
// 	}, err
// }

// func (p productPgRepository) GetProduct(ctx context.Context, productId string) (product.Product, error) {
// 	queryBuilder := p.pgSqlBuilder.Select("*").From("products").Where(sq.Eq{"product_id": productId})
// 	query, args, _ := queryBuilder.ToSql()
// 	row := p.conn.QueryRow(ctx, query, args...)

// 	var result Product
// 	err := row.Scan(&result.ProductId, &result.Name, &result.Description, &result.Price, &result.Quantity)
// 	if err == pgx.ErrNoRows {
// 		return product.Product{}, fmt.Errorf("%w: %w", product.ErrProductNotFound, err)
// 	}
// 	if err != nil {
// 		return product.Product{}, err
// 	}
// 	prod := product.Product{
// 		ProductId:   result.ProductId,
// 		Name:        result.Name,
// 		Description: result.Description,
// 		Price:       result.Price,
// 		Quantity:    result.Quantity,
// 	}

// 	return prod, nil
// }

// func (p productPgRepository) CreateProduct(ctx context.Context, product product.Product) error {
// 	query, args, _ := p.pgSqlBuilder.Insert("products").
// 		Columns("product_id", "name", "description", "price", "stock_quantity").
// 		Values(product.ProductId, product.Name, product.Description, product.Price, product.Quantity).
// 		ToSql()

// 	_, err := p.conn.Exec(ctx, query, args...)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (p productPgRepository) DeleteProduct(ctx context.Context, productId string) error {
// 	queryBuilder := p.pgSqlBuilder.Delete("products").Where(sq.Eq{"product_id": productId})
// 	query, args, _ := queryBuilder.ToSql()

// 	_, err := p.conn.Exec(ctx, query, args...)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (p productPgRepository) UpdateProduct(ctx context.Context, prod product.Product) error {
// 	queryBuilder := p.pgSqlBuilder.Update("products").
// 		SetMap(sq.Eq{
// 			"name":           prod.Name,
// 			"description":    prod.Description,
// 			"price":          prod.Price,
// 			"stock_quantity": prod.Quantity,
// 		}).
// 		Where(sq.Eq{"product_id": prod.ProductId})
// 	query, args, _ := queryBuilder.ToSql()

// 	_, err := p.conn.Exec(ctx, query, args...)
// 	if err == pgx.ErrNoRows {
// 		return fmt.Errorf("%w: %w", product.ErrProductNotFound, err)
// 	}
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (p productPgRepository) GetProducts(ctx context.Context) ([]product.Product, error) {
// 	queryBuilder := p.pgSqlBuilder.Select("*").From("products")
// 	query, args, _ := queryBuilder.ToSql()

// 	rows, err := p.conn.Query(ctx, query, args...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var products []product.Product
// 	for rows.Next() {
// 		var result Product
// 		err := rows.Scan(&result.ProductId, &result.Name, &result.Description, &result.Price, &result.Quantity)
// 		if err != nil {
// 			return nil, nil
// 		}
// 		prod := product.Product{
// 			ProductId:   result.ProductId,
// 			Name:        result.Name,
// 			Description: result.Description,
// 			Price:       result.Price,
// 			Quantity:    result.Quantity,
// 		}
// 		products = append(products, prod)
// 	}

// 	return products, nil
// }
