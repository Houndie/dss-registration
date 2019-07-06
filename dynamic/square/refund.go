package square

import "time"

type RefundStatus string

const (
	RefundStatusPending  RefundStatus = "PENDING"
	RefundStatusApproved RefundStatus = "APPROVED"
	RefundStatusRejected RefundStatus = "REJECTED"
	RefundStatusFailed   RefundStatus = "FAILED"
)

type Refund struct {
	Id                   string                 `json:"id,omitempty"`
	LocationId           string                 `json:"location_id,omitempty"`
	TransactionId        string                 `json:"transaction_id,omitempty"`
	TenderId             string                 `json:"tender_id,omitempty"`
	CreatedAt            *time.Time             `json:"created_at,omitempty"`
	Reason               string                 `json:"reason,omitempty"`
	AmountMoney          *Money                 `json:"amount_money,omitempty"`
	Status               RefundStatus           `json:"status,omitempty"`
	ProcessingFeeMoney   *Money                 `json:"processing_fee_money,omitempty"`
	AdditionalRecipients []*AdditionalRecipient `json:"additional_recipients,omitempty"`
}
