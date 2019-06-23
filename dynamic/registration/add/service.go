package add

import (
	"context"

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

func (s *Service) Add(ctx context.Context, registration *Registration) error {
	s.logger.Trace("Fetching all items from square")
	_, err := s.client.ListCatalog(nil)
	if err != nil {
		wrap := "error fetching all items from square"
		utility.LogSquareError(s.logger, err, wrap)
		return errors.Wrap(err, wrap)
	}

	transactionID, err := uuid.NewV4()
	if err != nil {
		return err
	}

	return s.store.AddRegistration(ctx, &StoreRegistration{
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
		TransactionID:   transactionID,
	})
}
