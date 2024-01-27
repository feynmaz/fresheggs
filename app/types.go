package app

type ProductCreate interface {
	GetName() string
	GetDescription() string
	GetPrice() float64
	GetQuantity() int
}
