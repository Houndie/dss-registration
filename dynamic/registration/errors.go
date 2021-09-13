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

var ErrNoPurchaseItems = errors.New("registration contains no purchase items")

var ErrNoUnpaidItems = errors.New("registration contains no unpaid items")

type ErrFileTooBig struct {
	Filesize int64
}

func (e ErrFileTooBig) Error() string {
	return fmt.Sprintf("filesize %d is too big", e.Filesize)
}
