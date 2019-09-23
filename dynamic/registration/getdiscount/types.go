package getdiscount

import (
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/registration/common"
)

type ErrDiscountDoesNotExist struct {
	Code string
}

func (e ErrDiscountDoesNotExist) Error() string {
	return fmt.Sprintf("discount for code %s does not exist", e.Code)
}

type Discount struct {
	AppliedTo    common.PurchaseItem
	ItemDiscount common.ItemDiscount
}
