package datastore

import (
	"context"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/volunteer"
	"github.com/pkg/errors"
)

func (d *Datastore) InsertVolunteer(ctx context.Context, submission *volunteer.StoreVolunteerSubmission) error {
	q := datastore.NewQuery(volunteerKind).Filter("UserId = ", submission.UserId).Filter("Disabled = ", false).Limit(1).KeysOnly()
	keys, err := d.client.GetAll(ctx, q, nil)
	if err != nil {
		return errors.Wrap(err, "error fetching existing volunteer submission")
	}
	if len(keys) != 0 {
		return storage.ErrVolunteerExists{UserId: submission.UserId}
	}
	key := datastore.IncompleteKey(volunteerKind, nil)

	_, err = d.client.Put(ctx, key, &volunteerEntity{
		UserId:   submission.UserId,
		Name:     submission.Name,
		Email:    submission.Email,
		Disabled: false,
		Created:  time.Now(),
	})
	if err != nil {
		return errors.Wrap(err, "error inserting volunteer submission to store")
	}
	return nil
}
