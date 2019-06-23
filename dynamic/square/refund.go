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
	Id                   string                 `json:"id"`
	LocationId           string                 `json:"location_id"`
	TransactionId        string                 `json:"transaction_id"`
	TenderId             string                 `json:"tender_id"`
	CreatedAt            time.Time              `json:"created_at"`
	Reason               string                 `json:"reason"`
	AmountMoney          *Money                 `json:"amount_money"`
	Status               RefundStatus           `json:"status"`
	ProcessingFeeMoney   *Money                 `json:"processing_fee_money"`
	AdditionalRecipients []*AdditionalRecipient `json:"additional_recipients"`
}
