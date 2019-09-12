package adddiscount

type DiscountTarget string

const (
	FullWeekendDiscountTarget     DiscountTarget = "Full Weekend"
	DanceOnlyDiscountTarget       DiscountTarget = "Dance Only"
	MixAndMatchDiscountTarget     DiscountTarget = "Mix And Match"
	SoloJazzDiscountTarget        DiscountTarget = "Solo Jazz"
	TeamCompetitionDiscountTarget DiscountTarget = "Team Competition"
	TShirtDiscountTarget          DiscountTarget = "TShirt"
)

type SingleDiscount struct {
	Name      string
	AppliedTo DiscountTarget
}

type Discount struct {
	Code      string
	Discounts []*SingleDiscount
}

type ErrUnauthorized struct{}

func (ErrUnauthorized) Error() string {
	return "User is not authorized for this operation"
}
