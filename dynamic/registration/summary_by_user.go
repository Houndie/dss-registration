package registration

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/square"
)

func (s *Service) SummaryByUser(ctx context.Context, token string) ([]*Summary, error) {
	s.logger.Trace("In list by user service")
	s.logger.Tracef("fetching user-info for token %s", token)
	userinfo, err := s.authorizer.Userinfo(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("could not authorize user: %w", err)
	}
	s.logger.Tracef("found user %s", userinfo.UserID)

	s.logger.Tracef("fetching registrations for user %s", userinfo.UserID)
	r, err := s.store.GetRegistrationsByUser(ctx, userinfo.UserID)
	if err != nil {
		return nil, fmt.Errorf("error fetching registrations from store: %w", err)
	}
	s.logger.Tracef("found %d registrations", len(r))
	if len(r) == 0 {
		return nil, nil
	}

	s.logger.Trace("fetching locations from square")
	locations, err := s.client.ListLocations(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing locations from square: %w", err)
	}
	s.logger.Tracef("found %d locations", len(locations))

	if len(locations) != 1 {
		return nil, fmt.Errorf("found unexpected number of locations %d", len(locations))
	}
	s.logger.Tracef("found location %s", locations[0].ID)

	orderIDs := []string{}
	for _, reg := range r {
		orderIDs = append(orderIDs, reg.OrderIDs...)
	}
	s.logger.Tracef("found %d total orders between all locations", len(orderIDs))

	orderMap := map[string]*square.Order{}
	if len(orderIDs) > 0 {
		s.logger.Trace("retrieving orders from square")
		orders, err := s.client.BatchRetrieveOrders(ctx, locations[0].ID, orderIDs)
		if err != nil {
			return nil, fmt.Errorf("error retrieving orders matching ids: %w", err)
		}

		for _, order := range orders {
			orderMap[order.ID] = order
		}
	}

	s.logger.Trace("assembling registration response")
	registrations := make([]*Summary, len(r))
	for i, reg := range r {
		paid := true
		for _, id := range reg.OrderIDs {
			if orderMap[id].State != square.OrderStateCompleted {
				paid = false
				break
			}
		}

		registrations[i] = &Summary{
			ID:        reg.ID,
			FirstName: reg.FirstName,
			LastName:  reg.LastName,
			Email:     reg.Email,
			CreatedAt: reg.CreatedAt,
			Paid:      paid,
		}
	}
	s.logger.Tracef("returning %d registrations", len(registrations))

	return registrations, nil
}
