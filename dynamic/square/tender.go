package square

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

type tender struct {
	Id                   string                 `json:"id,omitempty"`
	LocationId           string                 `json:"location_id,omitempty"`
	TransactionId        string                 `json:"transaction_id,omitempty"`
	CreatedAt            *time.Time             `json:"created_at,omitempty"`
	Note                 string                 `json:"note,omitempty"`
	AmountMoney          *Money                 `json:"amount_money,omitempty"`
	TipMoney             *Money                 `json:"tip_money,omitempty"`
	ProcessingFeeMoney   *Money                 `json:"processing_fee_money,omitempty"`
	CustomerId           string                 `json:"customer_id,omitempty"`
	Type                 TenderType             `json:"type,omitempty"`
	CardDetails          *TenderCardDetails     `json:"card_details,omitempty"`
	CashDetails          *TenderCashDetails     `json:"cash_details,omitempty"`
	AdditionalRecipients []*AdditionalRecipient `json:"additional_recipients,omitempty"`
}

type tenderType interface {
	isTenderType()
}

type ThirdPartyCardDetails struct{}
type SquareGiftCardDetails struct{}
type NoSaleDetails struct{}
type OtherTenderTypeDetails struct{}

func (*ThirdPartyCardDetails) isTenderType()  {}
func (*SquareGiftCardDetails) isTenderType()  {}
func (*NoSaleDetails) isTenderType()          {}
func (*OtherTenderTypeDetails) isTenderType() {}

type Tender struct {
	Id                   string
	LocationId           string
	TransactionId        string
	CreatedAt            *time.Time
	Note                 string
	AmountMoney          *Money
	TipMoney             *Money
	ProcessingFeeMoney   *Money
	CustomerId           string
	Type                 tenderType
	AdditionalRecipients []*AdditionalRecipient
}

type TenderType string

const (
	tenderTypeCard           TenderType = "CARD"
	tenderTypeCash           TenderType = "CASH"
	tenderTypeThirdPartyCard TenderType = "THIRD_PARTY_CARD"
	tenderTypeSquareGiftCard TenderType = "SQUARE_GIFT_CARD"
	tenderTypeNoSale         TenderType = "NO_SALE"
	tenderTypeOther          TenderType = "OTHER"
)

func (t *Tender) MarshalJSON() ([]byte, error) {
	tJson := tender{
		Id:                   t.Id,
		LocationId:           t.LocationId,
		TransactionId:        t.TransactionId,
		CreatedAt:            t.CreatedAt,
		Note:                 t.Note,
		AmountMoney:          t.AmountMoney,
		TipMoney:             t.TipMoney,
		ProcessingFeeMoney:   t.ProcessingFeeMoney,
		CustomerId:           t.CustomerId,
		AdditionalRecipients: t.AdditionalRecipients,
	}

	switch details := t.Type.(type) {
	case *TenderCardDetails:
		tJson.Type = tenderTypeCard
		tJson.CardDetails = details
	case *TenderCashDetails:
		tJson.Type = tenderTypeCash
		tJson.CashDetails = details
	case *ThirdPartyCardDetails:
		tJson.Type = tenderTypeThirdPartyCard
	case *SquareGiftCardDetails:
		tJson.Type = tenderTypeSquareGiftCard
	case *NoSaleDetails:
		tJson.Type = tenderTypeNoSale
	case *OtherTenderTypeDetails:
		tJson.Type = tenderTypeOther
	default:
		return nil, errors.New("Found unknown tender type")
	}

	return json.Marshal(tJson)
}

func (t *Tender) UnmarshalJSON(b []byte) error {
	tJson := tender{}
	err := json.Unmarshal(b, &tJson)
	if err != nil {
		return errors.Wrap(err, "Error unmarshaling Tender json")
	}

	t.Id = tJson.Id
	t.LocationId = tJson.LocationId
	t.TransactionId = tJson.TransactionId
	t.CreatedAt = tJson.CreatedAt
	t.Note = tJson.Note
	t.AmountMoney = tJson.AmountMoney
	t.TipMoney = tJson.TipMoney
	t.ProcessingFeeMoney = tJson.ProcessingFeeMoney
	t.CustomerId = tJson.CustomerId
	t.AdditionalRecipients = tJson.AdditionalRecipients

	switch tJson.Type {
	case tenderTypeCard:
		t.Type = tJson.CardDetails
	case tenderTypeCash:
		t.Type = tJson.CashDetails
	case tenderTypeThirdPartyCard:
		t.Type = &ThirdPartyCardDetails{}
	case tenderTypeSquareGiftCard:
		t.Type = &SquareGiftCardDetails{}
	case tenderTypeNoSale:
		t.Type = &NoSaleDetails{}
	case tenderTypeOther:
		t.Type = &OtherTenderTypeDetails{}
	default:
		return errors.New("Unknown tender type found")
	}
	return nil
}
