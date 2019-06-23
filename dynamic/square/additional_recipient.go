package square

type AdditionalRecipient struct {
	LocationId   string `json:"location_id"`
	Description  string `json:"description"`
	AmountMoney  *Money `json:"amount_money"`
	ReceivableId string `json:"receivable_id"`
}
