package registration

import (
	"context"
	"fmt"

	"github.com/Houndie/square-go/objects"
	"github.com/Houndie/square-go/orders"
)

func (s *Service) SummaryByUser(ctx context.Context, token string) ([]*Summary, error) {
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

	s.logger.Trace("fetching locationsListRes.Locations from square")
	locationsListRes, err := s.client.Locations.List(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error listing locationsListRes.Locations from square: %w", err)
	}
	s.logger.Tracef("found %d locationsListRes.Locations", len(locationsListRes.Locations))

	if len(locationsListRes.Locations) != 1 {
		return nil, fmt.Errorf("found unexpected number of locationsListRes.Locations %d", len(locationsListRes.Locations))
	}
	s.logger.Tracef("found location %s", locationsListRes.Locations[0].ID)

	orderIDs := []string{}
	for _, reg := range r {
		orderIDs = append(orderIDs, reg.OrderIDs...)
	}
	s.logger.Tracef("found %d total orders between all locationsListRes.Locations", len(orderIDs))

	orderMap := map[string]*objects.Order{}
	if len(orderIDs) > 0 {
		s.logger.Trace("retrieving orders from square")
		ordersRes, err := s.client.Orders.BatchRetrieve(ctx, &orders.BatchRetrieveRequest{
			LocationID: locationsListRes.Locations[0].ID,
			OrderIDs:   orderIDs,
		})
		if err != nil {
			return nil, fmt.Errorf("error retrieving orders matching ids: %w", err)
		}

		for _, order := range ordersRes.Orders {
			orderMap[order.ID] = order
		}
	}

	s.logger.Trace("assembling registration response")
	registrations := make([]*Summary, len(r))
	for i, reg := range r {
		paid := true
		for _, id := range reg.OrderIDs {
			if orderMap[id].State != objects.OrderStateCompleted {
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
