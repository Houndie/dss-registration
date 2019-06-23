package square

type TenderCashDetails struct {
	BuyerTenderedMoney *Money `json:"buyer_tendered_money"`
	ChangeBackMoney    *Money `json:"change_back_money"`
}

func (*TenderCashDetails) isTenderType() {}
