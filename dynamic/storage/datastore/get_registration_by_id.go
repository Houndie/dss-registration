package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/storage"
)

func (d *Datastore) GetRegistration(ctx context.Context, id string) (*storage.Registration, error) {
	re := &registrationEntity{}
	key, err := datastore.DecodeKey(id)
	if err != nil {
		return nil, fmt.Errorf("error decoding datastore key: %w", err)
	}

	err = d.client.Get(ctx, key, re)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			return nil, storage.ErrNotFound{
				Key: id,
			}
		}
		return nil, fmt.Errorf("error fetching registration from database: %w", err)
	}

	r, err := fromRegistrationEntity(key, re)
	if err != nil {
		return nil, fmt.Errorf("error converting registration: %w", err)
	}
	return r, nil
}
