package registration

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/gofrs/uuid"
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
	userinfo, err := s.authorizer.GetUserinfo(ctx, token)
	if err != nil {
		return "", fmt.Errorf("could not authorize user: %w", err)
	}
	s.logger.Tracef("found user %s", userinfo.UserID())

	if oldRegistration.UserID != userinfo.UserID() {
		s.logger.WithError(err).Debug("user id does not match that of found registration")
		s.logger.WithError(err).Tracef("registration provided user id %s, user provided %s", oldRegistration.UserID, userinfo.UserID())
		return "", storage.ErrNoRegistrationForID{ID: registration.ID}
	}

	if err := checkOldPurchases(registration, oldRegistration); err != nil {
		return "", err
	}

	returnerURL := redirectUrl
	orderID := ""
	if hasPurchase := hasUpdatePurchase(registration, oldRegistration); hasPurchase || len(oldRegistration.OrderIDs) > 0 {
		s.logger.Trace("generating reference id")
		referenceID, err := uuid.NewV4()
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

		squareData, err := common.GetSquareCatalog(ctx, s.client)
		if err != nil {
			return "", err
		}

		paymentData := &common.PaymentData{}
		if len(oldRegistration.OrderIDs) > 0 {
			paymentData, err = common.GetSquarePayments(ctx, s.client, squareData, locations[0].ID, oldRegistration.OrderIDs)
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
		} else if hasOldFullWeekend && !paymentData.WeekendPassPaid {
			newFullWeekendPurchase = true
			newFullWeekendPurchaseTier = oldFullWeekend.Tier
		}
		if newFullWeekendPurchase {
			bestTier, bestCost, err := common.LowestInStockTier(ctx, squareData, s.client)
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

		lineItems, lineDiscounts, err := makeLineItems(registration, squareData, paymentData, discounts)
		if err != nil {
			return "", err
		}

		order := &square.CreateOrderRequest{
			IdempotencyKey: idempotencyKey,
			Order: &square.Order{
				ReferenceID: referenceID.String(),
				LocationID:  locations[0].ID,
				LineItems:   lineItems,
				Discounts:   lineDiscounts,
			},
		}

		s.logger.Trace("creating checkout with square")
		returnerURL, orderID, err = common.CreateCheckout(ctx, s.client, locations[0].ID, idempotencyKey, order, registration.Email, redirectUrl)
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
		UserID:          userinfo.UserID(),
		DiscountCodes:   registration.DiscountCodes,
		OrderIDs:        orderIDs,
	}
	err = s.store.UpdateRegistration(ctx, storeRegistration)
	if err != nil {
		return "", fmt.Errorf("error updating registration in database: %w", err)
	}
	return returnerURL, nil
}
