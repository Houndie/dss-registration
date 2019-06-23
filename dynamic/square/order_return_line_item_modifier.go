package square

type OrderReturnLineItemModifier struct {
	Uid               string `json:"uid"`
	SourceModifierUid string `json:"source_modifier_uid"`
	CatalogObjectId   string `json:"catalog_object_id"`
	Name              string `json:"name"`
	BasePriceMoney    *Money `json:"base_price_money"`
	TotalPriceMoney   *Money `json:"total_price_money"`
}
