package adddiscount

import "github.com/Houndie/dss-registration/dynamic/registration/common"

type SingleDiscount struct {
	Name      string
	AppliedTo common.PurchaseItem
}

type Discount struct {
	Code      string
	Discounts []*SingleDiscount
}

type ErrUnauthorized struct{}

func (ErrUnauthorized) Error() string {
	return "User is not authorized for this operation"
}
