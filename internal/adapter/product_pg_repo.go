package adapter

import (
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/feynmaz/fresheggs/internal/adapter/postgres"
	"github.com/feynmaz/fresheggs/internal/domain/product"
	"github.com/jackc/pgx/v5"
)

var (
	ErrGetDb = errors.New("failed to get db")
)

type productPgRepository struct {
	conn         *pgx.Conn
	pgSqlBuilder sq.StatementBuilderType
}

func NewProductPgRepository(pgDsn string) (*productPgRepository, error) {
	conn, err := postgres.GetDb(pgDsn)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrGetDb, err)
	}
	return &productPgRepository{
		conn:         conn,
		pgSqlBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}, err
}

func (p productPgRepository) GetProduct(ctx context.Context, productId string) (product.Product, error) {
	products := p.pgSqlBuilder.Select("*").From("products").Join("items USING (product_id)")
	query, args, _ := products.Where(sq.Eq{"product_id": productId}).ToSql()

	prod := product.Product{}
	err := p.conn.QueryRow(ctx, query, args...).Scan(prod)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return prod, product.ErrProductNotFound
		}
		return prod, err
	}
	return prod, nil
}

func (p productPgRepository) CreateProduct(ctx context.Context, product product.Product) error {
	return nil
}

func (p productPgRepository) DeleteProduct(ctx context.Context, productId string) error {
	return nil
}

func (p productPgRepository) UpdateProduct(ctx context.Context, product product.Product) error {
	return nil
}

func (p productPgRepository) GetProducts(ctx context.Context) ([]product.Product, error) {
	return nil, nil
}
