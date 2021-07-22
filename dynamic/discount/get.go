package discount

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/storage"
)

func (s *Service) Get(ctx context.Context, code string) ([]*common.Discount, error) {
	s.logger.Tracef("in get discount service, with code %s", code)

	discounts, ok := s.squareData.Discounts.CodeDiscounts[code]
	if !ok {
		return nil, storage.ErrDiscountNotFound{
			Code: code,
		}
	}

	return discounts, nil
}
