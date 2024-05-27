package api

import (
	"net/http"

	"github.com/feynmaz/fresheggs/internal/types"
)

// LastBlock godoc
// @Summary      Create product
// @Tags         post
// @Produce      json
// @Param        request   body    types.CreateProduct  true  "Create product"
// @Success      200  {object}  types.Product
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v1/product [post]
func (api *API) createProduct(w http.ResponseWriter, r *http.Request) {
	response := types.Product{
		ProductId: 1,
		Name:      "quail egg",
		Price:     12.34,
		Count:     100,
	}
	api.WriteJSON(w, r, response)
}

// LastBlock godoc
// @Summary      Get products
// @Tags         get
// @Produce      json
// @Success      200  {array}  types.Product
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v1/products [get]
func (api *API) getProducts(w http.ResponseWriter, r *http.Request) {
	response := []types.Product{{
		ProductId: 1,
		Name:      "quail egg",
		Price:     12.34,
		Count:     100,
	}, {
		ProductId: 2,
		Name:      "chicken egg",
		Price:     23.43,
		Count:     120,
	}}
	api.WriteJSON(w, r, response)
}
