package datastore

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/registration/add"
	"github.com/Houndie/dss-registration/dynamic/registration/common"
	"github.com/pkg/errors"
)

func (s *Datastore) AddRegistration(ctx context.Context, r *add.StoreRegistration) (string, error) {
	keys := make([]*datastore.Key, len(r.Discounts))

	for i, discount := range r.Discounts {
		d, err := datastore.DecodeKey(discount)
		if err != nil {
			return "", fmt.Errorf("Key not of datastore type")
		}
		keys[i] = d
	}
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
		UserId:        r.UserId,
		OrderIds:      r.OrderIds,
		CreatedAt:     time.Now().Format(time.RFC3339),
		Disabled:      false,
		Discounts:     keys,
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
		return "", fmt.Errorf("Found unknown type of weekend pass")
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
		return "", fmt.Errorf("Found unknown type of housing")
	}
	registrationKey := datastore.IncompleteKey(registrationKind, nil)
	completeKey, err := s.client.Put(ctx, registrationKey, registration)
	if err != nil {
		return "", errors.Wrap(err, "Error inserting registration into database")
	}
	return completeKey.Encode(), nil
}
