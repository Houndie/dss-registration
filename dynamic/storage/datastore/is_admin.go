package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/pkg/errors"
)

func (d *Datastore) IsAdmin(ctx context.Context, userId string) (bool, error) {
	q := datastore.NewQuery(adminKind).Filter("UserId =", userId).Limit(1)
	admins := []adminEntity{}
	_, err := d.client.GetAll(ctx, q, &admins)
	if err != nil {
		return false, errors.Wrap(err, "error getting admins")
	}
	return len(admins) > 0, nil
}
