package getbyid

import (
	"fmt"
	"time"

	"github.com/Houndie/dss-registration/dynamic/registration/common"
)

type UnpaidItems struct {
	OrderIds []string
	Items    []string
	Cost     int
}

type Registration struct {
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
	UnpaidItems     *UnpaidItems
	CreatedAt       time.Time
	Discounts       []*Discount
}

type StoreRegistration struct {
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
	UserId          string
	OrderIds        []string
	CreatedAt       time.Time
	Discounts       []*StoreDiscount
}

type ErrBadRegistrationId struct {
	RegistrationId string
}

func (e ErrBadRegistrationId) Error() string {
	return fmt.Sprintf("registration id %s does not correspond to a registration", e.RegistrationId)
}

type SingleStoreDiscount struct {
	Name      string
	AppliedTo common.PurchaseItem
}

type StoreDiscount struct {
	Code      string
	Discounts []*SingleStoreDiscount
}

type SingleDiscount struct {
	ItemDiscount common.ItemDiscount
	AppliedTo    common.PurchaseItem
}

type Discount struct {
	Code      string
	Discounts []*SingleDiscount
}
