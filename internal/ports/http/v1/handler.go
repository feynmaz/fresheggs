package v1

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/feynmaz/fresheggs/internal/app"
	httptools "github.com/feynmaz/fresheggs/internal/common/http"
	"github.com/feynmaz/fresheggs/internal/domain/product"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	productService app.ProductService
}

func NewRouteHandler(productService app.ProductService) handler {
	return handler{
		productService: productService,
	}
}

func (h handler) Register(router chi.Router) {
	HandlerFromMux(h, router)
}

// (POST /product)
func (h handler) PostProduct(w http.ResponseWriter, r *http.Request) {
	var productCreate ProductCreate
	body, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &productCreate)

	p, err := h.productService.CreateProduct(r.Context(), productCreate)
	if err != nil {
		httptools.SendJSONResponse(w, http.StatusInternalServerError, err)
		return
	}
	result := Product{
		Description:   p.Description,
		Name:          p.Name,
		Price:         p.Price,
		ProductId:     "",
		StockQuantity: p.Quantity,
	}
	httptools.SendJSONResponse(w, http.StatusCreated, result)
}

// Delete product by ID
// (DELETE /product/{product_id})
func (h handler) DeleteProductProductId(w http.ResponseWriter, r *http.Request, productId string) {
	if err := h.productService.DeleteProduct(r.Context(), productId); err != nil {
		httptools.SendJSONResponse(w, http.StatusInternalServerError, err)
		return
	}
	httptools.SendJSONResponse(w, http.StatusOK, productId)
}

// Get product by ID
// (GET /product/{product_id})
func (h handler) GetProductProductId(w http.ResponseWriter, r *http.Request, productId string) {
	p, err := h.productService.GetProduct(r.Context(), productId)
	if err != nil {
		if errors.Is(err, product.ErrProductNotFound) {
			httptools.SendJSONResponse(w, http.StatusNotFound, err)
			return
		}
		httptools.SendJSONResponse(w, http.StatusInternalServerError, err)
		return
	}
	product := Product{
		Description:   p.Description,
		Name:          p.Name,
		Price:         p.Price,
		ProductId:     "",
		StockQuantity: p.Quantity,
	}
	httptools.SendJSONResponse(w, http.StatusOK, product)
}

// Update product by ID
// (PATCH /product/{product_id})
func (h handler) PatchProductProductId(w http.ResponseWriter, r *http.Request, productId string) {
	var productPatch ProductPatch
	body, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(body, &productPatch)

	p, err := h.productService.UpdateProduct(r.Context(), productId, productPatch)
	if err != nil {
		if errors.Is(err, product.ErrProductNotFound) {
			httptools.SendJSONResponse(w, http.StatusNotFound, err)
			return
		}
		httptools.SendJSONResponse(w, http.StatusInternalServerError, err)
		return
	}
	product := Product{
		Description:   p.Description,
		Name:          p.Name,
		Price:         p.Price,
		StockQuantity: p.Quantity,
		ProductId:     p.ProductId,
	}
	httptools.SendJSONResponse(w, http.StatusOK, product)
}

// Get list of products
// (GET /products)
func (h handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	ps, err := h.productService.GetProducts(r.Context())
	if err != nil {
		httptools.SendJSONResponse(w, http.StatusInternalServerError, err)
		return
	}

	products := make([]Product, 0, len(ps))
	for _, p := range ps {
		products = append(products, Product{
			Description:   p.Description,
			Name:          p.Name,
			Price:         p.Price,
			StockQuantity: p.Quantity,
			ProductId:     p.ProductId,
		})
	}

	httptools.SendJSONResponse(w, http.StatusOK, products)
}
