package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/pkg/errors"
)

func (d *Datastore) VolunteerExists(ctx context.Context, userId string) (bool, error) {
	q := datastore.NewQuery(volunteerKind).Filter("UserId = ", userId).Filter("Disabled = ", false).Limit(1).KeysOnly()
	keys, err := d.client.GetAll(ctx, q, nil)
	if err != nil {
		return false, errors.Wrap(err, "error fetching volunteer submissions")
	}
	return (len(keys) > 0), nil
}
