package square

type OrderRoundingAdjustment struct {
	Uid         string `json:"uid"`
	Name        string `json:"name"`
	AmountMoney *Money `json:"amount_money"`
}
