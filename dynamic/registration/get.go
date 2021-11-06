package registration

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/storage"
)

func (s *Service) Get(ctx context.Context, token, registrationID string) (*Info, error) {
	s.logger.Trace("In get by id service")
	userinfo, err := s.authorizer.GetUserinfo(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("could not authorize user: %w", err)
	}

	s.logger.Tracef("fetching registrations for user %s", userinfo.UserID())
	r, err := s.store.GetRegistration(ctx, registrationID)
	if err != nil {
		return nil, fmt.Errorf("error getting registration: %w", err)
	}
	s.logger.Trace("found registration")

	if !userinfo.IsAllowed(s.permissionConfig.List) && r.UserID != userinfo.UserID() {
		s.logger.WithError(err).Debug("user id does not match that of found registration")
		s.logger.WithError(err).Tracef("registration provided user id %s, user provided %s", r.UserID, userinfo.UserID())
		return nil, storage.ErrNotFound{Key: registrationID}
	}

	pd := map[string]*common.PaymentData{r.ID: &common.PaymentData{}}
	if len(r.OrderIDs) > 0 {
		s.logger.Trace("fetching locations from square")
		locationsListRes, err := s.client.Locations.List(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("error listing locations from square: %w", err)
		}
		s.logger.Tracef("found %d locations", len(locationsListRes.Locations))

		if len(locationsListRes.Locations) != 1 {
			return nil, fmt.Errorf("found unexpected number of locations %d", len(locationsListRes.Locations))
		}
		s.logger.Tracef("found location %s", locationsListRes.Locations[0].ID)

		pd, err = common.GetSquarePayments(ctx, s.client, s.squareData.PurchaseItems, locationsListRes.Locations[0].ID, map[string][]string{r.ID: r.OrderIDs})
		if err != nil {
			return nil, err
		}
	}

	return fromStorageRegistration(r, pd[r.ID]), nil
}
