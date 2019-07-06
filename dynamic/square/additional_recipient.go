package square

type AdditionalRecipient struct {
	LocationId   string `json:"location_id,omitempty"`
	Description  string `json:"description,omitempty"`
	AmountMoney  *Money `json:"amount_money,omitempty"`
	ReceivableId string `json:"receivable_id,omitempty"`
}
