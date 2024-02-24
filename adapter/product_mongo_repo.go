package adapter

import (
	"context"
	"errors"
	"fmt"

	"github.com/feynmaz/fresheggs/adapter/mongo"
	"github.com/feynmaz/fresheggs/domain/product"
	"go.mongodb.org/mongo-driver/bson"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
)

type MongoProductRepo struct {
	products *mongoDriver.Collection
}

func NewMongoProductRepo(ctx context.Context, mongoURI string) (MongoProductRepo, error) {
	client, err := mongo.GetMongoClient(ctx, mongoURI)
	if err != nil {
		return MongoProductRepo{}, err
	}
	products := client.Database("fresheggs").Collection("products")

	return MongoProductRepo{
		products: products,
	}, nil
}

func (r MongoProductRepo) GetProduct(ctx context.Context, productId string) (product.Product, error) {
	var result mongo.Product
	err := r.products.FindOne(ctx, bson.D{{Key: "product_id", Value: productId}}).
		Decode(&result)
	if err != nil {
		if errors.Is(err, mongoDriver.ErrNoDocuments) {
			return product.Product{}, fmt.Errorf("%w: %w", product.ErrProductNotFound, err)
		}
	}

	return product.Product{
		ProductId:   result.ProductId,
		Name:        result.Name,
		Description: result.Description,
		Price:       result.Price,
		Quantity:    result.Quantity,
	}, nil
}

func (r MongoProductRepo) CreateProduct(ctx context.Context, p product.Product) error {
	insertValue := mongo.Product{
		ProductId:   p.ProductId,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Quantity:    p.Quantity,
	}
	_, err := r.products.InsertOne(ctx, insertValue)
	if err != nil {
		return err
	}
	return nil
}

func (r MongoProductRepo) DeleteProduct(ctx context.Context, productId string) error {
	_, err := r.products.DeleteOne(ctx, bson.D{{Key: "product_id", Value: productId}})
	if err != nil {
		// TODO: test if this error can occur
		if errors.Is(err, mongoDriver.ErrNoDocuments) {
			return fmt.Errorf("%w: %w", product.ErrProductNotFound, err)
		}
	}
	return nil
}

func (r MongoProductRepo) UpdateProduct(ctx context.Context, p product.Product) error {
	update := mongo.Product{
		ProductId:   p.ProductId,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Quantity:    p.Quantity,
	}
	_, err := r.products.UpdateOne(ctx, bson.D{{Key: "product_id", Value: p.ProductId}}, update)
	if err != nil {
		// TODO: test if this error can occur
		if errors.Is(err, mongoDriver.ErrNoDocuments) {
			return fmt.Errorf("%w: %w", product.ErrProductNotFound, err)
		}
	}

	return nil
}

func (r MongoProductRepo) GetProducts(ctx context.Context) ([]product.Product, error) {
	cursor, err := r.products.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var results []mongo.Product
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	products := make([]product.Product, 0, len(results))
	for _, result := range results {
		products = append(products, product.Product{
			ProductId:   result.ProductId,
			Name:        result.Name,
			Description: result.Description,
			Price:       result.Price,
			Quantity:    result.Quantity,
		})
	}

	return products, nil
}
