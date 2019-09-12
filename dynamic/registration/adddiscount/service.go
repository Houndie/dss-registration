package adddiscount

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Store interface {
	IsAdmin(context.Context, string) (bool, error)
	AddDiscount(context.Context, *Discount) error
}

type Authorizer interface {
	Userinfo(ctx context.Context, accessToken string) (*authorizer.Userinfo, error)
}

type Service struct {
	store      Store
	authorizer Authorizer
	logger     *logrus.Logger
}

func NewService(logger *logrus.Logger, store Store, authorizer Authorizer) *Service {
	return &Service{
		store:      store,
		logger:     logger,
		authorizer: authorizer,
	}
}

func (s *Service) AddDiscount(ctx context.Context, token string, discount *Discount) error {
	s.logger.Trace("add discount service")
	s.logger.Tracef("fetching user-info for token %s", token)
	userinfo, err := s.authorizer.Userinfo(ctx, token)
	if err != nil {
		msg := "could not authorize user"
		s.logger.WithError(err).Debug(msg)
		return errors.Wrap(err, msg)
	}
	s.logger.Tracef("found user %s", userinfo.UserId)
	isAdmin, err := s.store.IsAdmin(ctx, userinfo.UserId)
	if err != nil {
		msg := "could not fetch admin status from store"
		s.logger.WithError(err).Error(msg)
		return errors.Wrap(err, msg)
	}

	if !isAdmin {
		e := ErrUnauthorized{}
		s.logger.Debug(e)
		return e
	}

	err = s.store.AddDiscount(ctx, discount)
	if err != nil {
		msg := "could not add discount to store"
		s.logger.WithError(err).Error(msg)
		return errors.Wrap(err, msg)
	}

	return nil
}
