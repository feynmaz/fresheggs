package v1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/feynmaz/fresheggs/internal/domain/entity"
	"github.com/feynmaz/fresheggs/internal/domain/usecase/product"
	"github.com/go-chi/chi/v5"
)

type ProductUsecase interface {
	GetProducts(ctx context.Context, limit, offset int) ([]*entity.Product, error)
	GetProduct(ctx context.Context, productId string) (*entity.Product, error)
	CreateProduct(ctx context.Context, productId string, productCreate product.ProductCreate) (string, error)
	UpdateProduct(ctx context.Context, productId string, updatedProduct product.ProductCreate) (*entity.Product, error)
	DeleteProduct(ctx context.Context, productId string) error
}

type productHandler struct {
	productUsecase ProductUsecase
}

func NewProductHandler(productUsecase ProductUsecase) *productHandler {
	return &productHandler{
		productUsecase: productUsecase,
	}
}

func (h *productHandler) Register(router chi.Router) {
	HandlerFromMux(h, router)
}

// Delete product by ID
// (DELETE /product/{product_id})
func (h *productHandler) DeleteProductProductId(w http.ResponseWriter, r *http.Request, productId string) {
	_ = h.productUsecase.DeleteProduct(r.Context(), productId)
	w.WriteHeader(200)
}

// Get product by ID
// (GET /product/{product_id})
func (h *productHandler) GetProductProductId(w http.ResponseWriter, r *http.Request, productId string) {
	product, _ := h.productUsecase.GetProduct(r.Context(), productId)

	reponseProduct := Product{
		Description: &product.Description,
		Name:        &product.Name,
		Price:       &product.Price,
		ProductId:   &product.ProductId,
	}

	WriteJson(w, reponseProduct)
}

// Update product by ID
// (PATCH /product/{product_id})
func (h *productHandler) PatchProductProductId(w http.ResponseWriter, r *http.Request, productId string) {
	var productPatch ProductCreate
	body, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &productPatch)

	product, _ := h.productUsecase.UpdateProduct(r.Context(), productId, productPatch)
	reponseProduct := Product{
		Description: &product.Description,
		Name:        &product.Name,
		Price:       &product.Price,
		ProductId:   &product.ProductId,
	}

	WriteJson(w, reponseProduct)
}

// Get list of products
// (GET /products)
func (h *productHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, _ := h.productUsecase.GetProducts(r.Context(), 10, 0)
	productsSummary := make([]ProductSummary, 0, len(products))
	for _, product := range products {
		productsSummary = append(productsSummary, ProductSummary{
			Name:      &product.Name,
			Price:     &product.Price,
			ProductId: &product.ProductId,
		})
	}

	WriteJson(w, productsSummary)
}

// Create product
// (POST /product)
func (h *productHandler) PostProduct(w http.ResponseWriter, r *http.Request) {
	var productCreate ProductCreate
	body, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &productCreate)

	productId, _ := h.productUsecase.CreateProduct(r.Context(), "", productCreate)
	
	responseProduct := Product{
		Description: productCreate.Description,
		Name: productCreate.Name,
		Price: productCreate.Price,
		ProductId: &productId,
		StockQuantity: productCreate.StockQuantity,
	}

	WriteJson(w, responseProduct)
}
