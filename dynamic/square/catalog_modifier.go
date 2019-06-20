package square

type CatalogModifier struct {
	Name       string `json:"name"`
	PriceMoney *Money `json:"price_money"`
}

func (*CatalogModifier) isCatalogObjectType() {}
