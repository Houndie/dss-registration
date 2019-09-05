package update

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/registration/common"
	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/utility"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Authorizer interface {
	Userinfo(ctx context.Context, accessToken string) (*authorizer.Userinfo, error)
}

type Store interface {
	GetUpdateRegistration(context.Context, string) (*StoreOldRegistration, error)
	UpdateRegistration(context.Context, *StoreUpdateRegistration, string) error
}
type SquareClient interface {
	ListCatalog(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator
	ListLocations(ctx context.Context) ([]*square.Location, error)
	CreateCheckout(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error)
	UpdateOrder(ctx context.Context, locationId, orderId string, order *square.Order, fieldsToClear []string, idempotencyKey string) (*square.Order, error)
	BatchRetrieveOrders(ctx context.Context, locationId string, orderIds []string) ([]*square.Order, error)
}

type Service struct {
	logger     *logrus.Logger
	store      Store
	authorizer Authorizer
	client     SquareClient
}

func NewService(logger *logrus.Logger, authorizer Authorizer, store Store, client SquareClient) *Service {
	return &Service{
		logger:     logger,
		store:      store,
		authorizer: authorizer,
		client:     client,
	}
}

func hasPurchase(newRegistration *Registration, oldRegistration *StoreOldRegistration) bool {
	switch newRegistration.PassType.(type) {
	case *common.WeekendPass:
		if _, ok := oldRegistration.PassType.(*common.WeekendPass); !ok {
			return true
		}
	case *common.DanceOnlyPass:
		if _, ok := oldRegistration.PassType.(*common.DanceOnlyPass); !ok {
			return true
		}
	}

	if newRegistration.MixAndMatch != nil && oldRegistration.MixAndMatch == nil {
		return true
	}

	if newRegistration.TShirt != nil && oldRegistration.TShirt == nil {
		return true
	}

	if newRegistration.TeamCompetition != nil && oldRegistration.TeamCompetition == nil {
		return true
	}

	if newRegistration.SoloJazz && !oldRegistration.SoloJazz {
		return true
	}
	return false
}

func (s *Service) Update(ctx context.Context, token string, registration *Registration, redirectUrl string) (string, error) {
	s.logger.Tracef("fetching old registration id %s", registration.Id)
	oldRegistration, err := s.store.GetUpdateRegistration(ctx, registration.Id)
	if err != nil {
		switch errors.Cause(err).(type) {
		case storage.ErrNotFound:
			newErr := ErrBadRegistrationId{registration.Id}
			s.logger.WithError(err).Debug(newErr)
			return "", newErr
		default:
			msg := "error fetching registrations from store"
			s.logger.WithError(err).Error(msg)
			return "", errors.Wrap(err, msg)
		}
	}

	s.logger.Tracef("fetching user-info for token %s", token)
	userinfo, err := s.authorizer.Userinfo(ctx, token)
	if err != nil {
		msg := "could not authorize user"
		s.logger.WithError(err).Debug(msg)
		return "", errors.Wrap(err, msg)
	}
	s.logger.Tracef("found user %s", userinfo.UserId)

	if oldRegistration.UserId != userinfo.UserId {
		s.logger.Debugf("registration found does not belong to user")
		return "", ErrBadRegistrationId{registration.Id}
	}

	switch oldRegistration.PassType.(type) {
	case *common.WeekendPass:
		if _, ok := registration.PassType.(*common.WeekendPass); !ok {
			err := ErrAlreadyPurchased{
				Field:         "Pass Type",
				ExistingValue: "Full Weekend",
			}
			s.logger.Debug(err)
			return "", err
		}
	case *common.DanceOnlyPass:
		if _, ok := registration.PassType.(*common.DanceOnlyPass); !ok {
			err := ErrAlreadyPurchased{
				Field:         "Pass Type",
				ExistingValue: "DanceOnly",
			}
			s.logger.Debug(err)
			return "", err
		}
	}

	if oldRegistration.MixAndMatch != nil && registration.MixAndMatch == nil {
		err := ErrAlreadyPurchased{
			Field:         "Mix and Match",
			ExistingValue: "Yes",
		}
		s.logger.Debug(err)
		return "", err
	}

	if oldRegistration.SoloJazz && !registration.SoloJazz {
		err := ErrAlreadyPurchased{
			Field:         "Solo Jazz",
			ExistingValue: "Yes",
		}
		s.logger.Debug(err)
		return "", err
	}

	if oldRegistration.TeamCompetition != nil && registration.TeamCompetition == nil {
		err := ErrAlreadyPurchased{
			Field:         "Team Competition",
			ExistingValue: "Yes",
		}
		s.logger.Debug(err)
		return "", err
	}

	if oldRegistration.TShirt != nil && registration.TShirt == nil {
		err := ErrAlreadyPurchased{
			Field:         "TShirt",
			ExistingValue: "Yes",
		}
		s.logger.Debug(err)
		return "", err
	}

	updateRegistration := &StoreUpdateRegistration{
		FirstName:       registration.FirstName,
		LastName:        registration.LastName,
		StreetAddress:   registration.StreetAddress,
		City:            registration.City,
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
	}

	/*if !hasUserPurchase && len(oldRegistration.OrderIds) == 0 {
		err = s.store.UpdateRegistration(ctx, updateRegistration, registration.Id)
		if err != nil {
			msg := "error updating registration in database"
			s.logger.WithError(err).Error(msg)
			return "", errors.Wrap(err, msg)
		}
		return redirectUrl, nil
	}*/
	s.logger.Trace("generating reference id")
	referenceId, err := uuid.NewV4()
	if err != nil {
		wrap := "error generating reference id"
		s.logger.WithError(err).Error(wrap)
		return "", errors.Wrap(err, wrap)
	}

	s.logger.Trace("registration contians items that must be paid, making square calls")

	s.logger.Trace("Fetching all locations from square")
	locations, err := s.client.ListLocations(ctx)
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

	idempotencyKey, err := uuid.NewV4()
	if err != nil {
		wrap := "error generating idempotency key"
		s.logger.WithError(err).Error(wrap)
		return "", errors.Wrap(err, wrap)
	}

	order := &square.CreateOrderRequest{
		IdempotencyKey: idempotencyKey.String(),
		Order: &square.Order{
			ReferenceId: referenceId.String(),
			LocationId:  locations[0].Id,
			LineItems:   []*square.OrderLineItem{},
			Version:     1,
		},
	}

	if hasPurchase(registration, oldRegistration) {

		s.logger.Trace("Fetching all items from square")
		objects := s.client.ListCatalog(ctx, nil)

		for objects.Next() {
			item, ok := objects.Value().CatalogObjectType.(*square.CatalogItem)
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
					if registration.MixAndMatch == nil || oldRegistration.MixAndMatch != nil {
						continue
					}
					order.Order.LineItems = append(order.Order.LineItems, &square.OrderLineItem{
						Quantity:        "1",
						CatalogObjectId: v.Id,
					})
				case utility.TeamCompItem:
					if registration.TeamCompetition == nil || oldRegistration.TeamCompetition != nil {
						continue
					}
					order.Order.LineItems = append(order.Order.LineItems, &square.OrderLineItem{
						Quantity:        "1",
						CatalogObjectId: v.Id,
					})
				case utility.SoloJazzItem:
					if !registration.SoloJazz || oldRegistration.SoloJazz {
						continue
					}
					order.Order.LineItems = append(order.Order.LineItems, &square.OrderLineItem{
						Quantity:        "1",
						CatalogObjectId: v.Id,
					})
				case utility.TShirtItem:
					if registration.TShirt == nil || oldRegistration.TShirt != nil {
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
				if _, ok := registration.PassType.(*common.DanceOnlyPass); !ok {
					continue
				}
				if _, ok := oldRegistration.PassType.(*common.DanceOnlyPass); ok {
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
				weekendPass, ok := registration.PassType.(*common.WeekendPass)
				if !ok {
					continue
				}
				if _, ok := oldRegistration.PassType.(*common.WeekendPass); ok {
					continue
				}

				var tierString string
				switch weekendPass.Tier {
				case common.WeekendPassTier1:
					tierString = utility.WeekendPassTier1Name
				case common.WeekendPassTier2:
					tierString = utility.WeekendPassTier2Name
				case common.WeekendPassTier3:
					tierString = utility.WeekendPassTier3Name
				case common.WeekendPassTier4:
					tierString = utility.WeekendPassTier4Name
				case common.WeekendPassTier5:
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
		if objects.Error() != nil {
			wrap := "error fetching all items from square"
			utility.LogSquareError(s.logger, objects.Error(), wrap)
			return "", errors.Wrap(objects.Error(), wrap)
		}
	}
	unpaidOrderIds := []string{}
	unpaidOrders := []*square.Order{}
	s.logger.Tracef("found %d existing orders", len(oldRegistration.OrderIds))
	if len(oldRegistration.OrderIds) > 0 {
		existingOrders, err := s.client.BatchRetrieveOrders(ctx, locations[0].Id, oldRegistration.OrderIds)
		if err != nil {
			wrap := "error fetching all items from square"
			utility.LogSquareError(s.logger, err, wrap)
			return "", errors.Wrap(err, wrap)
		}
		for _, existingOrder := range existingOrders {
			if existingOrder.State == square.OrderStateOpen {
				s.logger.Tracef("Order %s still unpaid", existingOrder.Id)
				for _, existingLineItem := range existingOrder.LineItems {
					order.Order.LineItems = append(order.Order.LineItems, &square.OrderLineItem{
						Quantity:        existingLineItem.Quantity,
						CatalogObjectId: existingLineItem.CatalogObjectId,
					})
				}
				unpaidOrderIds = append(unpaidOrderIds, existingOrder.Id)
				unpaidOrders = append(unpaidOrders, existingOrder)
			}
		}
	}

	if len(order.Order.LineItems) == 0 {
		err = s.store.UpdateRegistration(ctx, updateRegistration, registration.Id)
		if err != nil {
			msg := "error updating registration in database"
			s.logger.WithError(err).Error(msg)
			return "", errors.Wrap(err, msg)
		}
		return redirectUrl, nil
	}

	s.logger.Trace("creating checkout with square")
	checkout, err := s.client.CreateCheckout(ctx, locations[0].Id, idempotencyKey.String(), order, false, utility.SmackdownEmail, registration.Email, nil, redirectUrl, nil, "")
	if err != nil {
		wrap := "error creating square checkout"
		utility.LogSquareError(s.logger, err, wrap)
		return "", errors.Wrap(err, wrap)
	}

	updateRegistration.OrderUpdate = &StoreOrderUpdate{
		NewId:       checkout.Order.Id,
		ObsoleteIds: unpaidOrderIds,
	}

	s.logger.Trace("Adding registration to database")
	err = s.store.UpdateRegistration(ctx, updateRegistration, registration.Id)
	if err != nil {
		wrap := "error adding registration to database"
		s.logger.WithError(err).Error(wrap)
		return "", errors.Wrap(err, wrap)
	}
	for _, unpaidOrder := range unpaidOrders {

		sparseOrder := &square.Order{
			Id:      unpaidOrder.Id,
			Version: unpaidOrder.Version,
			State:   square.OrderStateCanceled,
		}

		_, err := s.client.UpdateOrder(ctx, locations[0].Id, unpaidOrder.Id, sparseOrder, nil, idempotencyKey.String())
		if err != nil {
			wrap := "error cleaning up order from square"
			s.logger.WithError(err).Error(wrap)
			return "", errors.Wrap(err, wrap)
		}
	}

	return checkout.CheckoutPageUrl, nil
}
