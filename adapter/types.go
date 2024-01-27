package adapter

type Product struct {
	ProductId   string  `db:"product_id"`
	Name        string  `db:"name"`
	Description string  `db:"description"`
	Price       float64 `db:"price"`
	Quantity    int     `db:"stock_quantity"`
}
