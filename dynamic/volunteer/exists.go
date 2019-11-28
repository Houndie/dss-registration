package volunteer

import (
	"context"

	"github.com/pkg/errors"
)

func (s *Service) Exists(ctx context.Context, token string) (bool, error) {
	s.logger.Tracef("fetching user-info for token %s", token)
	userinfo, err := s.authorizer.Userinfo(ctx, token)
	if err != nil {
		msg := "could not authorize user"
		s.logger.WithError(err).Debug(msg)
		return false, errors.Wrap(err, msg)
	}
	s.logger.Tracef("found user %s", userinfo.UserId)

	exists, err := s.store.VolunteerExists(ctx, userinfo.UserId)
	if err != nil {
		s.logger.WithError(err).Error("error checking for volunteer existance")
		return false, errors.Wrap(err, "error checking for volunteer existance")
	}
	return exists, nil
}
