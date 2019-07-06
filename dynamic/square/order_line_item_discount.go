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
	Uid             string                     `json:"uid,omitempty"`
	CatalogObjectId string                     `json:"catalog_object_id,omitempty"`
	Name            string                     `json:"name,omitempty"`
	Type            OrderLineItemDiscountType  `json:"type,omitempty"`
	Percentage      string                     `json:"percentage,omitempty"`
	AmountMoney     *Money                     `json:"amount_money,omitempty"`
	AppliedMoney    *Money                     `json:"applied_money,omitempty"`
	Scope           OrderLineItemDiscountScope `json:"scope,omitempty"`
}
