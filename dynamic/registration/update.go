package registration

import (
	"context"
	"fmt"
	"time"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/storage"
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

	if oldRegistration.SoloJazz != nil && newRegistration.SoloJazz == nil {
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

func (s *Service) Update(ctx context.Context, token string, registration *Info) (*Info, error) {
	s.logger.Tracef("fetching old registration id %s", registration.ID)

	s.logger.Tracef("fetching user-info for token %s", token)
	userinfo, err := s.authorizer.GetUserinfo(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("could not authorize user: %w", err)
	}
	s.logger.Tracef("found user %s", userinfo.UserID())

	oldRegistration, err := s.store.GetRegistration(ctx, registration.ID)
	if err != nil {
		return nil, fmt.Errorf("error fetching registration from store: %w", err)
	}

	isAdmin := userinfo.IsAllowed(s.permissionConfig.Update)

	if !isAdmin && oldRegistration.UserID != userinfo.UserID() {
		s.logger.WithError(err).Debug("user id does not match that of found registration")
		s.logger.WithError(err).Tracef("registration provided user id %s, user provided %s", oldRegistration.UserID, userinfo.UserID())
		return nil, storage.ErrNoRegistrationForID{ID: registration.ID}
	}

	if err := checkOldPurchases(registration, oldRegistration); err != nil {
		return nil, err
	}

	s.logger.Trace("Fetching all locations from square")
	locationListRes, err := s.client.Locations.List(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error listing locations from square: %w", err)
	}
	if len(locationListRes.Locations) != 1 {
		return nil, fmt.Errorf("found wrong number of locations %v", len(locationListRes.Locations))
	}
	locationID := locationListRes.Locations[0].ID

	pdMap, err := common.GetSquarePayments(ctx, s.client, s.squareData.PurchaseItems, locationID, map[string][]string{registration.ID: oldRegistration.OrderIDs})
	if err != nil {
		return nil, err
	}

	pd := pdMap[registration.ID]

	if err := paymentCheck(registration, isAdmin, oldRegistration, pd); err != nil {
		return nil, err
	}

	if registration.Enabled != oldRegistration.Enabled && !isAdmin {
		return nil, authorizer.Unauthorized
	}

	if oldRegistration.CreatedAt.Sub(registration.CreatedAt) >= 1*time.Second {
		return nil, ErrHasImmutableField{Field: "created_at"}
	}

	s.logger.Trace("Updating registration in database")
	storeRegistration := toStorageRegistration(registration)
	storeRegistration.UserID = oldRegistration.UserID
	storeRegistration.OrderIDs = oldRegistration.OrderIDs

	err = s.store.UpdateRegistration(ctx, storeRegistration)
	if err != nil {
		return nil, fmt.Errorf("error updating registration in database: %w", err)
	}

	returnInfo := fromStorageRegistration(storeRegistration, pd)

	return returnInfo, nil
}
