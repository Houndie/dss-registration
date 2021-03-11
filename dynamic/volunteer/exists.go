package volunteer

import (
	"context"
	"fmt"
)

func (s *Service) Exists(ctx context.Context, token string) (bool, error) {
	s.logger.Tracef("fetching user-info for token %s", token)
	userinfo, err := s.authorizer.Userinfo(ctx, token)
	if err != nil {
		msg := "could not authorize user"
		s.logger.WithError(err).Debug(msg)
		return false, fmt.Errorf("%s: %w", msg, err)
	}
	s.logger.Tracef("found user %s", userinfo.UserID())

	exists, err := s.store.VolunteerExists(ctx, userinfo.UserID())
	if err != nil {
		s.logger.WithError(err).Error("error checking for volunteer existance")
		return false, fmt.Errorf("error checking for volunteer existance: %w", err)
	}
	return exists, nil
}
