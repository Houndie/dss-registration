package discount

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/storage"
)

func (s *Service) Add(ctx context.Context, token string, discount *Bundle) error {
	s.logger.Trace("add discount service")
	s.logger.Tracef("fetching user-info for token %s", token)
	userinfo, err := s.authorizer.Userinfo(ctx, token)
	if err != nil {
		return fmt.Errorf("could not authorize user: %w", err)
	}
	s.logger.Tracef("found user %s", userinfo.UserID)
	isAdmin, err := s.store.IsAdmin(ctx, userinfo.UserID)
	if err != nil {
		return fmt.Errorf("could not fetch admin status from store: %w", err)
	}

	if !isAdmin {
		return ErrUnauthorized
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
		return fmt.Errorf("could not add discount to store: %w", err)
	}

	return nil
}
