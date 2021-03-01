package discount

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/storage"
)

func (s *Service) Update(ctx context.Context, token, oldCode string, newDiscount *Bundle) error {
	s.logger.Trace("update discount service")
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

	singleDiscounts := make([]*storage.SingleDiscount, len(newDiscount.Discounts))
	for i, sd := range newDiscount.Discounts {
		singleDiscounts[i] = &storage.SingleDiscount{
			Name:      sd.Name,
			AppliedTo: sd.AppliedTo,
		}
	}
	err = s.store.UpdateDiscount(ctx, oldCode, &storage.Discount{
		Code:      newDiscount.Code,
		Discounts: singleDiscounts,
	})
	if err != nil {
		return fmt.Errorf("could not update discount to store: %w", err)
	}

	return nil
}
