package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/storage"
)

type ErrNoVolunteers struct{}

func (ErrNoVolunteers) Error() string {
	return "no volunteer submissions for this user are found"
}

func (d *Datastore) GetVolunteer(ctx context.Context, userId string) (*storage.Volunteer, error) {
	q := datastore.NewQuery(volunteerKind).Filter("UserId = ", userId).Filter("Disabled = ", false).Limit(1)
	volunteers := []volunteerEntity{}
	keys, err := d.client.GetAll(ctx, q, &volunteers)
	if err != nil {
		return nil, fmt.Errorf("error fetching volunteer submissions: %w", err)
	}
	if len(volunteers) == 0 {
		return nil, ErrNoVolunteers{}
	}
	return fromVolunteerEntity(keys[0], &volunteers[0]), nil
}
