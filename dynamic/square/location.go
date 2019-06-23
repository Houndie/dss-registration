package square

import "time"

type LocationCapability string

const LocationCapabilityCreditCardProcessing LocationCapability = "CREDIT_CARD_PROCESSING"

type LocationStatus string

const (
	LocationStatusActive   LocationStatus = "ACTIVE"
	LocationStatusInactive LocationStatus = "INACTIVE"
)

type LocationType string

const (
	LocationTypePhysical LocationStatus = "PHYSICAL"
	LocationTypeMobile   LocationStatus = "MOBILE"
)

type Location struct {
	Id                string               `json:"id"`
	Name              string               `json:"name"`
	Address           *Address             `json:"address"`
	Timezone          string               `json:"timezone"`
	Capabilities      []LocationCapability `json:"capabilities"`
	Status            LocationStatus       `json:"status"`
	CreatedAt         time.Time            `json:"created_at"`
	MerchantId        string               `json:"merchant_id"`
	Country           string               `json:"country"`
	LanguageCode      string               `json:"language_code"`
	Currency          string               `json:"currency"`
	PhoneNumber       string               `json:"phone_number"`
	BusinessName      string               `json:"business_name"`
	Type              LocationType         `json:"type"`
	WebsiteUrl        string               `json:"website_url"`
	BusinessHours     *BusinessHours       `json:"business_hours"`
	BusinessEmail     string               `json:"business_email"`
	Description       string               `json:"description"`
	TwitterUsername   string               `json:"twitter_username"`
	InstagramUsername string               `json:"instagram_username"`
	FacebookUrl       string               `json:"facebook_url"`
	Coordinates       *Coordinates         `json:"coordinates"`
}
