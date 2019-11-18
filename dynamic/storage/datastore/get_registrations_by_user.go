package datastore

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/registration/listbyuser"
	"github.com/pkg/errors"
	"google.golang.org/api/iterator"
)

func (d *Datastore) GetRegistrationsByUser(ctx context.Context, userId string) ([]*listbyuser.StoreRegistration, error) {
	q := datastore.NewQuery(registrationKind).Filter("UserId =", userId).Filter("Disabled =", false)
	t := d.client.Run(ctx, q)
	registrations := []*listbyuser.StoreRegistration{}
	for {
		var r registrationEntity
		key, err := t.Next(&r)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errors.Wrap(err, "error fetching registration from datastore")
		}
		createdAt, err := time.Parse(time.RFC3339, r.CreatedAt)
		if err != nil {
			return nil, errors.Wrapf(err, "error converting registration created at %s to understandable time", r.CreatedAt)
		}
		registrations = append(registrations, &listbyuser.StoreRegistration{
			Id:        key.Encode(),
			FirstName: r.FirstName,
			LastName:  r.LastName,
			Email:     r.Email,
			CreatedAt: createdAt,
			OrderIds:  r.OrderIds,
		})
	}
	return registrations, nil
}
