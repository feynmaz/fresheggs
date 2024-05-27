package types

type CreateProduct struct {
	Address string  `json:"address"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Count   int     `json:"count"`
}

type ProductResponse struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Count int     `json:"count"`
}

type Product struct {
	ProductId uint64  `json:"productId"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Count     int     `json:"count"`
}
