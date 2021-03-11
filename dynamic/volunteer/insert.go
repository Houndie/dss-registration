package volunteer

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/storage"
)

func (s *Service) Insert(ctx context.Context, token string, submission *VolunteerSubmission) error {
	s.logger.Tracef("fetching user-info for token %s", token)
	userinfo, err := s.authorizer.Userinfo(ctx, token)
	if err != nil {
		msg := "could not authorize user"
		s.logger.WithError(err).Debug(msg)
		return fmt.Errorf("%s: %w", msg, err)
	}
	s.logger.Tracef("found user %s", userinfo.UserID)

	err = s.store.InsertVolunteer(ctx, &storage.Volunteer{
		UserID: userinfo.UserID,
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
		return fmt.Errorf("error inserting volunteer submission into store: %w", err)
	}

	return nil
}
