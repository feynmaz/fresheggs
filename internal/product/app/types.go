package app

type ProductCreate interface {
	GetDescription() string
	GetName() string
	GetPrice() float32
	GetStockQuantity() int
}
