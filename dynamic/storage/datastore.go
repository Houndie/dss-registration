package storage

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/registration/add"
)

const danceOnlyPassKind = "DanceOnlyPass"

type danceOnlyPassEntity struct{}

const fullWeekendPassKind = "FullWeekendPass"

type fullWeekendPassEntity struct {
	Level int
}

const mixAndMatchKind = "MixAndMatch"

type mixAndMatchEntity struct {
	Role string
}

const teamCompetitionKind = "TeamCompetition"

type teamCompetitionEntity struct {
	Name string
}

const tShirtKind = "TShirt"

type tShirtEntity struct {
	Style string
}

const provideHousingKind = "Provide Housing"

type provideHousingEntity struct {
	Pets     string
	Quantity int
	Details  string
}

const requireHousingKind = "Require Housing"

type requireHousingEntity struct {
	PetAllergies string
	Details      string
}

const registrationKind = "Registration"

type registrationEntity struct {
	FirstName     string
	LastName      string
	StreetAddress string
	City          string
	State         string
	ZipCode       string
	Email         string
	HomeScene     string
	IsStudent     bool
	SoloJazz      bool
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
	// Inserting the main registration table can't be in a transaction, as we need the key
	registrationKey := datastore.IncompleteKey(registrationKind, nil)
	registrationKey, err := s.client.Put(ctx, registrationKey, &registrationEntity{
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
	})
	if err != nil {
		return err
	}
	_, err = s.client.RunInTransaction(ctx, func(tx *datastore.Transaction) error {

		switch p := r.PassType.(type) {
		case *add.WeekendPass:
			weekendPassKey := datastore.IncompleteKey(fullWeekendPassKind, registrationKey)
			_, err := tx.Put(weekendPassKey, &fullWeekendPassEntity{
				Level: int(p.Level),
			})
			if err != nil {
				return err
			}
		case *add.DanceOnlyPass:
			danceOnlyPassKey := datastore.IncompleteKey(danceOnlyPassKind, registrationKey)
			_, err := tx.Put(danceOnlyPassKey, &danceOnlyPassEntity{})
			if err != nil {
				return err
			}
		case *add.NoPass: //Do nothing
		default:
			return fmt.Errorf("Found unknown type of weekend pass")
		}

		if r.MixAndMatch != nil {
			mixAndMatchKey := datastore.IncompleteKey(mixAndMatchKind, registrationKey)
			_, err := tx.Put(mixAndMatchKey, &mixAndMatchEntity{
				Role: r.MixAndMatch.Role,
			})
			if err != nil {
				return err
			}
		}

		if r.TeamCompetition != nil {
			teamCompetitionKey := datastore.IncompleteKey(teamCompetitionKind, registrationKey)
			_, err := tx.Put(teamCompetitionKey, &teamCompetitionEntity{
				Name: r.TeamCompetition.Name,
			})
			if err != nil {
				return err
			}
		}

		if r.TShirt != nil {
			tShirtKey := datastore.IncompleteKey(tShirtKind, registrationKey)
			_, err := tx.Put(tShirtKey, &tShirtEntity{
				Style: string(r.TShirt.Style),
			})
			if err != nil {
				return err
			}
		}

		switch h := r.Housing.(type) {
		case *add.ProvideHousing:
			provideHousingKey := datastore.IncompleteKey(provideHousingKind, registrationKey)
			_, err := tx.Put(provideHousingKey, &provideHousingEntity{
				Pets:     h.Pets,
				Quantity: h.Quantity,
				Details:  h.Details,
			})
			if err != nil {
				return err
			}
		case *add.RequireHousing:
			requireHousingKey := datastore.IncompleteKey(requireHousingKind, registrationKey)
			_, err := tx.Put(requireHousingKey, &requireHousingEntity{
				PetAllergies: h.PetAllergies,
				Details:      h.Details,
			})
			if err != nil {
				return err
			}
		case *add.NoHousing: //Nothing to do
		default:
			return fmt.Errorf("Found unknown type of housing")
		}
		return nil
	})
	if err != nil {
		_ = s.client.Delete(ctx, registrationKey)
	}
	return err
}
