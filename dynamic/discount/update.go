package discount

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/common"
)

func (s *Service) Update(ctx context.Context, token, oldCode string, newDiscount *Bundle) error {
	s.logger.Trace("update discount service")
	if err := common.IsAdmin(ctx, s.store, s.authorizer, s.logger, token); err != nil {
		return fmt.Errorf("error checking for admin: %w", err)
	}

	if err := s.store.UpdateDiscount(ctx, oldCode, toStore(newDiscount)); err != nil {
		return fmt.Errorf("could not update discount to store: %w", err)
	}

	return nil
}
