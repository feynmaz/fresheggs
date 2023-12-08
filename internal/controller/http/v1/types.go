package v1

func (p ProductPatch) GetDescription() string {
	if p.Description != nil {
		return *p.Description
	}
	return ""
}

func (p ProductPatch) GetName() string {
	if p.Name != nil {
		return *p.Name
	}
	return ""
}

func (p ProductPatch) GetPrice() float32 {
	if p.Price != nil {
		return *p.Price
	}
	return 0
}

func (p ProductPatch) GetStockQuantity() int {
	if p.StockQuantity != nil {
		return *p.StockQuantity
	}
	return 0
}
