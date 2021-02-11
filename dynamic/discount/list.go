package discount

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/common"
)

func (s *Service) List(ctx context.Context, accessToken string) ([]*Bundle, error) {
	userinfo, err := s.authorizer.Userinfo(ctx, accessToken)
	if err != nil {
		return nil, fmt.Errorf("error authenticating user: %w", err)
	}

	isAdmin, err := s.store.IsAdmin(ctx, userinfo.UserID)
	if err != nil {
		return nil, fmt.Errorf("error authorizing user: %w", err)
	}

	if !isAdmin {
		return nil, ErrUnauthorized
	}

	bundles, err := s.store.ListDiscounts(ctx)
	if err != nil {
		return nil, fmt.Errorf("error listing discounts from store: %w", err)
	}

	// Early return to avoid square call
	if len(bundles) == 0 {
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
