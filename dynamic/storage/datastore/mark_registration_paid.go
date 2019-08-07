package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

func (d *Datastore) MarkRegistrationPaid(ctx context.Context, referenceId uuid.UUID) error {
	var registrations []*registrationEntity
	q := datastore.NewQuery(registrationKind).Filter("ReferenceId = ", referenceId.String()).Limit(1)
	keys, err := d.client.GetAll(ctx, q, &registrations)
	if err != nil {
		return errors.Wrap(err, "error retrieving registration from datastore")
	}
	if len(registrations) != 1 {
		return fmt.Errorf("found incorrect number of registrations %d", len(registrations))
	}
	if len(keys) != len(registrations) {
		return fmt.Errorf("found incorrect number of keys %d, expected %d", len(keys), len(registrations))
	}

	if registrations[0].Paid {
		return storage.ErrAlreadyPaid{}
	}

	registrations[0].Paid = true
	_, err = d.client.Put(ctx, keys[0], registrations[0])
	return errors.Wrap(err, "error putting registration back in datastore")
}
