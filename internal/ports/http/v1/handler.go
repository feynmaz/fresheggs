package v1

import (
	"net/http"

	httptools "github.com/feynmaz/fresheggs/internal/common/http"
	"github.com/go-chi/chi/v5"

)

type handler struct{}

func NewRouteHandler() handler {
	return handler{}
}

func (h handler) Register(router chi.Router) {
	HandlerFromMux(h, router)
}

// (POST /product)
func (h handler) PostProduct(w http.ResponseWriter, r *http.Request) {
	product := Product{
		Name:          "Chicken egg",
		Price:         0.2,
		StockQuantity: 100,
		ProductId:     "37c4b517-67b1-49d5-ac89-f3f2acb0bb31",
	}
	httptools.SendJSONResponse(w, http.StatusOK, product)
}

// Delete product by ID
// (DELETE /product/{product_id})
func (h handler) DeleteProductProductId(w http.ResponseWriter, r *http.Request, productId string) {
	deletedId := "37c4b517-67b1-49d5-ac89-f3f2acb0bb31"
	httptools.SendJSONResponse(w, http.StatusOK, deletedId)
}

// Get product by ID
// (GET /product/{product_id})
func (h handler) GetProductProductId(w http.ResponseWriter, r *http.Request, productId string) {
	product := Product{
		Name:          "Chicken egg",
		Price:         0.2,
		StockQuantity: 100,
		ProductId:     "37c4b517-67b1-49d5-ac89-f3f2acb0bb31",
	}
	httptools.SendJSONResponse(w, http.StatusOK, product)
}

// Update product by ID
// (PATCH /product/{product_id})
func (h handler) PatchProductProductId(w http.ResponseWriter, r *http.Request, productId string) {
	product := Product{
		Name:          "Chicken egg",
		Price:         0.2,
		StockQuantity: 100,
		ProductId:     "37c4b517-67b1-49d5-ac89-f3f2acb0bb31",
	}
	httptools.SendJSONResponse(w, http.StatusOK, product)
}

// Get list of products
// (GET /products)
func (h handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	result := []ProductSummary{{
		Name:          "Chicken egg",
		Price:         0.2,
		StockQuantity: 100,
	}, {
		Name:          "Quail egg",
		Price:         0.12,
		StockQuantity: 400,
	}}
	httptools.SendJSONResponse(w, http.StatusOK, result)
}
