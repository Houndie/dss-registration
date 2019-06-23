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
	Id              string             `json:"id"`
	LocationId      string             `json:"location_id"`
	CreatedAt       time.Time          `json:"created_at"`
	Tenders         []*Tender          `json:"tenders"`
	Refunds         []*Refund          `json:"refunds"`
	ReferenceId     string             `json:"reference_id"`
	Product         TransactionProduct `json:"product"`
	ClientId        string             `json:"client_id"`
	ShippingAddress *Address           `json:"shipping_address"`
	OrderId         string             `json:"order_id"`
}
