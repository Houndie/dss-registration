package discount

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/common"
)

func (s *Service) Add(ctx context.Context, token string, discount *Bundle) error {
	s.logger.Trace("add discount service")
	if err := common.IsAdmin(ctx, s.store, s.authorizer, s.logger, token); err != nil {
		return fmt.Errorf("error checking for admin: %w", err)
	}
	if err := s.store.AddDiscount(ctx, toStore(discount)); err != nil {
		return fmt.Errorf("could not add discount to store: %w", err)
	}

	return nil
}
