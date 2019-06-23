package square

type CardBrand string

const (
	CardBrandOtherBrand      CardBrand = "OTHER_BRAND"
	CardBrandVisa            CardBrand = "VISA"
	CardBrandMastercard      CardBrand = "MASTERCARD"
	CardBrandAmericanExpress CardBrand = "AMERICAN_EXPRESS"
	CardBrandDiscover        CardBrand = "DISCOVER"
	CardBrandDiscoverDiners  CardBrand = "DINERS"
	CardBrandJcb             CardBrand = "JCB"
	CardBrandChinaUnionpay   CardBrand = "CHINA_UNIONPAY"
	CardBrandSquareGiftCard  CardBrand = "SQUARE_GIFT_CARD"
)

type Card struct {
	Id             string    `json:"string"`
	CardBrand      CardBrand `json:"card_brand"`
	Last4          string    `json:"last_4"`
	ExpMonth       int       `json:"exp_month"`
	ExpYear        int       `json:"exp_year"`
	CardholderName string    `json:"cardholder_name"`
	BillingAddress Address   `json:"address"`
	Fingerprint    string    `json:"fingerprint"`
}
