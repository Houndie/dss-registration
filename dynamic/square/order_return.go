package square

type OrderReturn struct {
	Uid                  string                      `json:"uid"`
	SourceOrderId        string                      `json:"source_order_id"`
	ReturnLineItems      []*OrderReturnLineItem      `json:"return_line_items"`
	ReturnServiceCharges []*OrderReturnServiceCharge `json:"return_service_charges"`
	ReturnTaxes          []*OrderReturnTax           `json:"return_taxes"`
	ReturnDiscounts      []*OrderReturnDiscount      `json:"return_discounts"`
	RoundingAdjustment   *OrderRoundingAdjustment    `json:"rounding_adjustment"`
	ReturnAmounts        *OrderMoneyAmounts          `json:"return_amounts"`
}
