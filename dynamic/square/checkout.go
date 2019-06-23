package square

import "time"

type Checkout struct {
	Id                         string                 `json:"id"`
	CheckoutPageUrl            string                 `json:"checkout_page_url"`
	AskForShippingAddress      bool                   `json:"ask_for_shipping_address"`
	MerchantSupportEmail       string                 `json:"merchant_support_email"`
	PrePopulateBuyerEmail      string                 `json:"pre_populate_buyer_email"`
	PrePopulateShippingAddress *Address               `json:"pre_populate_shipping_address"`
	RedirectUrl                string                 `json:"redirect_url"`
	Order                      *Order                 `json:"order"`
	CreatedAt                  time.Time              `json:"created_at"`
	AdditionalRecipients       []*AdditionalRecipient `json:"additional_recipients"`
}
