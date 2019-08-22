package finalize

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/utility"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type ErrReferenceIdNotEqual struct {
	FromSquare uuid.UUID
	FromClient uuid.UUID
}

func (e ErrReferenceIdNotEqual) Error() string {
	return fmt.Sprintf("found id %v from square, was given id %v from client (should be equal)", e.FromSquare, e.FromClient)
}

type Service struct {
	logger *logrus.Logger
	store  Store
	client SquareClient
}

func NewService(logger *logrus.Logger, store Store, client SquareClient) *Service {
	return &Service{
		logger: logger,
		store:  store,
		client: client,
	}
}

type Store interface {
	MarkRegistrationPaid(ctx context.Context, referenceId uuid.UUID, paymentId string) error
}

type SquareClient interface {
	RetrieveTransaction(ctx context.Context, locationId, transactionId string) (*square.Transaction, error)
	ListLocations(ctx context.Context) ([]*square.Location, error)
}

func (s *Service) Finalize(ctx context.Context, referenceId uuid.UUID, transactionId string) error {
	locations, err := s.client.ListLocations(ctx)
	if err != nil {
		wrap := "error listing locations from square"
		utility.LogSquareError(s.logger, err, wrap)
		return errors.Wrap(err, wrap)
	}
	if len(locations) != 1 {
		err := fmt.Errorf("found wrong number of locations %v", len(locations))
		s.logger.Error(err)
		return err
	}

	transaction, err := s.client.RetrieveTransaction(ctx, locations[0].Id, transactionId)
	if err != nil {
		// TODO find the error for a transaction that doesn't exist and report it appropriately
		wrap := "error retrieving transaction from square"
		utility.LogSquareError(s.logger, err, wrap)
		return errors.Wrap(err, wrap)
	}
	transactionReferenceId, err := uuid.FromString(transaction.ReferenceId)
	if err != nil {
		wrap := fmt.Sprintf("transaction reference id %s is not a uuid", transactionReferenceId)
		s.logger.WithError(err).Warn(wrap)
		return errors.Wrap(err, wrap)
	}
	if transactionReferenceId != referenceId {
		err := ErrReferenceIdNotEqual{
			FromSquare: transactionReferenceId,
			FromClient: referenceId,
		}
		s.logger.Debug(err)
		return err
	}

	err = s.store.MarkRegistrationPaid(ctx, referenceId)
	if err != nil {
		wrap := "error marking registration as paid"
		s.logger.Error(err)
		return errors.Wrap(err, wrap)
	}
	return nil
}
