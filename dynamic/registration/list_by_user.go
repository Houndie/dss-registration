package registration

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/common"
)

func (s *Service) ListByUser(ctx context.Context, token string) ([]*Info, error) {
	s.logger.Trace("In list by user service")
	s.logger.Tracef("fetching user-info for token %s", token)
	userinfo, err := s.authorizer.GetUserinfo(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("could not authorize user: %w", err)
	}
	s.logger.Tracef("found user %s", userinfo.UserID())

	s.logger.Tracef("fetching registrations for user %s", userinfo.UserID())
	r, err := s.store.GetRegistrationsByUser(ctx, userinfo.UserID())
	if err != nil {
		return nil, fmt.Errorf("error fetching registrations from store: %w", err)
	}
	s.logger.Tracef("found %d registrations", len(r))
	if len(r) == 0 {
		return nil, nil
	}

	orderIDs := map[string][]string{}
	for _, reg := range r {
		orderIDs[reg.ID] = reg.OrderIDs
	}
	s.logger.Tracef("found %d total orders", len(orderIDs))

	pd := map[string]*common.PaymentData{}
	for _, reg := range r {
		pd[reg.ID] = &common.PaymentData{}
	}
	if len(orderIDs) > 0 {
		s.logger.Trace("fetching locations from square")
		locationsListRes, err := s.client.Locations.List(ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("error listing locations from square: %w", err)
		}
		s.logger.Tracef("found %d locations", len(locationsListRes.Locations))

		if len(locationsListRes.Locations) != 1 {
			return nil, fmt.Errorf("found unexpected number of locations %d", len(locationsListRes.Locations))
		}

		pd, err = common.GetSquarePayments(ctx, s.client, s.squareData.PurchaseItems, locationsListRes.Locations[0].ID, orderIDs)
		if err != nil {
			return nil, err
		}

	}

	s.logger.Trace("assembling registration response")
	registrations := make([]*Info, len(r))
	for i, reg := range r {
		registrations[i] = fromStorageRegistration(reg, pd[reg.ID])
	}
	s.logger.Tracef("returning %d registrations", len(registrations))

	return registrations, nil
}
