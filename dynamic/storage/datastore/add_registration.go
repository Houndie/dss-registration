package datastore

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/storage"
)

func (s *Datastore) AddRegistration(ctx context.Context, r *storage.Registration) (string, error) {
	r.CreatedAt = time.Now()
	_, re, err := toRegistrationEntity(r)
	if err != nil {
		return "", fmt.Errorf("error parsing registration: %w", err)
	}
	registrationKey := datastore.IncompleteKey(registrationKind, nil)
	completeKey, err := s.client.Put(ctx, registrationKey, re)
	if err != nil {
		return "", fmt.Errorf("Error inserting registration into database: %w", err)
	}
	return completeKey.Encode(), nil
}
