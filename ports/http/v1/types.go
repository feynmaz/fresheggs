package v1

// Implement ProductCreate interface
func (p ProductCreate) GetName() string {
	return p.Name
}

func (p ProductCreate) GetDescription() string {
	return p.Description
}

func (p ProductCreate) GetPrice() float64 {
	return p.Price
}

func (p ProductCreate) GetQuantity() int {
	return p.StockQuantity
}

func (p ProductPatch) GetName() string {
	if p.Name != nil {
		return *p.Name
	}
	return ""
}

func (p ProductPatch) GetDescription() string {
	if p.Description != nil {
		return *p.Description
	}
	return ""
}

func (p ProductPatch) GetPrice() float64 {
	if p.Price != nil {
		return *p.Price
	}
	return 0
}

func (p ProductPatch) GetQuantity() int {
	if p.StockQuantity != nil {
		return *p.StockQuantity
	}
	return 0
}
