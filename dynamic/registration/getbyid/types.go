package getbyid

import (
	"fmt"
	"time"

	"github.com/Houndie/dss-registration/dynamic/registration/common"
)

type OrderItem struct {
	Name string
	Cost int
}

type Order struct {
	Items []*OrderItem
	Paid  bool
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
	Orders          []*Order
	CreatedAt       time.Time
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
}

type ErrBadRegistrationId struct {
	RegistrationId string
}

func (e ErrBadRegistrationId) Error() string {
	return fmt.Sprintf("registration id %s does not correspond to a registration", e.RegistrationId)
}
