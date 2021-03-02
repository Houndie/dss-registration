package discount

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/common"
)

func (s *Service) Delete(ctx context.Context, token, code string) error {
	s.logger.Trace("delete discount service")
	if err := common.IsAdmin(ctx, s.store, s.authorizer, s.logger, token); err != nil {
		return fmt.Errorf("error checking for admin: %w", err)
	}

	if err := s.store.DeleteDiscount(ctx, code); err != nil {
		return fmt.Errorf("could not delete discount from store: %w", err)
	}

	return nil
}
