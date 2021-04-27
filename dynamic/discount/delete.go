package discount

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/common"
)

func (s *Service) Delete(ctx context.Context, token, code string) error {
	s.logger.Trace("delete discount service")
	err := common.IsAllowed(ctx, s.authorizer, token, authorizer.DeleteDiscountPermission)
	if err != nil {
		return fmt.Errorf("error checking authorization: %w", err)
	}

	if err := s.store.DeleteDiscount(ctx, code); err != nil {
		return fmt.Errorf("could not delete discount from store: %w", err)
	}

	return nil
}
