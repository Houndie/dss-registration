package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/registration/common"
	"github.com/Houndie/dss-registration/dynamic/registration/update"
	"github.com/pkg/errors"
)

func (d *Datastore) GetUpdateRegistration(ctx context.Context, id string) (*update.StoreOldRegistration, error) {
	key, err := datastore.DecodeKey(id)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding key")
	}
	registration := registrationEntity{}
	err = d.client.Get(ctx, key, &registration)
	if err != nil {
		return nil, errors.Wrap(err, "error fetching registration from datastore")
	}

	var passType common.PassType
	switch registration.WeekendPass {
	case fullWeekendPass:
		passType = &common.WeekendPass{
			Level: common.WeekendPassLevel(registration.FullWeekendPassInfo.Level),
			Tier:  common.WeekendPassTier(registration.FullWeekendPassInfo.Tier),
		}
	case danceOnlyPass:
		passType = &common.DanceOnlyPass{}
	case noPass:
		passType = &common.NoPass{}
	}

	var mixAndMatch *common.MixAndMatch
	if registration.HasMixAndMatch {
		mixAndMatch = &common.MixAndMatch{
			Role: common.MixAndMatchRole(registration.MixAndMatchRole),
		}
	}

	var teamCompetition *common.TeamCompetition
	if registration.HasTeamCompetition {
		teamCompetition = &common.TeamCompetition{
			Name: registration.TeamCompetitionName,
		}
	}

	var tShirt *common.TShirt
	if registration.WantsTShirt {
		tShirt = &common.TShirt{
			Style: common.TShirtStyle(registration.TShirtStyle),
		}
	}

	return &update.StoreOldRegistration{
		IsStudent:       registration.IsStudent,
		PassType:        passType,
		MixAndMatch:     mixAndMatch,
		SoloJazz:        registration.SoloJazz,
		TeamCompetition: teamCompetition,
		TShirt:          tShirt,
		UserId:          registration.UserId,
	}, nil
}
