package registration

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/pkg/errors"
)

func (s *Service) Get(ctx context.Context, token, registrationID string) (*Info, error) {
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
	r, err := s.store.GetRegistration(ctx, registrationID)
	if err != nil {
		return nil, fmt.Errorf("error getting registration: %w", err)
	}
	s.logger.Trace("found registration")

	if r.UserId != userinfo.UserId {
		s.logger.WithError(err).Debug("user id does not match that of found registration")
		s.logger.WithError(err).Tracef("registration provided user id %s, user provided %s", r.UserId, userinfo.UserId)
		return nil, storage.ErrNotFound{Key: registrationID}
	}

	pd := &paymentData{}
	if len(r.OrderIds) > 0 {
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

		squareData, err := getSquareCatalog(ctx, s.client)
		if err != nil {
			return nil, err
		}

		pd, err = getSquarePayments(ctx, s.client, squareData, locations[0].Id, r.OrderIds)
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
		PassType:        fromStoragePassType(r.PassType, pd.weekendPassPaid, pd.danceOnlyPaid),
		MixAndMatch:     fromStorageMixAndMatch(r.MixAndMatch, pd.mixAndMatchPaid),
		SoloJazz:        fromStorageSoloJazz(r.SoloJazz, pd.soloJazzPaid),
		TeamCompetition: fromStorageTeamCompetition(r.TeamCompetition, pd.teamCompetitionPaid),
		TShirt:          fromStorageTShirt(r.TShirt, pd.tShirtPaid),
		Housing:         r.Housing,
		DiscountCodes:   r.DiscountCodes,
		CreatedAt:       r.CreatedAt,
	}, nil
}
