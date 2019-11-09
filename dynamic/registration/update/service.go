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
	GetDiscounts(ctx context.Context, codes []string) ([]string, []*common.StoreDiscount, error)
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

	discountKeys, discounts, err := s.store.GetDiscounts(ctx, registration.DiscountCodes)
	if err != nil {
		wrap := "error fetching discount codes from datastore"
		if _, ok := errors.Cause(err).(common.ErrDiscountDoesNotExist); ok {
			s.logger.WithError(err).Debug(wrap)
		} else {
			s.logger.WithError(err).Error(wrap)
		}
		return "", errors.Wrap(err, wrap)
	}

	for _, code := range registration.DiscountCodes {
		for _, d := range oldRegistration.Discounts {
			if d.Code == code {
				err := ErrDiscountAlreadyApplied{
					Code: code,
				}
				s.logger.Debug(err)
				return "", err
			}
		}
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
		NewDiscounts:    discountKeys,
	}

	s.logger.Trace("generating reference id")
	referenceId, err := uuid.NewV4()
	if err != nil {
		wrap := "error generating reference id"
		s.logger.WithError(err).Error(wrap)
		return "", errors.Wrap(err, wrap)
	}

	retUrl := redirectUrl
	if hasPurchase(registration, oldRegistration) || len(oldRegistration.OrderIds) > 0 {
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
		discountsMap := map[string]common.PurchaseItem{}
		for _, discount := range discounts {
			discountsMap[discount.Name] = discount.AppliedTo
		}
		for _, discount := range oldRegistration.Discounts {
			for _, singleDiscount := range discount.Discounts {
				discountsMap[singleDiscount.Name] = singleDiscount.AppliedTo
			}
		}

		discountIds := map[common.PurchaseItem][]string{}

		purchaseItems := map[common.PurchaseItem]*square.OrderLineItem{}
		var tierString string
		switch t := registration.PassType.(type) {
		case *common.WeekendPass:
			purchaseItems[common.FullWeekendPurchaseItem] = &square.OrderLineItem{Quantity: "1"}
			switch t.Tier {
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
		case *common.DanceOnlyPass:
			purchaseItems[common.DanceOnlyPurchaseItem] = &square.OrderLineItem{Quantity: "1"}
		default:
			//Do nothing
		}

		if registration.MixAndMatch != nil {
			purchaseItems[common.MixAndMatchPurchaseItem] = &square.OrderLineItem{Quantity: "1"}
		}

		if registration.SoloJazz {
			purchaseItems[common.SoloJazzPurchaseItem] = &square.OrderLineItem{Quantity: "1"}
		}

		if registration.TeamCompetition != nil {
			purchaseItems[common.TeamCompetitionPurchaseItem] = &square.OrderLineItem{Quantity: "1"}
		}

		if registration.TShirt != nil {
			purchaseItems[common.TShirtPurchaseItem] = &square.OrderLineItem{Quantity: "1"}
		}

		unpaidItems := map[string]string{}
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
					unpaidOrderIds = append(unpaidOrderIds, existingOrder.Id)
					unpaidOrders = append(unpaidOrders, existingOrder)
					for _, lineItem := range existingOrder.LineItems {
						unpaidItems[lineItem.Uid] = lineItem.Quantity
					}
				}
			}
		}

		s.logger.Trace("Fetching all items from square")
		objects := s.client.ListCatalog(ctx, []square.CatalogObjectType{square.CatalogObjectTypeItem, square.CatalogObjectTypeDiscount})
		for objects.Next() {
			switch o := objects.Value().CatalogObjectType.(type) {
			case *square.CatalogItem:
				s.logger.Tracef("Comparing item name %s to legend", o.Name)
				switch o.Name {
				case utility.MixAndMatchItem, utility.TeamCompItem, utility.SoloJazzItem, utility.TShirtItem:
					if len(o.Variations) != 1 {
						err := fmt.Errorf("Found unexpected number of variations: %v", len(o.Variations))
						s.logger.Error(err)
						return "", err
					}
					v := o.Variations[0]
					quantity, ok := unpaidItems[v.Id]
					if ok {
						var purchaseItem common.PurchaseItem
						switch o.Name {
						case utility.MixAndMatchItem:
							purchaseItem = common.MixAndMatchPurchaseItem
						case utility.TeamCompItem:
							purchaseItem = common.TeamCompetitionPurchaseItem
						case utility.SoloJazzItem:
							purchaseItem = common.SoloJazzPurchaseItem
						case utility.TShirtItem:
							purchaseItem = common.TShirtPurchaseItem
						default:
							err := errors.New("Impossible code path...how did I get here")
							s.logger.Error(err)
							return "", err
						}
						purchaseItems[purchaseItem] = &square.OrderLineItem{CatalogObjectId: v.Id, Quantity: quantity}
						continue
					}
					_, ok = v.CatalogObjectType.(*square.CatalogItemVariation)
					if !ok {
						err := "Invalid response from square...item variation isn't a variation?"
						s.logger.Error(err)
						return "", errors.New(err)
					}
					var pi *square.OrderLineItem
					switch o.Name {
					case utility.MixAndMatchItem:
						pi, ok = purchaseItems[common.MixAndMatchPurchaseItem]
						if !ok {
							continue
						}
					case utility.TeamCompItem:
						pi, ok = purchaseItems[common.TeamCompetitionPurchaseItem]
						if !ok {
							continue
						}
					case utility.SoloJazzItem:
						pi, ok = purchaseItems[common.SoloJazzPurchaseItem]
						if !ok {
							continue
						}
					case utility.TShirtItem:
						pi, ok = purchaseItems[common.TShirtPurchaseItem]
						if !ok {
							continue
						}
					default:
						err := errors.New("Impossible code path...how did I get here")
						s.logger.Error(err)
						return "", err
					}
					pi.CatalogObjectId = v.Id
					continue
				case utility.DancePassItem:
					s.logger.Trace("Found dance pass item")
					for _, v := range o.Variations {
						quantity, ok := unpaidItems[v.Id]
						if ok {
							purchaseItems[common.DanceOnlyPurchaseItem] = &square.OrderLineItem{CatalogObjectId: v.Id, Quantity: quantity}
							break
						}

						pi, ok := purchaseItems[common.DanceOnlyPurchaseItem]
						if !ok {
							continue
						}
						variation, ok := v.CatalogObjectType.(*square.CatalogItemVariation)
						if !ok {
							err := "Invalid response from square...item variation isn't a variation?"
							s.logger.Error(err)
							return "", errors.New(err)
						}
						if variation.Name != "Presale" {
							s.logger.Tracef("Did not find dance pass variant Presale (found %s), moving on", variation.Name)
							continue
						}
						s.logger.Trace("Found dance pass variant Presale")
						pi.CatalogObjectId = v.Id
						break
					}
				case utility.WeekendPassItem:
					s.logger.Trace("Found full weekend pass item")

					for _, v := range o.Variations {
						quantity, ok := unpaidItems[v.Id]
						if ok {
							purchaseItems[common.FullWeekendPurchaseItem] = &square.OrderLineItem{CatalogObjectId: v.Id, Quantity: quantity}
							break
						}
						pi, ok := purchaseItems[common.FullWeekendPurchaseItem]
						if !ok {
							continue
						}

						variation, ok := v.CatalogObjectType.(*square.CatalogItemVariation)
						if !ok {
							err := "Invalid response from square...item variation isn't a variation?"
							s.logger.Error(err)
							return "", errors.New(err)
						}
						if variation.Name != tierString {
							continue
						}
						s.logger.Trace("Found weekend pass")
						pi.CatalogObjectId = v.Id
						break
					}
				}
			case *square.CatalogDiscount:
				appliedTo, ok := discountsMap[o.Name]
				if !ok {
					continue
				}
				delete(discountsMap, o.Name)
				ids, ok := discountIds[appliedTo]
				if !ok || ids == nil {
					discountIds[appliedTo] = []string{objects.Value().Id}
					continue
				}
				ids = append(ids, objects.Value().Id)
			default:
				err := errors.New("found unknown catalog object type (should be impossible)")
				s.logger.Error(err)
				return "", err
			}
		}
		if objects.Error() != nil {
			wrap := "error fetching all items from square"
			utility.LogSquareError(s.logger, objects.Error(), wrap)
			return "", errors.Wrap(objects.Error(), wrap)
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
		for appliedTo, orderItem := range purchaseItems {
			discountIds, ok := discountIds[appliedTo]
			if ok {
				orderItem.Discounts = make([]*square.OrderLineItemDiscount, len(discountIds))
				for i, discountId := range discountIds {
					orderItem.Discounts[i] = &square.OrderLineItemDiscount{
						CatalogObjectId: discountId,
					}
				}
			}
			order.Order.LineItems = append(order.Order.LineItems, orderItem)
		}

		updateRegistration.ObsoleteOrderIds = unpaidOrderIds

		s.logger.Trace("creating checkout with square")
		checkout, err := s.client.CreateCheckout(ctx, locations[0].Id, idempotencyKey.String(), order, false, utility.SmackdownEmail, registration.Email, nil, redirectUrl, nil, "")
		if err != nil {
			errorList, ok := err.(*square.ErrorList)

			// If this error is anything other than "can't create checkouts worth less than a dollar"
			if !ok || len(errorList.Errors) > 1 || errorList.Errors[0].Category != square.ErrorCategoryInvalidRequestError || errorList.Errors[0].Code != square.ErrorCodeValueTooLow || errorList.Errors[0].Field != "order.total_money.amount" {
				wrap := "error creating square checkout"
				utility.LogSquareError(s.logger, err, wrap)
				return "", errors.Wrap(err, wrap)
			}
			s.logger.Trace("registration with checkout amount less than a dollar found, not creating order or checkout")
		} else {
			updateRegistration.NewOrderId = checkout.Order.Id
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
		retUrl = checkout.CheckoutPageUrl
	}

	s.logger.Trace("Adding registration to database")
	err = s.store.UpdateRegistration(ctx, updateRegistration, registration.Id)
	if err != nil {
		wrap := "error adding registration to database"
		s.logger.WithError(err).Error(wrap)
		return "", errors.Wrap(err, wrap)
	}

	return retUrl, nil
}
