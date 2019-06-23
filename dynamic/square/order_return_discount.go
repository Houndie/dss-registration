package square

type OrderReturnDiscount struct {
	Uid               string                     `json:"uid"`
	SourceDiscountUid string                     `json:"source_discount_uid"`
	CatalogObjectId   string                     `json:"catalog_object_id"`
	Name              string                     `json:"name"`
	Type              OrderLineItemDiscountType  `json:"type"`
	Percentage        string                     `json:"percentage"`
	AmountMoney       *Money                     `json:"amount_money"`
	AppliedMoney      *Money                     `json:"applied_money"`
	Scope             OrderLineItemDiscountScope `json:"scope"`
}
