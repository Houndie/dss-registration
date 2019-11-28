package datastore

import (
	"time"

	"cloud.google.com/go/datastore"
)

const (
	registrationKind = "Registration"

	fullWeekendPass = "Full Weekend Pass"
	danceOnlyPass   = "Dance Only Pass"
	noPass          = "No Pass"

	requiresHousing = "Requires Housing"
	providesHousing = "Provides Housing"
	noHousing       = "No Housing"
)

type registrationEntity struct {
	FirstName      string
	LastName       string
	StreetAddress  string
	City           string
	State          string
	ZipCode        string
	Email          string
	HomeScene      string
	IsStudent      bool
	SoloJazz       bool
	HousingRequest string
	RequireHousing struct {
		PetAllergies string
		Details      string
	}
	ProvideHousing struct {
		Pets     string
		Quantity int
		Details  string
	}
	WantsTShirt         bool
	TShirtStyle         string
	HasTeamCompetition  bool
	TeamCompetitionName string
	HasMixAndMatch      bool
	MixAndMatchRole     string
	WeekendPass         string
	FullWeekendPassInfo struct {
		Level int
		Tier  int
	}
	UserId    string
	OrderIds  []string
	CreatedAt string
	Discounts []*datastore.Key
	Disabled  bool
}

const (
	discountKind = "Discount"

	fullWeekendDiscount     = "full weekend"
	danceOnlyDiscount       = "dance only"
	mixAndMatchDiscount     = "mix and match"
	soloJazzDiscount        = "solo jazz"
	teamCompetitionDiscount = "team competition"
	tshirtDiscount          = "tshirt"
)

type singleDiscount struct {
	Name      string
	AppliedTo string
}

type discountEntity struct {
	Code      string
	Discounts []singleDiscount `datastore:",noindex"`
}

const adminKind = "Admin"

type adminEntity struct {
	UserId string
}

const (
	volunteerKind = "Volunteer"
)

type volunteerEntity struct {
	UserId   string
	Name     string
	Email    string
	Disabled bool
	Created  time.Time
}
