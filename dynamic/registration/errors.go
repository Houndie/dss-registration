package registration

import (
	"errors"
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

type ErrAlreadyPurchased struct {
	Field         string
	ExistingValue string
}

func (e ErrAlreadyPurchased) Error() string {
	return fmt.Sprintf("cannot update field %s to value %s, as a different value was already purchased", e.Field, e.ExistingValue)
}

var ErrRegistrationDisabled = errors.New("registration found when service is not active")

var ErrNoPaidItems = errors.New("registration contains no paid items")
