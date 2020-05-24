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
	singleDiscounts := make([]*Single, len(discount.Discounts))
	for i, singleDiscount := range discount.Discounts {
		singleDiscounts[i] = &Single{
			Amount:    squareData.Discounts[singleDiscount.Name].Amount,
			Name:      singleDiscount.Name,
			AppliedTo: singleDiscount.AppliedTo,
		}

	}
	return &Bundle{
		Code:      discount.Code,
		Discounts: singleDiscounts,
	}, nil
}
