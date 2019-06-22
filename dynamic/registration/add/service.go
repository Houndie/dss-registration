package add

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

type Store interface {
	AddRegistration(context.Context, *StoreRegistration) error
}

type Service struct {
	store  Store
	logger *logrus.Logger
}

func NewService(logger *logrus.Logger, store Store) *Service {
	return &Service{
		store:  store,
		logger: logger,
	}
}

func (s *Service) Add(ctx context.Context, registration *Registration) error {
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
