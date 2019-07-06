package square

import "time"

type TransactionProduct string

const (
	TransactionProductRegister     TransactionProduct = "REGISTER"
	TransactionProductExternalApi  TransactionProduct = "EXTERNAL_API"
	TransactionProductBilling      TransactionProduct = "BILLING"
	TransactionProductAppointments TransactionProduct = "APPOINTMENTS"
	TransactionProductInvoices     TransactionProduct = "INVOICES"
	TransactionProductOnlineStore  TransactionProduct = "ONLINE_STORE"
	TransactionProductPayroll      TransactionProduct = "PAYROLL"
	TransactionProductOther        TransactionProduct = "OTHER"
)

type Transaction struct {
	Id              string             `json:"id,omitempty"`
	LocationId      string             `json:"location_id,omitempty"`
	CreatedAt       *time.Time         `json:"created_at,omitempty"`
	Tenders         []*Tender          `json:"tenders,omitempty"`
	Refunds         []*Refund          `json:"refunds,omitempty"`
	ReferenceId     string             `json:"reference_id,omitempty"`
	Product         TransactionProduct `json:"product,omitempty"`
	ClientId        string             `json:"client_id,omitempty"`
	ShippingAddress *Address           `json:"shipping_address,omitempty"`
	OrderId         string             `json:"order_id,omitempty"`
}
