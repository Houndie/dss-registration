package add

import "github.com/gofrs/uuid"

type WeekendPassLevel int
type WeekendPassTier int

const (
	Level1 WeekendPassLevel = 1
	Level2 WeekendPassLevel = 2
	Level3 WeekendPassLevel = 3

	Tier1 WeekendPassTier = 1
	Tier2 WeekendPassTier = 2
	Tier3 WeekendPassTier = 3
	Tier4 WeekendPassTier = 4
	Tier5 WeekendPassTier = 5
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

type MixAndMatch struct {
	Role string
}

type TeamCompetition struct {
	Name string
}

type TShirtStyle string

const (
	UnisexS   TShirtStyle = "Unisex S"
	UnisexM   TShirtStyle = "Unisex M"
	UnisexL   TShirtStyle = "Unisex L"
	UnisexXL  TShirtStyle = "Unisex XL"
	Unisex2XL TShirtStyle = "Unisex 2XL"
	Unisex3XL TShirtStyle = "Unisex 3XL"
	BellaS    TShirtStyle = "Bella S"
	BellaM    TShirtStyle = "Bella M"
	BellaL    TShirtStyle = "Bella L"
	BellaXL   TShirtStyle = "Bella XL"
	Bella2XL  TShirtStyle = "Bella 2XL"
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
	TransactionID   uuid.UUID
}
