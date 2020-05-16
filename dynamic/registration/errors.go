package registration

import (
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/storage"
)

type ErrOutOfStock struct {
	NextTier storage.WeekendPassTier
	NextCost int
}

func (ErrOutOfStock) Error() string {
	return "this item is out of stock"
}

type ErrUnauthorized struct{}

func (ErrUnauthorized) Error() string {
	return "User is not authorized for this operation"
}

type ErrAlreadyPurchased struct {
	Field         string
	ExistingValue string
}

func (e ErrAlreadyPurchased) Error() string {
	return fmt.Sprintf("cannot update field %s to value %s, as a different value was already purchased", e.Field, e.ExistingValue)
}
