package datastore

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/storage"
)

func (s *Datastore) UpdateRegistration(ctx context.Context, r *storage.Registration) error {
	key, re, err := toRegistrationEntity(r)
	if err != nil {
		return fmt.Errorf("error parsing registration: %w", err)
	}
	_, err = s.client.Put(ctx, key, re)
	return fmt.Errorf("Error inserting registration into database", err)
}
