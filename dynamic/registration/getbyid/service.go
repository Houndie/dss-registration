package getbyid

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/registration/common"
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
	ListCatalog(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator
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

	squareDiscounts := map[string]common.ItemDiscount{}
	for _, d := range r.Discounts {
		for _, sd := range d.Discounts {
			squareDiscounts[sd.Name] = nil
		}
	}
	objects := s.client.ListCatalog(ctx, []square.CatalogObjectType{square.CatalogObjectTypeDiscount})
	for objects.Next() {
		discountObject, ok := objects.Value().CatalogObjectType.(*square.CatalogDiscount)
		if !ok {
			s.logger.Error("found non discount object when discount was requested")
			continue
		}
		_, ok = squareDiscounts[discountObject.Name]
		if !ok {
			continue
		}

		var itemDiscount common.ItemDiscount
		switch t := discountObject.DiscountType.(type) {
		case *square.CatalogDiscountFixedAmount:
			itemDiscount = &common.DollarDiscount{
				Amount: t.AmountMoney.Amount,
			}
		case *square.CatalogDiscountVariableAmount:
			itemDiscount = &common.DollarDiscount{
				Amount: t.AmountMoney.Amount,
			}
		case *square.CatalogDiscountFixedPercentage:
			itemDiscount = &common.PercentDiscount{
				Amount: t.Percentage,
			}
		case *square.CatalogDiscountVariablePercentage:
			itemDiscount = &common.PercentDiscount{
				Amount: t.Percentage,
			}
		default:
			err := errors.New("unknown item discount type found from square")
			s.logger.Error(err)
			return nil, err
		}

		squareDiscounts[discountObject.Name] = itemDiscount
	}
	if objects.Error() != nil {
		wrap := "error fetching catalog objects from square"
		s.logger.WithError(err).Error(wrap)
		return nil, errors.Wrap(err, wrap)
	}
	discounts := []*Discount{}
	s.logger.Tracef("parsing %d discounts", len(r.Discounts))
	for _, discount := range r.Discounts {
		singleDiscounts := make([]*SingleDiscount, len(discount.Discounts))
		allDiscountsFound := true
		for i, sd := range discount.Discounts {
			itemDiscount, ok := squareDiscounts[sd.Name]
			if !ok {
				err := errors.New("impossible code path, somehow a square discount was not added or removed from square discount map")
				s.logger.Error(err)
				return nil, err
			}
			if itemDiscount == nil {
				s.logger.Errorf("Discount %s was applied but is no longer found in square store.  Omitting this code in result")
				allDiscountsFound = false
				break
			}
			singleDiscounts[i] = &SingleDiscount{
				ItemDiscount: itemDiscount,
				AppliedTo:    sd.AppliedTo,
			}
		}
		if !allDiscountsFound {
			continue
		}
		discounts = append(discounts, &Discount{
			Code:      discount.Code,
			Discounts: singleDiscounts,
		})
	}
	s.logger.Tracef("%d discounts parsed and applied to registration", len(discounts))

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
		Discounts:       discounts,
	}, nil
}
