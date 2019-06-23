package square

type OrderMoneyAmounts struct {
	TotalMoney         *Money `json:"total_money"`
	TaxMoney           *Money `json:"tax_money"`
	DiscountMoney      *Money `json:"discount_money"`
	TipMoney           *Money `json:"tip_money"`
	ServiceChargeMoney *Money `json:"service_charge_money"`
}
