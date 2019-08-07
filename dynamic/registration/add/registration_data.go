package add

import "github.com/gofrs/uuid"

type WeekendPassLevel int
type WeekendPassTier int

const (
	WeekendPassLevel1 WeekendPassLevel = 1
	WeekendPassLevel2 WeekendPassLevel = 2
	WeekendPassLevel3 WeekendPassLevel = 3

	WeekendPassTier1 WeekendPassTier = 1
	WeekendPassTier2 WeekendPassTier = 2
	WeekendPassTier3 WeekendPassTier = 3
	WeekendPassTier4 WeekendPassTier = 4
	WeekendPassTier5 WeekendPassTier = 5
)

type PassType interface {
	isPassType()
}
type WeekendPass struct {
	Level WeekendPassLevel
	Tier  WeekendPassTier
}
type DanceOnlyPass struct{}
type NoPass struct{}

func (*WeekendPass) isPassType()   {}
func (*DanceOnlyPass) isPassType() {}
func (*NoPass) isPassType()        {}

type MixAndMatchRole string

const (
	MixAndMatchRoleLeader   MixAndMatchRole = "Leader"
	MixAndMatchRoleFollower MixAndMatchRole = "Follower"
)

type MixAndMatch struct {
	Role MixAndMatchRole
}

type TeamCompetition struct {
	Name string
}

type TShirtStyle string

const (
	TShirtStyleUnisexS   TShirtStyle = "Unisex S"
	TShirtStyleUnisexM   TShirtStyle = "Unisex M"
	TShirtStyleUnisexL   TShirtStyle = "Unisex L"
	TShirtStyleUnisexXL  TShirtStyle = "Unisex XL"
	TShirtStyleUnisex2XL TShirtStyle = "Unisex 2XL"
	TShirtStyleUnisex3XL TShirtStyle = "Unisex 3XL"
	TShirtStyleBellaS    TShirtStyle = "Bella S"
	TShirtStyleBellaM    TShirtStyle = "Bella M"
	TShirtStyleBellaL    TShirtStyle = "Bella L"
	TShirtStyleBellaXL   TShirtStyle = "Bella XL"
	TShirtStyleBella2XL  TShirtStyle = "Bella 2XL"
)

type TShirt struct {
	Style TShirtStyle
}

type Housing interface {
	isHousing()
}

type NoHousing struct{}
type ProvideHousing struct {
	Pets     string
	Quantity int
	Details  string
}
type RequireHousing struct {
	PetAllergies string
	Details      string
}

func (*NoHousing) isHousing()      {}
func (*ProvideHousing) isHousing() {}
func (*RequireHousing) isHousing() {}

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
	PassType        PassType
	MixAndMatch     *MixAndMatch
	SoloJazz        bool
	TeamCompetition *TeamCompetition
	TShirt          *TShirt
	Housing         Housing
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
	PassType        PassType
	MixAndMatch     *MixAndMatch
	SoloJazz        bool
	TeamCompetition *TeamCompetition
	TShirt          *TShirt
	Housing         Housing
	ReferenceId     uuid.UUID
	Paid            bool
}
