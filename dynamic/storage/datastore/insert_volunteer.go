package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/storage"
)

func (d *Datastore) InsertVolunteer(ctx context.Context, submission *storage.Volunteer) error {
	q := datastore.NewQuery(volunteerKind).Filter("UserId = ", submission.UserID).Filter("Disabled = ", false).Limit(1).KeysOnly()
	keys, err := d.client.GetAll(ctx, q, nil)
	if err != nil {
		return fmt.Errorf("error fetching existing volunteer submission: %w", err)
	}
	if len(keys) != 0 {
		return storage.ErrVolunteerExists{UserId: submission.UserID}
	}
	key := datastore.IncompleteKey(volunteerKind, nil)
	_, ve, err := toVolunteerEntity(submission)
	if err != nil {
		return fmt.Errorf("error converting volunteer to datastore type: %w", err)
	}

	_, err = d.client.Put(ctx, key, ve)
	if err != nil {
		return fmt.Errorf("error inserting volunteer submission to store: %w", err)
	}
	return nil
}
