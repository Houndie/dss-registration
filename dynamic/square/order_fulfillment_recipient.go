package square

type OrderFulfillmentRecipient struct {
	CustomerId   string `json:"customer_id,omitempty"`
	DisplayName  string `json:"display_name,omitempty"`
	EmailAddress string `json:"email_address,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
}
