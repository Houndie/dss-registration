package square

type OrderFulfillmentRecipient struct {
	CustomerId   string `json:"customer_id"`
	DisplayName  string `json:"display_name"`
	EmailAddress string `json:"email_address"`
	PhoneNumber  string `json:"phone_number"`
}
