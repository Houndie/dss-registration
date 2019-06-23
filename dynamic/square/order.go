package square

import "time"

type OrderState string

const (
	OrderStateOpen      OrderState = "OPEN"
	OrderStateCompleted OrderState = "COMPLETED"
	OrderStateCanceled  OrderState = "CANCELED"
)

type Order struct {
	Id                      string                   `json:"id"`
	LocationID              string                   `json:"location_id"`
	ReferenceID             string                   `json:"reference_id"`
	Source                  *OrderSource             `json:"source"`
	CustomerID              string                   `json:"customer_id"`
	LineItems               []*OrderLineItem         `json:"line_items"`
	Taxes                   []*OrderLineItemTax      `json:"taxes"`
	Discounts               []*OrderLineItemDiscount `json:"discounts"`
	ServiceCharges          []*OrderServiceCharge    `json:"service_charges"`
	Fulfillments            []*OrderFulfillment      `json:"fulfillments"`
	Returns                 []*OrderReturn           `json:"returns"`
	ReturnAmounts           *OrderMoneyAmounts       `json:"return_amounts"`
	NetAmounts              *OrderMoneyAmounts       `json:"net_amounts"`
	RoundingAdjustment      *OrderRoundingAdjustment `json:"rounding_adjustment"`
	Tenders                 *Tender                  `json:"tenders"`
	Refunds                 *Refund                  `json:"refunds"`
	CreatedAt               time.Time                `json:"created_at"`
	UpdatedAt               time.Time                `json:"updated_at"`
	ClosedAt                time.Time                `json:"closed_at"`
	State                   OrderState               `json:"state"`
	TotalMoney              *Money                   `json:"total_money"`
	TotalTaxMoney           *Money                   `json:"total_tax_money"`
	TotalDiscountMoney      *Money                   `json:"total_discount_money"`
	TotalServiceChargeMoney *Money                   `json:"total_service_charge_money"`
}
