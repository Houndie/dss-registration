package volunteer

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/pkg/errors"
)

func (s *Service) Insert(ctx context.Context, token string, submission *VolunteerSubmission) error {
	s.logger.Tracef("fetching user-info for token %s", token)
	userinfo, err := s.authorizer.Userinfo(ctx, token)
	if err != nil {
		msg := "could not authorize user"
		s.logger.WithError(err).Debug(msg)
		return errors.Wrap(err, msg)
	}
	s.logger.Tracef("found user %s", userinfo.UserId)

	err = s.store.InsertVolunteer(ctx, &StoreVolunteerSubmission{
		UserId: userinfo.UserId,
		Name:   submission.Name,
		Email:  submission.Email,
	})

	if err != nil {
		switch err.(type) {
		case storage.ErrVolunteerExists:
			s.logger.WithError(err).Debug("user attempted to submit existing volunteer form")
		default:
			s.logger.WithError(err).Error("error inserting volunteer submission into store")
		}
		return errors.Wrapf(err, "error inserting volunteer submission into store")
	}

	return nil
}
