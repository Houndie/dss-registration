package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/registration/common"
	"github.com/pkg/errors"
)

func (s *Datastore) UpdateRegistrationTier(ctx context.Context, id string, newTier common.WeekendPassTier) error {
	key, err := datastore.DecodeKey(id)
	if err != nil {
		return errors.Wrap(err, "error decoding registration key")
	}
	registration := registrationEntity{}
	err = s.client.Get(ctx, key, &registration)
	if err != nil {
		return errors.Wrap(err, "error fetching existing registration from database")
	}
	registration.FullWeekendPassInfo.Tier = int(newTier)
	_, err = s.client.Put(ctx, key, &registration)
	return errors.Wrap(err, "Error inserting registration into database")
}
