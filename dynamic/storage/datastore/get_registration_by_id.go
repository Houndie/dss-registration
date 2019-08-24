package datastore

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/registration/common"
	"github.com/Houndie/dss-registration/dynamic/registration/getbyid"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/pkg/errors"
)

func (d *Datastore) GetRegistrationById(ctx context.Context, id string) (*getbyid.StoreRegistration, error) {
	r := registrationEntity{}
	key, err := datastore.DecodeKey(id)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding datastore key")
	}

	err = d.client.Get(ctx, key, &r)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			return nil, storage.ErrNotFound{
				Key: id,
			}
		}
		return nil, errors.Wrap(err, "error fetching registration from database")
	}

	createdAt, err := time.Parse(time.RFC3339, r.CreatedAt)
	if err != nil {
		return nil, errors.Wrapf(err, "error converting registration created at %s to understandable time", r.CreatedAt)
	}

	var passType common.PassType
	switch r.WeekendPass {
	case fullWeekendPass:
		passType = &common.WeekendPass{
			Level: common.WeekendPassLevel(r.FullWeekendPassInfo.Level),
			Tier:  common.WeekendPassTier(r.FullWeekendPassInfo.Tier),
		}
	case danceOnlyPass:
		passType = &common.DanceOnlyPass{}
	case noPass:
		passType = &common.NoPass{}
	}

	var mixAndMatch *common.MixAndMatch
	if r.HasMixAndMatch {
		mixAndMatch = &common.MixAndMatch{
			Role: common.MixAndMatchRole(r.MixAndMatchRole),
		}
	}

	var teamCompetition *common.TeamCompetition
	if r.HasTeamCompetition {
		teamCompetition = &common.TeamCompetition{
			Name: r.TeamCompetitionName,
		}
	}

	var tShirt *common.TShirt
	if r.WantsTShirt {
		tShirt = &common.TShirt{
			Style: common.TShirtStyle(r.TShirtStyle),
		}
	}

	var housing common.Housing
	switch r.HousingRequest {
	case requiresHousing:
		housing = &common.RequireHousing{
			PetAllergies: r.RequireHousing.PetAllergies,
			Details:      r.RequireHousing.Details,
		}
	case providesHousing:
		housing = &common.ProvideHousing{
			Pets:     r.ProvideHousing.Pets,
			Quantity: r.ProvideHousing.Quantity,
			Details:  r.ProvideHousing.Details,
		}
	case noHousing:
		housing = &common.NoHousing{}
	}

	return &getbyid.StoreRegistration{
		FirstName:       r.FirstName,
		LastName:        r.LastName,
		StreetAddress:   r.StreetAddress,
		City:            r.City,
		State:           r.State,
		ZipCode:         r.ZipCode,
		Email:           r.Email,
		HomeScene:       r.HomeScene,
		IsStudent:       r.IsStudent,
		PassType:        passType,
		MixAndMatch:     mixAndMatch,
		SoloJazz:        r.SoloJazz,
		TeamCompetition: teamCompetition,
		TShirt:          tShirt,
		Housing:         housing,
		UserId:          r.UserId,
		OrderIds:        r.OrderIds,
		CreatedAt:       createdAt,
	}, nil
}
