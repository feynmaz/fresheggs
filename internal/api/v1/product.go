package v1

import (
	"fmt"
	"net/http"

	"github.com/feynmaz/fresheggs/internal/tools"
	"github.com/feynmaz/fresheggs/internal/types"
)

// Get product details by productID
// (GET /product/{productID})
func (api *API) GetProductById(w http.ResponseWriter, r *http.Request, productID string) {
	product, err := api.storage.GetProduct(productID)
	if err != nil {
		tools.WriteError(w, r, err)
		return
	}
	if product.ProductID == "" {
		err = types.NewErrNotFound(fmt.Sprintf("no product with id %s", productID))
		tools.WriteError(w, r, err)
		return
	}

	response := &ProductResponse{
		Description: &product.Description,
		Name:        product.Name,
		ProductID:   product.ProductID,
		Price:       product.Price,
	}

	tools.WriteJSON(w, r, response)
}
