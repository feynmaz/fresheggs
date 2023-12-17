package product

import (
	"encoding/json"
)

type Product struct {
	ProductId     string  `json:"product_id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float32 `json:"price"`
	StockQuantity int     `json:"stock_quantity"`
}

func (p Product) String() string {
	bytes, _ := json.Marshal(p)
	return string(bytes)
}

func (p Product) IsEmpty() bool {
	return p.ProductId == ""
}
