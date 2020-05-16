package registration

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/pkg/errors"
)

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

	singleDiscounts := make([]*storage.SingleDiscount, len(discount.Discounts))
	for i, sd := range discount.Discounts {
		singleDiscounts[i] = &storage.SingleDiscount{
			Name:      sd.Name,
			AppliedTo: sd.AppliedTo,
		}
	}
	err = s.store.AddDiscount(ctx, &storage.Discount{
		Code:      discount.Code,
		Discounts: singleDiscounts,
	})
	if err != nil {
		msg := "could not add discount to store"
		s.logger.WithError(err).Error(msg)
		return errors.Wrap(err, msg)
	}

	return nil
}
