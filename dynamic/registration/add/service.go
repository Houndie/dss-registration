package add

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/utility"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Store interface {
	AddRegistration(context.Context, *StoreRegistration) error
}

type SquareClient interface {
	ListCatalog([]string) ([]*square.CatalogObject, error)
	ListLocations() ([]*square.Location, error)
	CreateCheckout(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error)
}

type Service struct {
	client SquareClient
	store  Store
	logger *logrus.Logger
}

func NewService(logger *logrus.Logger, store Store, client SquareClient) *Service {
	return &Service{
		store:  store,
		logger: logger,
		client: client,
	}
}

func (s *Service) Add(ctx context.Context, registration *Registration, redirectUrl string) (string, error) {
	s.logger.Trace("Fetching all locations from square")
	locations, err := s.client.ListLocations()
	if err != nil {
		wrap := "error listing locations from square"
		utility.LogSquareError(s.logger, err, wrap)
		return "", errors.Wrap(err, wrap)
	}
	if len(locations) != 1 {
		err := fmt.Errorf("found wrong number of locations %v", len(locations))
		s.logger.Error(err)
		return "", err
	}

	s.logger.Trace("Fetching all items from square")
	objects, err := s.client.ListCatalog(nil)
	if err != nil {
		wrap := "error fetching all items from square"
		utility.LogSquareError(s.logger, err, wrap)
		return "", errors.Wrap(err, wrap)
	}

	idempotencyKey, err := uuid.NewV4()
	if err != nil {
		wrap := "error generating idempotency key"
		s.logger.WithError(err).Error(wrap)
		return "", errors.Wrap(err, wrap)
	}

	referenceId, err := uuid.NewV4()
	if err != nil {
		wrap := "error generating reference id"
		s.logger.WithError(err).Error(wrap)
		return "", errors.Wrap(err, wrap)
	}

	s.logger.Trace("Adding registration to database")
	err = s.store.AddRegistration(ctx, &StoreRegistration{
		FirstName:       registration.FirstName,
		LastName:        registration.LastName,
		StreetAddress:   registration.StreetAddress,
		City:            registration.City,
		State:           registration.State,
		ZipCode:         registration.ZipCode,
		Email:           registration.Email,
		HomeScene:       registration.HomeScene,
		IsStudent:       registration.IsStudent,
		PassType:        registration.PassType,
		MixAndMatch:     registration.MixAndMatch,
		SoloJazz:        registration.SoloJazz,
		TeamCompetition: registration.TeamCompetition,
		TShirt:          registration.TShirt,
		Housing:         registration.Housing,
		TransactionID:   idempotencyKey,
	})
	if err != nil {
		wrap := "error adding registration to database"
		s.logger.WithError(err).Error(wrap)
		return "", errors.Wrap(err, wrap)
	}

	order := &square.CreateOrderRequest{
		IdempotencyKey: idempotencyKey.String(),
		Order: &square.Order{
			ReferenceID: referenceId.String(),
			LocationID:  locations[0].Id,
			LineItems:   []*square.OrderLineItem{},
		},
	}

	for _, object := range objects {
		item, ok := object.CatalogObjectType.(*square.CatalogItem)
		if !ok {
			s.logger.Trace("Square object was not of type catalog item")
			continue
		}
		s.logger.Tracef("Comparing item name %s to legend", item.Name)
		switch item.Name {
		case utility.MixAndMatchItem, utility.TeamCompItem, utility.SoloJazzItem, utility.TShirtItem:
			if len(item.Variations) != 1 {
				err := fmt.Errorf("Found unexpected number of variations: %v", len(item.Variations))
				s.logger.Error(err)
				return "", err
			}
			v := item.Variations[0]
			_, ok := v.CatalogObjectType.(*square.CatalogItemVariation)
			if !ok {
				err := "Invalid response from square...item variation isn't a variation?"
				s.logger.Error(err)
				return "", errors.New(err)
			}
			switch item.Name {
			case utility.MixAndMatchItem:
				if registration.MixAndMatch == nil {
					continue
				}
				order.Order.LineItems = append(order.Order.LineItems, &square.OrderLineItem{
					Quantity:        "1",
					CatalogObjectId: v.Id,
				})
			case utility.TeamCompItem:
				if registration.TeamCompetition == nil {
					continue
				}
				order.Order.LineItems = append(order.Order.LineItems, &square.OrderLineItem{
					Quantity:        "1",
					CatalogObjectId: v.Id,
				})
			case utility.SoloJazzItem:
				if !registration.SoloJazz {
					continue
				}
				order.Order.LineItems = append(order.Order.LineItems, &square.OrderLineItem{
					Quantity:        "1",
					CatalogObjectId: v.Id,
				})
			case utility.TShirtItem:
				if registration.TShirt == nil {
					continue
				}
				order.Order.LineItems = append(order.Order.LineItems, &square.OrderLineItem{
					Quantity:        "1",
					CatalogObjectId: v.Id,
				})
			default:
				err := errors.New("Impossible code path...how did I get here")
				s.logger.Error(err)
				return "", err
			}
		case utility.DancePassItem:
			s.logger.Trace("Found dance pass item")
			if _, ok := registration.PassType.(*DanceOnlyPass); !ok {
				continue
			}
			for _, v := range item.Variations {
				variation, ok := v.CatalogObjectType.(*square.CatalogItemVariation)
				if !ok {
					err := "Invalid response from square...item variation isn't a variation?"
					s.logger.Error(err)
					return "", errors.New(err)
				}
				if variation.Name == "Presale" {
					s.logger.Trace("Found dance pass variant Presale")
					order.Order.LineItems = append(order.Order.LineItems, &square.OrderLineItem{
						Quantity:        "1",
						CatalogObjectId: v.Id,
					})
					break
				}
				s.logger.Tracef("Did not find dance pass variant Presale (found %s), moving on", variation.Name)
			}
		case utility.WeekendPassItem:
			s.logger.Trace("Found full weekend pass item")
			weekendPass, ok := registration.PassType.(*WeekendPass)
			if !ok {
				continue
			}

			var tierString string
			switch weekendPass.Tier {
			case Tier1:
				tierString = utility.WeekendPassTier1Name
			case Tier2:
				tierString = utility.WeekendPassTier2Name
			case Tier3:
				tierString = utility.WeekendPassTier3Name
			case Tier4:
				tierString = utility.WeekendPassTier4Name
			case Tier5:
				tierString = utility.WeekendPassTier5Name
			}
			for _, v := range item.Variations {
				variation, ok := v.CatalogObjectType.(*square.CatalogItemVariation)
				if !ok {
					err := "Invalid response from square...item variation isn't a variation?"
					s.logger.Error(err)
					return "", errors.New(err)
				}
				if variation.Name == tierString {
					s.logger.Trace("Found weekend pass")
					order.Order.LineItems = append(order.Order.LineItems, &square.OrderLineItem{
						Quantity:        "1",
						CatalogObjectId: v.Id,
					})
					break
				}
			}
		}
	}

	s.logger.Trace("creating checkout with square")
	checkout, err := s.client.CreateCheckout(ctx, locations[0].Id, idempotencyKey.String(), order, false, "info@daytonswingsmackdown.com", registration.Email, nil, redirectUrl, nil, "")
	if err != nil {
		wrap := "error creating square checkout"
		utility.LogSquareError(s.logger, err, wrap)
		return "", errors.Wrap(err, wrap)
	}
	return checkout.CheckoutPageUrl, nil
}
