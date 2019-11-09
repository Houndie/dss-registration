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

	discounts := make([]*update.ExistingDiscount, len(registration.Discounts))
	for i, discountKey := range registration.Discounts {
		de := discountEntity{}
		err := d.client.Get(ctx, discountKey, &de)
		if err != nil {
			return nil, errors.Wrapf(err, "error getting discount with key %v", discountKey)
		}
		singleDiscounts := make([]*common.StoreDiscount, len(de.Discounts))
		for j, sd := range de.Discounts {
			appliedTo, err := parseAppliedTo(sd.AppliedTo)
			if err != nil {
				return nil, errors.Wrap(nil, "found unknown error appliedto from store")
			}
			singleDiscounts[j] = &common.StoreDiscount{
				Name:      sd.Name,
				AppliedTo: appliedTo,
			}
		}
		discounts[i] = &update.ExistingDiscount{
			Code:      de.Code,
			Discounts: singleDiscounts,
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
		OrderIds:        registration.OrderIds,
		Discounts:       discounts,
	}, nil
}
