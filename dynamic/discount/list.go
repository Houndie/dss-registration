package discount

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/common"
)

func (s *Service) List(ctx context.Context, token string) ([]*Bundle, error) {
	s.logger.Trace("list discount service")
	if err := common.IsAdmin(ctx, s.store, s.authorizer, s.logger, token); err != nil {
		return nil, fmt.Errorf("error checking for admin: %w", err)
	}

	s.logger.Trace("fetching all discounts from store")
	bundles, err := s.store.ListDiscounts(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing discounts from store: %w", err)
	}
	s.logger.Tracef("found %d discounts", len(bundles))

	// Early return to avoid square call
	if len(bundles) == 0 {
		s.logger.Trace("fetching all discounts from store")
		return []*Bundle{}, nil
	}

	squareData, err := common.GetSquareCatalog(ctx, s.client)
	results := make([]*Bundle, len(bundles))
	for i, b := range bundles {
		results[i], err = fromStore(b, squareData)
		if err != nil {
			return nil, fmt.Errorf("error converting store bundle to model type: %w", err)
		}
	}

	return results, nil
}
