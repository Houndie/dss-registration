package registration

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/pkg/errors"
)

func (s *Service) GetDiscount(ctx context.Context, code string) (*Discount, error) {
	s.logger.Tracef("in get discount service, with code %s", code)

	discount, err := s.store.GetDiscount(ctx, code)
	if err != nil {
		switch errors.Cause(err).(type) {
		case storage.ErrDiscountDoesNotExist:
			s.logger.Debug(err)
			return nil, err
		default:
			msg := "error getting discount from store"
			s.logger.WithError(err).Error(msg)
			return nil, errors.Wrap(err, msg)
		}
	}

	squareData, err := getSquareCatalog(ctx, s.client)
	singleDiscounts := make([]*SingleDiscount, len(discount.Discounts))
	for i, singleDiscount := range discount.Discounts {
		singleDiscounts[i] = &SingleDiscount{
			Amount:    squareData.discounts[singleDiscount.Name].amount,
			Name:      singleDiscount.Name,
			AppliedTo: singleDiscount.AppliedTo,
		}

	}
	return &Discount{
		Code:      discount.Code,
		Discounts: singleDiscounts,
	}, nil
}
