package discount

import (
	"context"
	"fmt"
)

func (s *Service) Delete(ctx context.Context, token, code string) error {
	s.logger.Trace("update discount service")
	s.logger.Tracef("fetching user-info for token %s", token)
	userinfo, err := s.authorizer.Userinfo(ctx, token)
	if err != nil {
		return fmt.Errorf("could not authorize user: %w", err)
	}
	s.logger.Tracef("found user %s", userinfo.UserID)
	isAdmin, err := s.store.IsAdmin(ctx, userinfo.UserID)
	if err != nil {
		return fmt.Errorf("could not fetch admin status from store: %w", err)
	}

	if !isAdmin {
		return ErrUnauthorized
	}

	err = s.store.DeleteDiscount(ctx, code)
	if err != nil {
		return fmt.Errorf("could not delete discount from store: %w", err)
	}

	return nil
}
