package square

type CatalogDiscount struct {
	Name         string `json:"name"`
	DiscountType string `json:"discount_type"`
	Percentage   string `json:"percentage"`
	AmountMoney  *Money `json:"amount_money"`
	PinRequired  bool   `json:"pin_required"`
	LabelColor   string `json:"label_color"`
}

func (*CatalogDiscount) isCatalogObjectType() {}
