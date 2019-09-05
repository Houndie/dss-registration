package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/registration/common"
	"github.com/Houndie/dss-registration/dynamic/registration/update"
	"github.com/pkg/errors"
)

func (s *Datastore) UpdateRegistration(ctx context.Context, r *update.StoreUpdateRegistration, id string) error {
	key, err := datastore.DecodeKey(id)
	if err != nil {
		return errors.Wrap(err, "error decoding registration key")
	}
	registration := registrationEntity{}
	err = s.client.Get(ctx, key, &registration)
	if err != nil {
		return errors.Wrap(err, "error fetching existing registration from database")
	}
	registration.FirstName = r.FirstName
	registration.LastName = r.LastName
	registration.StreetAddress = r.StreetAddress
	registration.City = r.City
	registration.State = r.State
	registration.ZipCode = r.ZipCode
	registration.Email = r.Email
	registration.HomeScene = r.HomeScene
	registration.IsStudent = r.IsStudent
	registration.SoloJazz = r.SoloJazz
	if r.OrderUpdate != nil {
		var newOrderIds []string
		if len(r.OrderUpdate.ObsoleteIds) > 0 {
			newOrderIds = []string{}
			for _, oldOrderId := range registration.OrderIds {
				found := false
				for _, removeOrderId := range r.OrderUpdate.ObsoleteIds {
					if removeOrderId == oldOrderId {
						found = true
						break
					}
				}
				if found {
					continue
				}
				newOrderIds = append(newOrderIds, oldOrderId)
			}
		} else {
			newOrderIds = registration.OrderIds
		}
		newOrderIds = append(newOrderIds, r.OrderUpdate.NewId)
		registration.OrderIds = newOrderIds
	}

	switch p := r.PassType.(type) {
	case *common.WeekendPass:
		registration.WeekendPass = fullWeekendPass
		registration.FullWeekendPassInfo.Level = int(p.Level)
		registration.FullWeekendPassInfo.Tier = int(p.Tier)
	case *common.DanceOnlyPass:
		registration.WeekendPass = danceOnlyPass
	case *common.NoPass:
		registration.WeekendPass = noPass
	default:
		return fmt.Errorf("Found unknown type of weekend pass")
	}

	if r.MixAndMatch != nil {
		registration.HasMixAndMatch = true
		registration.MixAndMatchRole = string(r.MixAndMatch.Role)
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
	case *common.ProvideHousing:
		registration.HousingRequest = providesHousing
		registration.ProvideHousing.Pets = h.Pets
		registration.ProvideHousing.Quantity = h.Quantity
		registration.ProvideHousing.Details = h.Details
	case *common.RequireHousing:
		registration.HousingRequest = requiresHousing
		registration.RequireHousing.PetAllergies = h.PetAllergies
		registration.RequireHousing.Details = h.Details
	case *common.NoHousing:
		registration.HousingRequest = noHousing
	default:
		return fmt.Errorf("Found unknown type of housing")
	}
	_, err = s.client.Put(ctx, key, &registration)
	return errors.Wrap(err, "Error inserting registration into database")
}
