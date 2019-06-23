package square

type OrderLineItem struct {
	Uid                      string                   `json:"uid"`
	Name                     string                   `json:"name"`
	Quantity                 string                   `json:"quantity"`
	QuantityUnit             *OrderQuantityUnit       `json:"quantity_unit"`
	Note                     string                   `json:"note"`
	CatalogObjectId          string                   `json:"catalog_object_id"`
	VariationName            string                   `json:"variation_name"`
	Modifiers                []*OrderLineItemModifier `json:"modifiers"`
	Taxes                    []*OrderLineItemTax      `json:"taxes"`
	Discounts                []*OrderLineItemDiscount `json:"discounts"`
	BasePriceMoney           *Money                   `json:"base_price_money"`
	VariationTotalPriceMoney *Money                   `json:"variation_total_price_money"`
	GrossSalesMoney          *Money                   `json:"gross_sales_money"`
	TotalTaxMoney            *Money                   `json:"total_tax_money"`
	TotalDiscountMoney       *Money                   `json:"total_discount_money"`
	TotalMoney               *Money                   `json:"total_money"`
}
