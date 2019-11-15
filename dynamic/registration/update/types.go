package update

import (
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/registration/common"
)

type Registration struct {
	Id              string
	FirstName       string
	LastName        string
	StreetAddress   string
	City            string
	State           string
	ZipCode         string
	Email           string
	HomeScene       string
	IsStudent       bool
	PassType        common.PassType
	MixAndMatch     *common.MixAndMatch
	SoloJazz        bool
	TeamCompetition *common.TeamCompetition
	TShirt          *common.TShirt
	Housing         common.Housing
	DiscountCodes   []string
}

type StoreOldRegistration struct {
	IsStudent       bool
	PassType        common.PassType
	MixAndMatch     *common.MixAndMatch
	SoloJazz        bool
	TeamCompetition *common.TeamCompetition
	TShirt          *common.TShirt
	UserId          string
	OrderIds        []string
	Discounts       []*ExistingDiscount
}

type ExistingDiscount struct {
	Code      string
	Discounts []*common.StoreDiscount
}

type StoreOrderUpdate struct {
	NewId       string
	ObsoleteIds []string
}

type StoreUpdateRegistration struct {
	FirstName        string
	LastName         string
	StreetAddress    string
	City             string
	State            string
	ZipCode          string
	Email            string
	HomeScene        string
	IsStudent        bool
	PassType         common.PassType
	MixAndMatch      *common.MixAndMatch
	SoloJazz         bool
	TeamCompetition  *common.TeamCompetition
	TShirt           *common.TShirt
	Housing          common.Housing
	NewOrderId       string
	ObsoleteOrderIds []string
	NewDiscounts     []string
}

type ErrBadRegistrationId struct {
	RegistrationId string
}

func (e ErrBadRegistrationId) Error() string {
	return fmt.Sprintf("registration id %s does not correspond to a registration", e.RegistrationId)
}

type ErrAlreadyPurchased struct {
	Field         string
	ExistingValue string
}

func (e ErrAlreadyPurchased) Error() string {
	return fmt.Sprintf("cannot update field %s to value %s, as a different value was already purchased", e.Field, e.ExistingValue)
}

type ErrDiscountAlreadyApplied struct {
	Code string
}

func (e ErrDiscountAlreadyApplied) Error() string {
	return fmt.Sprintf("discount with code %s already applied to this registration", e.Code)
}

type ErrOutOfStock struct {
	NextTier int
	NextCost int
}

func (ErrOutOfStock) Error() string {
	return "this item is out of stock"
}
