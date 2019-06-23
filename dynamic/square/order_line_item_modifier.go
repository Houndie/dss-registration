package square

type OrderLineItemModifier struct {
	Uid             string `json:"uid"`
	CatalogObjectId string `json:"catalog_object_id"`
	Name            string `json:"name"`
	BasePriceMoney  *Money `json:"base_price_money"`
	TotalPriceMoney *Money `json:"total_price_money"`
}
