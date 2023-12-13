package entity

import "encoding/json"

type Product struct {
	ProductId   string  `json:"id" db:"product_id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Price       float32 `json:"price" db:"price"`
}

func (p Product) String() string {
	bytes, _ := json.Marshal(p)
	return string(bytes)
}
