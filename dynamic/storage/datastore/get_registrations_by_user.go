package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
)

func (d *Datastore) GetRegistrationsByUser(ctx context.Context, userId string) ([]*storage.Registration, error) {
	q := datastore.NewQuery(registrationKind).Filter("UserId =", userId).Filter("Disabled =", false)
	t := d.client.Run(ctx, q)
	registrations := []*storage.Registration{}
	for {
		var re *registrationEntity
		key, err := t.Next(re)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errors.Wrap(err, "error fetching registration from datastore")
		}
		r, err := fromRegistrationEntity(key, re)
		if err != nil {
			return nil, fmt.Errorf("error converting registration: %w", err)
		}
		registrations = append(registrations, r)
	}
	return registrations, nil
}
