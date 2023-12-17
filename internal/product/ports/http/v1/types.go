package v1

func (p ProductCreate) GetDescription() string {
	if p.Description != nil {
		return *p.Description
	}
	return ""
}

func (p ProductCreate) GetName() string {
	if p.Name != nil {
		return *p.Name
	}
	return ""
}

func (p ProductCreate) GetPrice() float32 {
	if p.Price != nil {
		return *p.Price
	}
	return 0
}

func (p ProductCreate) GetStockQuantity() int {
	if p.StockQuantity != nil {
		return *p.StockQuantity
	}
	return 0
}
