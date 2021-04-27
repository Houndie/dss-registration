package discount

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/common"
)

func (s *Service) Add(ctx context.Context, token string, discount *Bundle) error {
	s.logger.Trace("add discount service")
	err := common.IsAllowed(ctx, s.authorizer, token, authorizer.AddDiscountPermission)
	if err != nil {
		return fmt.Errorf("error checking authorization: %w", err)
	}
	if err := s.store.AddDiscount(ctx, toStore(discount)); err != nil {
		return fmt.Errorf("could not add discount to store: %w", err)
	}

	return nil
}
