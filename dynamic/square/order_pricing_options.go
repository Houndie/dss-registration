package square

type OrderPricingOptions struct {
	AutoApplyDiscounts bool `json:"auto_apply_discounts,omitempty"`
	AutoApplyTaxes     bool `json:"auto_apply_taxes,omitempty"`
}
