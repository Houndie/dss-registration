package registration

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/storage"
)

func (s *Service) Get(ctx context.Context, token, registrationID string) (*Info, error) {
	s.logger.Trace("In get by id service")
	userinfo, err := s.authorizer.GetUserinfo(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("could not authorize user: %w", err)
	}

	s.logger.Tracef("fetching registrations for user %s", userinfo.UserID())
	r, err := s.store.GetRegistration(ctx, registrationID)
	if err != nil {
		return nil, fmt.Errorf("error getting registration: %w", err)
	}
	s.logger.Trace("found registration")

	if r.UserID != userinfo.UserID() {
		s.logger.WithError(err).Debug("user id does not match that of found registration")
		s.logger.WithError(err).Tracef("registration provided user id %s, user provided %s", r.UserID, userinfo.UserID())
		return nil, storage.ErrNotFound{Key: registrationID}
	}

	pd := &common.PaymentData{}
	if len(r.OrderIDs) > 0 {
		s.logger.Trace("fetching locations from square")
		locations, err := s.client.Locations.List(ctx)
		if err != nil {
			return nil, fmt.Errorf("error listing locations from square: %w", err)
		}
		s.logger.Tracef("found %d locations", len(locations))

		if len(locations) != 1 {
			return nil, fmt.Errorf("found unexpected number of locations %d", len(locations))
		}
		s.logger.Tracef("found location %s", locations[0].ID)

		squareData, err := common.GetSquareCatalog(ctx, s.client)
		if err != nil {
			return nil, err
		}

		pd, err = common.GetSquarePayments(ctx, s.client, squareData, locations[0].ID, r.OrderIDs)
		if err != nil {
			return nil, err
		}
	}

	return &Info{
		ID:              r.ID,
		FirstName:       r.FirstName,
		LastName:        r.LastName,
		StreetAddress:   r.StreetAddress,
		City:            r.City,
		State:           r.State,
		ZipCode:         r.ZipCode,
		Email:           r.Email,
		HomeScene:       r.HomeScene,
		IsStudent:       r.IsStudent,
		PassType:        fromStoragePassType(r.PassType, pd.WeekendPassPaid, pd.DanceOnlyPaid),
		MixAndMatch:     fromStorageMixAndMatch(r.MixAndMatch, pd.MixAndMatchPaid),
		SoloJazz:        fromStorageSoloJazz(r.SoloJazz, pd.SoloJazzPaid),
		TeamCompetition: fromStorageTeamCompetition(r.TeamCompetition, pd.TeamCompetitionPaid),
		TShirt:          fromStorageTShirt(r.TShirt, pd.TShirtPaid),
		Housing:         r.Housing,
		DiscountCodes:   r.DiscountCodes,
		CreatedAt:       r.CreatedAt,
	}, nil
}
