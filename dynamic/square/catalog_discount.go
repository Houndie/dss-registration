package square

type CatalogDiscount struct {
	Name         string `json:"name,omitempty"`
	DiscountType string `json:"discount_type,omitempty"`
	Percentage   string `json:"percentage,omitempty"`
	AmountMoney  *Money `json:"amount_money,omitempty"`
	PinRequired  bool   `json:"pin_required,omitempty"`
	LabelColor   string `json:"label_color,omitempty"`
}

func (*CatalogDiscount) isCatalogObjectType() {}
