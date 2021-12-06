package vaccine

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/object"
	"github.com/Houndie/dss-registration/dynamic/storage"
)

func (s *Service) Upload(ctx context.Context, token string, filesize int64, id string) (string, error) {
	s.logger.Trace("vaccine upload service")
	userinfo, err := s.authorizer.GetUserinfo(ctx, token)
	if err != nil {
		return "", fmt.Errorf("could not authorize user: %w", err)
	}

	s.logger.Tracef("fetching registrations for user %s", userinfo.UserID())
	r, err := s.store.GetRegistration(ctx, id)
	if err != nil {
		return "", fmt.Errorf("error getting registration: %w", err)
	}
	s.logger.Trace("found registration")

	if r.UserID != userinfo.UserID() && !userinfo.IsAllowed(s.permissionConfig.Upload) {
		s.logger.WithError(err).Debug("user id does not match that of found registration")
		s.logger.WithError(err).Tracef("registration provided user id %s, user provided %s", r.UserID, userinfo.UserID())
		return "", storage.ErrNotFound{Key: id}
	}

	approved, err := s.store.GetVaccine(ctx, id)
	if err != nil {
		return "", fmt.Errorf("error getting status of vaccine approval from store: %w", err)
	}

	if approved {
		return "", ErrAlreadyApproved
	}

	if filesize > object.PutMaxSize {
		return "", ErrFileTooBig{Filesize: filesize}
	}

	url, err := s.objectClient.SignedPut(filesize, id)
	if err != nil {
		return "", fmt.Errorf("error creating signed put url: %w", err)
	}

	return url, nil
}
