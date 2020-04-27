package storage

import "fmt"

type SingleDiscount struct {
	Name      string
	AppliedTo PurchaseItem
}

type Discount struct {
	Code      string
	Discounts []*SingleDiscount
}

type ErrDiscountDoesNotExist struct {
	Code string
}

func (e ErrDiscountDoesNotExist) Error() string {
	return fmt.Sprintf("discount for code %s does not exist", e.Code)
}
