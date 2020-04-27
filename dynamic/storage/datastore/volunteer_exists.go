package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
)

func (d *Datastore) VolunteerExists(ctx context.Context, userId string) (bool, error) {
	q := datastore.NewQuery(volunteerKind).Filter("UserId = ", userId).Filter("Disabled = ", false).Limit(1).KeysOnly()
	keys, err := d.client.GetAll(ctx, q, nil)
	if err != nil {
		return false, fmt.Errorf("error fetching volunteer submissions: %w", err)
	}
	return (len(keys) > 0), nil
}
