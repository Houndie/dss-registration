package getbyid

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Authorizer interface {
	Userinfo(ctx context.Context, accessToken string) (*authorizer.Userinfo, error)
}

type Store interface {
	GetRegistrationById(ctx context.Context, id string) (*StoreRegistration, error)
}

type SquareClient interface {
	ListLocations(ctx context.Context) ([]*square.Location, error)
	BatchRetrieveOrders(ctx context.Context, locationId string, orderIds []string) ([]*square.Order, error)
}

type Service struct {
	logger     *logrus.Logger
	authorizer Authorizer
	store      Store
	client     SquareClient
}

func NewService(logger *logrus.Logger, authorizer Authorizer, store Store, client SquareClient) *Service {
	return &Service{
		logger:     logger,
		authorizer: authorizer,
		store:      store,
		client:     client,
	}
}

func (s *Service) GetById(ctx context.Context, token, registrationId string) (*Registration, error) {
	s.logger.Trace("In get by id service")
	s.logger.Tracef("fetching user-info for token %s", token)
	userinfo, err := s.authorizer.Userinfo(ctx, token)
	if err != nil {
		msg := "could not authorize user"
		s.logger.WithError(err).Debug(msg)
		return nil, errors.Wrap(err, msg)
	}
	s.logger.Tracef("found user %s", userinfo.UserId)

	s.logger.Tracef("fetching registrations for user %s", userinfo.UserId)
	r, err := s.store.GetRegistrationById(ctx, registrationId)
	if err != nil {
		switch errors.Cause(err).(type) {
		case storage.ErrNotFound:
			newErr := ErrBadRegistrationId{registrationId}
			s.logger.WithError(err).Debug(newErr)
			return nil, newErr
		default:
			msg := "error fetching registrations from store"
			s.logger.WithError(err).Error(msg)
			return nil, errors.Wrap(err, msg)
		}
	}
	s.logger.Trace("found registration")

	if r.UserId != userinfo.UserId {
		err := ErrBadRegistrationId{registrationId}
		s.logger.WithError(err).Debug("user id does not match that of found registration")
		s.logger.WithError(err).Tracef("registration provided user id %s, user provided %s", r.UserId, userinfo.UserId)
		return nil, err
	}

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

	var unpaidItems *UnpaidItems
	if len(r.OrderIds) > 0 {
		s.logger.Trace("retrieving orders from square")
		squareOrders, err := s.client.BatchRetrieveOrders(ctx, locations[0].Id, r.OrderIds)
		if err != nil {
			msg := "error retrieving orders matching ids"
			s.logger.WithError(err).Error(msg)
			return nil, errors.Wrap(err, msg)
		}
		unpaidItems = &UnpaidItems{
			OrderIds: []string{},
			Items:    []string{},
			Cost:     0,
		}
		for _, squareOrder := range squareOrders {
			if squareOrder.State != square.OrderStateOpen {
				continue
			}
			for _, squareOrderItem := range squareOrder.LineItems {
				unpaidItems.Items = append(unpaidItems.Items, squareOrderItem.Name)
				unpaidItems.Cost += squareOrderItem.TotalMoney.Amount
			}
			unpaidItems.OrderIds = append(unpaidItems.OrderIds, squareOrder.Id)
		}
		if len(unpaidItems.OrderIds) == 0 {
			unpaidItems = nil
		}
	}

	return &Registration{
		FirstName:       r.FirstName,
		LastName:        r.LastName,
		StreetAddress:   r.StreetAddress,
		City:            r.City,
		State:           r.State,
		ZipCode:         r.ZipCode,
		Email:           r.Email,
		HomeScene:       r.HomeScene,
		IsStudent:       r.IsStudent,
		PassType:        r.PassType,
		MixAndMatch:     r.MixAndMatch,
		SoloJazz:        r.SoloJazz,
		TeamCompetition: r.TeamCompetition,
		TShirt:          r.TShirt,
		Housing:         r.Housing,
		UnpaidItems:     unpaidItems,
	}, nil
}
