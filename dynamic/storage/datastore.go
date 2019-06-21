package storage

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/registration/add"
	"github.com/pkg/errors"
)

type weekendPass struct {
	Level int
}

type mixAndMatch struct {
	Role string
}

type tShirt struct {
	Style string
}

type teamCompetition struct {
	Name string
}

const registrationKind = "Registration"

type registrationEntity struct {
	FirstName       string
	LastName        string
	StreetAddress   string
	City            string
	State           string
	ZipCode         string
	Email           string
	HomeScene       string
	IsStudent       bool
	SoloJazz        bool
	RequiresHousing bool
	RequireHousing  struct {
		PetAllergies string
		Details      string
	}
	ProvidesHousing bool
	ProvideHousing  struct {
		Pets     string
		Quantity int
		Details  string
	}
	WantsTShirt          bool
	TShirtStyle          string
	HasTeamCompetition   bool
	TeamCompetitionName  string
	HasMixAndMatch       bool
	MixAndMatchRole      string
	HasFullWeekendPass   bool
	FullWeekendPassLevel int
	HasDanceOnlyPass     bool
}

type Datastore struct {
	client *datastore.Client
}

func NewDatastore(client *datastore.Client) *Datastore {
	return &Datastore{
		client: client,
	}

}

func (s *Datastore) AddRegistration(ctx context.Context, r *add.Registration) error {
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
	}

	switch p := r.PassType.(type) {
	case *add.WeekendPass:
		registration.HasFullWeekendPass = true
		registration.FullWeekendPassLevel = int(p.Level)
	case *add.DanceOnlyPass:
		registration.HasDanceOnlyPass = true
	case *add.NoPass: //Do nothing
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
		registration.ProvidesHousing = true
		registration.ProvideHousing.Pets = h.Pets
		registration.ProvideHousing.Quantity = h.Quantity
		registration.ProvideHousing.Details = h.Details
	case *add.RequireHousing:
		registration.RequiresHousing = true
		registration.RequireHousing.PetAllergies = h.PetAllergies
		registration.RequireHousing.Details = h.Details
	case *add.NoHousing: //Nothing to do
	default:
		return fmt.Errorf("Found unknown type of housing")
	}
	registrationKey := datastore.IncompleteKey(registrationKind, nil)
	_, err := s.client.Put(ctx, registrationKey, registration)
	return errors.Wrap(err, "Error inserting registration into database")
}
