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

type ErrHasAdminOverride struct {
	Field string
}

func (e ErrHasAdminOverride) Error() string {
	return fmt.Sprintf("%s has payment admin override", e.Field)
}

type ErrHasSquarePayment struct {
	Field string
}

func (e ErrHasSquarePayment) Error() string {
	return fmt.Sprintf("%s square payment cannot be set", e.Field)
}

type ErrHasImmutableField struct {
	Field string
}

func (e ErrHasImmutableField) Error() string {
	return fmt.Sprintf("cannot modify field %s", e.Field)
}
