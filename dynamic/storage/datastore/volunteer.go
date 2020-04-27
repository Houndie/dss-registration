package datastore

import (
	"fmt"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/storage"
)

const (
	volunteerKind = "Volunteer"
)

type volunteerEntity struct {
	UserID    string
	Name      string
	Email     string
	Disabled  bool
	CreatedAt time.Time
}

func toVolunteerEntity(v *storage.Volunteer) (*datastore.Key, *volunteerEntity, error) {

	var key *datastore.Key
	if v.ID != "" {
		var err error
		key, err = datastore.DecodeKey(v.ID)
		if err != nil {
			return nil, nil, fmt.Errorf("error decoding volunteer ID: %w", err)
		}
	}

	return key, &volunteerEntity{
		UserID:    v.UserID,
		Name:      v.Name,
		Email:     v.Email,
		Disabled:  v.Disabled,
		CreatedAt: v.CreatedAt,
	}, nil
}

func fromVolunteerEntity(key *datastore.Key, ve *volunteerEntity) *storage.Volunteer {
	return &storage.Volunteer{
		ID:        key.Encode(),
		Name:      ve.Name,
		Email:     ve.Email,
		Disabled:  ve.Disabled,
		CreatedAt: ve.CreatedAt,
	}
}
