package v1

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	httpTools "github.com/feynmaz/fresheggs/internal/common/http"
	"github.com/feynmaz/fresheggs/internal/product/app"
	"github.com/feynmaz/fresheggs/internal/product/domain/product"
	"github.com/go-chi/chi/v5"
)

type productHandler struct {
	app app.ProductService
}

func NewProductHandler(app app.ProductService) *productHandler {
	return &productHandler{
		app: app,
	}
}

func (h *productHandler) Register(router chi.Router) {
	HandlerFromMux(h, router)
}

// Create product
// (POST /product)
func (h *productHandler) PostProduct(w http.ResponseWriter, r *http.Request) {
	var productCreate ProductCreate
	body, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &productCreate)

	p, err := h.app.CreateProduct(r.Context(), productCreate)
	if err != nil {
		code := http.StatusInternalServerError
		httpTools.WriteErr(w, code, err)
		return
	}

	responseProduct := Product{
		Description:   &p.Description,
		Name:          &p.Name,
		Price:         &p.Price,
		ProductId:     &p.ProductId,
		StockQuantity: &p.StockQuantity,
	}

	httpTools.WriteJson(w, responseProduct)
}

// Get product by ID
// (GET /product/{product_id})
func (h *productHandler) GetProductProductId(w http.ResponseWriter, r *http.Request, productId string) {
	p, err := h.app.GetProduct(r.Context(), productId)
	if err != nil {
		code := http.StatusInternalServerError
		if errors.As(err, &product.NotFoundError{}) {
			code = http.StatusNotFound
		}
		httpTools.WriteErr(w, code, err)
		return
	}

	reponseProduct := Product{
		Description: &p.Description,
		Name:        &p.Name,
		Price:       &p.Price,
		ProductId:   &p.ProductId,
	}

	httpTools.WriteJson(w, reponseProduct)
}

// Update product by ID
// (PATCH /product/{product_id})
func (h *productHandler) PatchProductProductId(w http.ResponseWriter, r *http.Request, productId string) {
	// TODO:: implement
	// var productPatch ProductCreate
	// body, _ := io.ReadAll(r.Body)
	// _ = json.Unmarshal(body, &productPatch)

	// product, _ := h.app.UpdateProduct(r.Context(), productId, productPatch)
	// reponseProduct := Product{
	// 	Description: &product.Description,
	// 	Name:        &product.Name,
	// 	Price:       &product.Price,
	// 	ProductId:   &product.ProductId,
	// }

	// httpTools.WriteJson(w, reponseProduct)
}

// Delete product by ID
// (DELETE /product/{product_id})
func (h *productHandler) DeleteProductProductId(w http.ResponseWriter, r *http.Request, productId string) {
	err := h.app.DeleteProduct(r.Context(), productId)
	if err != nil {
		code := http.StatusInternalServerError
		if errors.As(err, &product.NotFoundError{}) {
			code = http.StatusNotFound
		}
		httpTools.WriteErr(w, code, err)
		return
	}
	w.WriteHeader(200)
}

// Get list of products
// (GET /products)
func (h *productHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.app.GetProducts(r.Context(), 10, 0)
	if err != nil {
		code := http.StatusInternalServerError
		httpTools.WriteErr(w, code, err)
	}
	productsSummary := make([]ProductSummary, 0, len(products))
	for _, product := range products {
		productsSummary = append(productsSummary, ProductSummary{
			Name:      &product.Name,
			Price:     &product.Price,
			ProductId: &product.ProductId,
		})
	}

	httpTools.WriteJson(w, productsSummary)
}
