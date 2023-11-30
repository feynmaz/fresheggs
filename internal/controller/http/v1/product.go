package v1

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/feynmaz/fresheggs/internal/domain/entity"
	"github.com/go-chi/chi"
)

const (
	getProductURL  = "/product/{productId}"
	getProductsURL = "/products"
)

type ProductUsecase interface {
	GetProducts(ctx context.Context, limit, offset int) ([]*entity.Product, error)
	GetProduct(ctx context.Context, productId string) (*entity.Product, error)
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
	router.Get(getProductURL, h.GetProduct)
	router.Get(getProductsURL, h.GetAllProducts)
}

func (h *productHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := chi.URLParam(r, "productId")
	product, _ := h.productUsecase.GetProduct(r.Context(), productId)
	reponseBody, _ := json.Marshal(product)
	w.Write(reponseBody)
}

func (h *productHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, _ := h.productUsecase.GetProducts(r.Context(), 10, 0)
	reponseBody, _ := json.Marshal(products)
	w.Write(reponseBody)
}
