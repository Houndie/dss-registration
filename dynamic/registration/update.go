package registration

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

func checkOldPurchases(newRegistration *Info, oldRegistration *storage.Registration) error {
	switch oldRegistration.PassType.(type) {
	case *storage.WeekendPass:
		if _, ok := newRegistration.PassType.(*WeekendPass); !ok {
			return ErrAlreadyPurchased{
				Field:         "Pass Type",
				ExistingValue: "Full Weekend",
			}
		}
	case *storage.DanceOnlyPass:
		if _, ok := newRegistration.PassType.(*DanceOnlyPass); !ok {
			return ErrAlreadyPurchased{
				Field:         "Pass Type",
				ExistingValue: "DanceOnly",
			}
		}
	}

	if oldRegistration.MixAndMatch != nil && newRegistration.MixAndMatch == nil {
		return ErrAlreadyPurchased{
			Field:         "Mix and Match",
			ExistingValue: "Yes",
		}
	}

	if oldRegistration.SoloJazz && newRegistration.SoloJazz == nil {
		return ErrAlreadyPurchased{
			Field:         "Solo Jazz",
			ExistingValue: "Yes",
		}
	}

	if oldRegistration.TeamCompetition != nil && newRegistration.TeamCompetition == nil {
		return ErrAlreadyPurchased{
			Field:         "Team Competition",
			ExistingValue: "Yes",
		}
	}

	if oldRegistration.TShirt != nil && newRegistration.TShirt == nil {
		return ErrAlreadyPurchased{
			Field:         "TShirt",
			ExistingValue: "Yes",
		}
	}

	for _, od := range oldRegistration.DiscountCodes {
		found := false
		for _, d := range newRegistration.DiscountCodes {
			if d == od {
				found = true
				break
			}
		}
		if !found {
			// TODO error
			return fmt.Errorf("missing discount code")
		}

	}
	return nil
}

func hasUpdatePurchase(newRegistration *Info, oldRegistration *storage.Registration) bool {
	switch newRegistration.PassType.(type) {
	case *WeekendPass:
		if _, ok := oldRegistration.PassType.(*storage.WeekendPass); !ok {
			return true
		}
	case *DanceOnlyPass:
		if _, ok := oldRegistration.PassType.(*storage.DanceOnlyPass); !ok {
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

	if newRegistration.SoloJazz != nil && !oldRegistration.SoloJazz {
		return true
	}
	return false
}

func (s *Service) Update(ctx context.Context, token string, idempotencyKey string, registration *Info, redirectUrl string) (string, error) {
	s.logger.Tracef("fetching old registration id %s", registration.ID)
	oldRegistration, err := s.store.GetRegistration(ctx, registration.ID)
	if err != nil {
		return "", fmt.Errorf("error fetching registration from store: %w", err)
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
		return "", storage.ErrNotFound{Key: registration.ID}
	}

	if err := checkOldPurchases(registration, oldRegistration); err != nil {
		return "", err
	}

	returnerURL := redirectUrl
	orderID := ""
	if hasPurchase := hasUpdatePurchase(registration, oldRegistration); hasPurchase || len(oldRegistration.OrderIds) > 0 {
		s.logger.Trace("generating reference id")
		referenceId, err := uuid.NewV4()
		if err != nil {
			return "", fmt.Errorf("error generating reference id: %w", err)
		}

		locations, err := s.client.ListLocations(ctx)
		if err != nil {
			return "", fmt.Errorf("error listing locations from square: %w", err)
		}
		if len(locations) != 1 {
			return "", fmt.Errorf("found wrong number of locations %v", len(locations))
		}

		squareData, err := getSquareCatalog(ctx, s.client)
		if err != nil {
			return "", err
		}

		paymentData := &paymentData{}
		if len(oldRegistration.OrderIds) > 0 {
			paymentData, err = getSquarePayments(ctx, s.client, squareData, locations[0].Id, oldRegistration.OrderIds)
			if err != nil {
				return "", err
			}
		}

		newFullWeekend, hasNewFullWeekend := registration.PassType.(*WeekendPass)
		oldFullWeekend, hasOldFullWeekend := oldRegistration.PassType.(*storage.WeekendPass)
		newFullWeekendPurchase := false
		var newFullWeekendPurchaseTier storage.WeekendPassTier
		if hasNewFullWeekend && !hasOldFullWeekend {
			newFullWeekendPurchase = true
			newFullWeekendPurchaseTier = newFullWeekend.Tier
		} else if hasOldFullWeekend && !paymentData.weekendPassPaid {
			newFullWeekendPurchase = true
			newFullWeekendPurchaseTier = oldFullWeekend.Tier
		}
		if newFullWeekendPurchase {
			bestTier, bestCost, err := lowestInStockTier(ctx, squareData, s.client)
			if err != nil {
				return "", fmt.Errorf("error finding best tier and cost: %w", err)
			}
			if newFullWeekendPurchaseTier < bestTier {
				return "", ErrOutOfStock{
					NextTier: bestTier,
					NextCost: bestCost,
				}
			}
		}

		discounts, err := discountCodeMap(ctx, s.store, registration.DiscountCodes)
		if err != nil {
			return "", err
		}

		lineItems, err := makeLineItems(registration, squareData, paymentData, discounts)
		if err != nil {
			return "", err
		}

		order := &square.CreateOrderRequest{
			IdempotencyKey: idempotencyKey,
			Order: &square.Order{
				ReferenceId: referenceId.String(),
				LocationId:  locations[0].Id,
				LineItems:   lineItems,
			},
		}

		s.logger.Trace("creating checkout with square")
		returnerURL, orderID, err = createCheckout(ctx, s.client, locations[0].Id, idempotencyKey, order, registration.Email, redirectUrl)
	}

	s.logger.Trace("Updating registration in database")
	var orderIDs []string
	if orderID != "" {
		orderIDs = []string{orderID}
	}

	storeRegistration := &storage.Registration{
		ID:              oldRegistration.ID,
		CreatedAt:       oldRegistration.CreatedAt,
		FirstName:       registration.FirstName,
		LastName:        registration.LastName,
		StreetAddress:   registration.StreetAddress,
		City:            registration.City,
		State:           registration.State,
		ZipCode:         registration.ZipCode,
		Email:           registration.Email,
		HomeScene:       registration.HomeScene,
		IsStudent:       registration.IsStudent,
		PassType:        toStoragePassType(registration.PassType),
		MixAndMatch:     toStorageMixAndMatch(registration.MixAndMatch),
		SoloJazz:        toStorageSoloJazz(registration.SoloJazz),
		TeamCompetition: toStorageTeamCompetition(registration.TeamCompetition),
		TShirt:          toStorageTShirt(registration.TShirt),
		Housing:         registration.Housing,
		UserId:          userinfo.UserId,
		DiscountCodes:   registration.DiscountCodes,
		OrderIds:        orderIDs,
	}
	err = s.store.UpdateRegistration(ctx, storeRegistration)
	if err != nil {
		return "", fmt.Errorf("error updating registration in database: %w", err)
	}
	return returnerURL, nil
}
