package entity

import "encoding/json"

type Product struct {
	ProductId   string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

func (p Product) String() string {
	bytes, _ := json.Marshal(p)
	return string(bytes)
}
