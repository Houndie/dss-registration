package add

import (
	"context"

	"github.com/sirupsen/logrus"
)

type Store interface {
	AddRegistration(context.Context, *Registration) error
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
	return s.store.AddRegistration(ctx, registration)
}
