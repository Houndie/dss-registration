package storage

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/registration/add"
	"github.com/pkg/errors"
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
	TransactionID string
	Paid          bool
}

type Datastore struct {
	client *datastore.Client
}

func NewDatastore(client *datastore.Client) *Datastore {
	return &Datastore{
		client: client,
	}

}

func (s *Datastore) AddRegistration(ctx context.Context, r *add.StoreRegistration) error {
	registration := &registrationEntity{
		FirstName:     r.FirstName,
		LastName:      r.LastName,
		StreetAddress: r.StreetAddress,
		City:          r.City,
		State:         r.State,
		ZipCode:       r.ZipCode,
		Email:         r.Email,
		HomeScene:     r.HomeScene,
		IsStudent:     r.IsStudent,
		SoloJazz:      r.SoloJazz,
		TransactionID: r.TransactionID.String(),
		Paid:          false,
	}

	switch p := r.PassType.(type) {
	case *add.WeekendPass:
		registration.WeekendPass = fullWeekendPass
		registration.FullWeekendPassInfo.Level = int(p.Level)
		registration.FullWeekendPassInfo.Tier = int(p.Tier)
	case *add.DanceOnlyPass:
		registration.WeekendPass = danceOnlyPass
	case *add.NoPass:
		registration.WeekendPass = noPass
	default:
		return fmt.Errorf("Found unknown type of weekend pass")
	}

	if r.MixAndMatch != nil {
		registration.HasMixAndMatch = true
		registration.MixAndMatchRole = r.MixAndMatch.Role
	}

	if r.TeamCompetition != nil {
		registration.HasTeamCompetition = true
		registration.TeamCompetitionName = r.TeamCompetition.Name
	}

	if r.TShirt != nil {
		registration.WantsTShirt = true
		registration.TShirtStyle = string(r.TShirt.Style)
	}

	switch h := r.Housing.(type) {
	case *add.ProvideHousing:
		registration.HousingRequest = providesHousing
		registration.ProvideHousing.Pets = h.Pets
		registration.ProvideHousing.Quantity = h.Quantity
		registration.ProvideHousing.Details = h.Details
	case *add.RequireHousing:
		registration.HousingRequest = requiresHousing
		registration.RequireHousing.PetAllergies = h.PetAllergies
		registration.RequireHousing.Details = h.Details
	case *add.NoHousing:
		registration.HousingRequest = noHousing
	default:
		return fmt.Errorf("Found unknown type of housing")
	}
	registrationKey := datastore.IncompleteKey(registrationKind, nil)
	_, err := s.client.Put(ctx, registrationKey, registration)
	return errors.Wrap(err, "Error inserting registration into database")
}
