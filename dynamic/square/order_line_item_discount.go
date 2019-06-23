package square

type OrderLineItemDiscountScope string

const (
	OrderLineItemDiscountScopeOtherDiscountScope OrderLineItemDiscountScope = "OTHER_DISCOUNT_SCOPE"
	OrderLineItemDiscountScopeLineItem           OrderLineItemDiscountScope = "LINE_ITEM"
	OrderLineItemDiscountScopeOrder              OrderLineItemDiscountScope = "ORDER"
)

type OrderLineItemDiscountType string

const (
	OrderLineItemDiscountTypeUnknownDiscount    OrderLineItemDiscountType = "UNKNOWN_DISCOUNT"
	OrderLineItemDiscountTypeFixedPercentage    OrderLineItemDiscountType = "FIXED_PERCENTAGE"
	OrderLineItemDiscountTypeFixedAmount        OrderLineItemDiscountType = "FIXED_AMOUNT"
	OrderLineItemDiscountTypeVariablePercentage OrderLineItemDiscountType = "VARIABLE_PERCENTAGE"
	OrderLineItemDiscountTypeVariableAmount     OrderLineItemDiscountType = "VARIABLE_AMOUNT"
)

type OrderLineItemDiscount struct {
	Uid             string                     `json:"uid"`
	CatalogObjectId string                     `json:"catalog_object_id"`
	Name            string                     `json:"name"`
	Type            OrderLineItemDiscountType  `json:"type"`
	Percentage      string                     `json:"percentage"`
	AmountMoney     *Money                     `json:"amount_money"`
	AppliedMoney    *Money                     `json:"applied_money"`
	Scope           OrderLineItemDiscountScope `json:"scope"`
}
