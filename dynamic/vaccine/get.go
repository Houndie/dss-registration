package vaccine

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/storage"
)

func (s *Service) Get(ctx context.Context, token, id string) (Info, error) {
	s.logger.Trace("vaccine upload service")
	userinfo, err := s.authorizer.GetUserinfo(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("could not authorize user: %w", err)
	}

	s.logger.Tracef("fetching registrations for user %s", userinfo.UserID())
	r, err := s.store.GetRegistration(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting registration: %w", err)
	}
	s.logger.Trace("found registration")

	if r.UserID != userinfo.UserID() && !userinfo.IsAllowed(s.permissionConfig.Get) {
		s.logger.WithError(err).Debug("user id does not match that of found registration")
		s.logger.WithError(err).Tracef("registration provided user id %s, user provided %s", r.UserID, userinfo.UserID())
		return nil, storage.ErrNotFound{Key: id}
	}

	approved, err := s.store.GetVaccine(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting status of vaccine approval from store: %w", err)
	}

	if approved {
		return &VaxApproved{}, nil
	}

	exists, err := s.objectClient.Exists(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error determining if object exists: %w", err)
	}

	if !exists {
		return &NoVaxProofSupplied{}, nil
	}

	url, err := s.objectClient.SignedGet(id)
	if err != nil {
		return nil, fmt.Errorf("error generating signed get URL: %w", err)
	}

	return &VaxApprovalPending{
		URL: url,
	}, nil
}
