package product

import (
	"context"

	"github.com/feynmaz/fresheggs/internal/domain/entity"
	"github.com/google/uuid"
)

type Service interface {
	GetProduct(ctx context.Context, productId string) (*entity.Product, error)
	GetProducts(ctx context.Context, limit, offset int) ([]*entity.Product, error)
	CreateProduct(ctx context.Context, product entity.Product) (string, error)
	UpdateProduct(ctx context.Context, productId string, product entity.Product) (*entity.Product, error)
	DeleteProduct(ctx context.Context, productId string) error
}

type ProductCreate interface {
	GetDescription() string
	GetName() string
	GetPrice() float32
	GetStockQuantity() int
}

type productUsecase struct {
	productService Service
}

func NewProductUsecase(productService Service) *productUsecase {
	return &productUsecase{
		productService: productService,
	}
}

func (u *productUsecase) CreateProduct(ctx context.Context, productId string, productCreate ProductCreate) (string, error) {
	if productId == "" {
		productId = uuid.New().String()
	}
	product := entity.Product{
		ProductId:     productId,
		Name:          productCreate.GetName(),
		Description:   productCreate.GetDescription(),
		Price:         productCreate.GetPrice(),
		StockQuantity: productCreate.GetStockQuantity(),
	}
	return u.productService.CreateProduct(ctx, product)
}

func (u *productUsecase) GetProducts(ctx context.Context, limit, offset int) ([]*entity.Product, error) {
	return u.productService.GetProducts(ctx, limit, offset)
}

func (u *productUsecase) GetProduct(ctx context.Context, id string) (*entity.Product, error) {
	return u.productService.GetProduct(ctx, id)
}

func (u *productUsecase) UpdateProduct(ctx context.Context, productId string, productPatch ProductCreate) (*entity.Product, error) {
	currentProduct, _ := u.productService.GetProduct(ctx, productId)
	updatedProduct := entity.Product{
		ProductId:   currentProduct.ProductId,
		Name:        currentProduct.Name,
		Description: currentProduct.Description,
		Price:       currentProduct.Price,
	}

	if productPatch.GetDescription() != "" {
		updatedProduct.Description = productPatch.GetDescription()
	}

	if productPatch.GetName() != "" {
		updatedProduct.Name = productPatch.GetName()
	}

	if productPatch.GetPrice() != 0 {
		updatedProduct.Price = productPatch.GetPrice()
	}

	return u.productService.UpdateProduct(ctx, productId, updatedProduct)
}

func (u *productUsecase) DeleteProduct(ctx context.Context, productId string) error {
	return u.productService.DeleteProduct(ctx, productId)
}
