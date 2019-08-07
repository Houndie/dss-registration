package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/pkg/errors"
)

func (d *Datastore) DeleteRegistration(ctx context.Context, key string) error {
	dKey, err := datastore.DecodeKey(key)
	if err != nil {
		return errors.Wrap(err, "error decoding key to datastore type")
	}

	err = d.client.Delete(ctx, dKey)
	return errors.Wrap(err, "error deleting registration from datastore")
}
