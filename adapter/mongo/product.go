package mongo

type Product struct {
	ProductId   string  `bson:"product_id"`
	Name        string  `bson:"name"`
	Description string  `bson:"description"`
	Price       float64 `bson:"price"`
	Quantity    int     `bson:"stock_quantity"`
}
