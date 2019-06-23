package square

type OrderLineItemTaxType string

const (
	OrderLineItemTaxTypeUnknownTax OrderLineItemTaxType = "UNKNOWN_TAX"
	OrderLineItemTaxTypeAdditive   OrderLineItemTaxType = "ADDITIVE"
	OrderLineItemTaxTypeInclusive  OrderLineItemTaxType = "INCLUSIVE"
)

type OrderLineItemTaxScope string

const (
	OrderLineItemTaxScopeOtherTaxScope OrderLineItemTaxScope = "OTHER_TAX_SCOPE"
	OrderLineItemTaxScopeLineItem      OrderLineItemTaxScope = "LINE_ITEM"
	OrderLineItemTaxScopeOrder         OrderLineItemTaxScope = "ORDER"
)

type OrderLineItemTax struct {
	Uid             string                `json:"uid"`
	CatalogObjectId string                `json:"catalog_object_id"`
	Name            string                `json:"name"`
	Type            OrderLineItemTaxType  `json:"type"`
	Percentage      string                `json:"percentage"`
	AppliedMoney    *Money                `json:"applied_money"`
	Scope           OrderLineItemTaxScope `json:"scope"`
}
