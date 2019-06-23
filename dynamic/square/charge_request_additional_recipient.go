package square

type ChargeRequestAdditionalRecipient struct {
	LocationId  string `json:"location_id"`
	Description string `json:"description"`
	AmountMoney *Money `json:"amount_money"`
}
