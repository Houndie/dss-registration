package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/volunteer"
	"github.com/pkg/errors"
)

type ErrNoVolunteers struct{}

func (ErrNoVolunteers) Error() string {
	return "no volunteer submissions for this user are found"
}

func (d *Datastore) GetVolunteer(ctx context.Context, userId string) (*volunteer.StoreVolunteerSubmission, error) {
	q := datastore.NewQuery(volunteerKind).Filter("UserId = ", userId).Filter("Disabled = ", false).Limit(1)
	volunteers := []volunteerEntity{}
	_, err := d.client.GetAll(ctx, q, &volunteers)
	if err != nil {
		return nil, errors.Wrap(err, "error fetching volunteer submissions")
	}
	if len(volunteers) == 0 {
		return nil, ErrNoVolunteers{}
	}
	return &volunteer.StoreVolunteerSubmission{
		UserId: volunteers[0].UserId,
		Name:   volunteers[0].Name,
		Email:  volunteers[0].Email,
	}, nil
}
