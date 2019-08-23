package listbyuser

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Authorizer interface {
	Userinfo(ctx context.Context, accessToken string) (*authorizer.Userinfo, error)
}

type Store interface {
	GetRegistrationsByUser(ctx context.Context, userId string) ([]*StoreRegistration, error)
}

type Service struct {
	authorizer Authorizer
	logger     *logrus.Logger
	store      Store
	client     SquareClient
}

type SquareClient interface {
	ListLocations(ctx context.Context) ([]*square.Location, error)
	BatchRetrieveOrders(ctx context.Context, locationId string, orderIds []string) ([]*square.Order, error)
}

func NewService(authorizer Authorizer, logger *logrus.Logger, store Store, client SquareClient) *Service {
	return &Service{
		authorizer: authorizer,
		logger:     logger,
		store:      store,
		client:     client,
	}
}

func (s *Service) ListByUser(ctx context.Context, token string) ([]*Registration, error) {
	s.logger.Trace("In list by user service")
	s.logger.Tracef("fetching user-info for token %s", token)
	userinfo, err := s.authorizer.Userinfo(ctx, token)
	if err != nil {
		msg := "could not authorize user"
		s.logger.WithError(err).Debug(msg)
		return nil, errors.Wrap(err, msg)
	}
	s.logger.Tracef("found user %s", userinfo.UserId)

	s.logger.Tracef("fetching registrations for user %s", userinfo.UserId)
	r, err := s.store.GetRegistrationsByUser(ctx, userinfo.UserId)
	if err != nil {
		msg := "error fetching registrations from store"
		s.logger.WithError(err).Error(msg)
		return nil, errors.Wrap(err, msg)
	}
	s.logger.Tracef("found %d registrations", len(r))

	s.logger.Trace("fetching locations from square")
	locations, err := s.client.ListLocations(ctx)
	if err != nil {
		msg := "error listing locations from square"
		s.logger.WithError(err).Error(msg)
		return nil, errors.Wrap(err, msg)
	}
	s.logger.Tracef("found %d locations", len(locations))

	if len(locations) != 1 {
		msg := fmt.Errorf("found unexpected number of locations %d", len(locations))
		s.logger.Error(msg)
		return nil, msg
	}
	s.logger.Tracef("found location %s", locations[0].Id)

	orderIds := []string{}
	for _, reg := range r {
		orderIds = append(orderIds, reg.OrderIds...)
	}
	s.logger.Tracef("found %d total orders between all locations", len(orderIds))

	orderMap := map[string]*square.Order{}
	if len(orderIds) > 0 {
		s.logger.Trace("retrieving orders from square")
		orders, err := s.client.BatchRetrieveOrders(ctx, locations[0].Id, orderIds)
		if err != nil {
			msg := "error retrieving orders matching ids"
			s.logger.WithError(err).Error(msg)
			return nil, errors.Wrap(err, msg)
		}

		for _, order := range orders {
			orderMap[order.Id] = order
		}
	}

	s.logger.Trace("assembling registration response")
	registrations := make([]*Registration, len(r))
	for i, reg := range r {
		paid := true
		for _, id := range reg.OrderIds {
			if orderMap[id].State != square.OrderStateCompleted {
				paid = false
				break
			}
		}

		registrations[i] = &Registration{
			Id:        reg.Id,
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
