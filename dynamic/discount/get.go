package discount

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/common"
)

func (s *Service) Get(ctx context.Context, code string) (*Bundle, error) {
	s.logger.Tracef("in get discount service, with code %s", code)

	discount, err := s.store.GetDiscount(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("error getting discount from store: %w", err)
	}

	squareData, err := common.GetSquareCatalog(ctx, s.client)
	bundle, err := fromStore(discount, squareData)
	if err != nil {
		return nil, fmt.Errorf("error converting store bundle to model type: %w", err)
	}
	return bundle, nil
}
