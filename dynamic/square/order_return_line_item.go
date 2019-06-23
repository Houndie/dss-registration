package square

type OrderReturnLineItem struct {
	Uid                      string                         `json:"uid"`
	SourceLineItemUid        string                         `json:"source_line_item_uid"`
	Name                     string                         `json:"name"`
	Quantity                 string                         `json:"quantity"`
	QuantityUnit             *OrderQuantityUnit             `json:"quantity_unit"`
	Note                     string                         `json:"note"`
	CatalogObjectId          string                         `json:"catalog_object_id"`
	VariationName            string                         `json:"variation_name"`
	ReturnModifiers          []*OrderReturnLineItemModifier `json:"return_modifiers"`
	ReturnTaxes              []*OrderReturnTax              `json:"return_taxes"`
	ReturnDiscounts          []*OrderReturnDiscount         `json:"return_discounts"`
	BasePriceMoney           *Money                         `json:"base_price_money"`
	VariationTotalPriceMoney *Money                         `json:"variation_total_price_money"`
	GrossReturnMoney         *Money                         `json:"gross_return_money"`
	TotalTaxMoney            *Money                         `json:"total_tax_money"`
	TotalDiscountMoney       *Money                         `json:"total_discount_money"`
	TotalMoney               *Money                         `json:"total_money"`
}
