package getdiscount

import (
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/registration/common"
)

type ItemDiscount interface {
	isItemDiscount()
}

type PercentDiscount struct {
	Amount string
}

type DollarDiscount struct {
	Amount int
}

func (*DollarDiscount) isItemDiscount()  {}
func (*PercentDiscount) isItemDiscount() {}

type Discount struct {
	AppliedTo    common.DiscountTarget
	ItemDiscount ItemDiscount
}

type StoreDiscount struct {
	Name      string
	AppliedTo common.DiscountTarget
}

type ErrDiscountDoesNotExist struct {
	Code string
}

func (e ErrDiscountDoesNotExist) Error() string {
	return fmt.Sprintf("discount for code %s does not exist", e.Code)
}
