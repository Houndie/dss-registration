package getdiscount

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/registration/common"
	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Store interface {
	GetDiscount(context.Context, string) ([]*common.StoreDiscount, error)
}

type SquareClient interface {
	ListCatalog(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator
}

type Service struct {
	client SquareClient
	store  Store
	logger *logrus.Logger
}

func NewService(logger *logrus.Logger, store Store, client SquareClient) *Service {
	return &Service{
		store:  store,
		logger: logger,
		client: client,
	}
}

func (s *Service) GetDiscount(ctx context.Context, code string) ([]*Discount, error) {
	s.logger.Tracef("in get discount service, with code %s", code)

	storeDiscounts, err := s.store.GetDiscount(ctx, code)
	if err != nil {
		switch errors.Cause(err).(type) {
		case ErrDiscountDoesNotExist:
			s.logger.Debug(err)
			return nil, err
		default:
			msg := "error getting discount from store"
			s.logger.WithError(err).Error(msg)
			return nil, errors.Wrap(err, msg)
		}
	}
	discounts := []*Discount{}

	objects := s.client.ListCatalog(ctx, []square.CatalogObjectType{square.CatalogObjectTypeDiscount})
	for objects.Next() {
		squareDiscount, ok := objects.Value().CatalogObjectType.(*square.CatalogDiscount)
		if !ok {
			s.logger.Error("somehow getting non-discount objects from square?")
			continue
		}
		s.logger.Tracef("found square object %s", squareDiscount.Name)
		for _, storeDiscount := range storeDiscounts {
			if storeDiscount.Name != squareDiscount.Name {
				continue
			}
			var itemDiscount common.ItemDiscount
			switch t := squareDiscount.DiscountType.(type) {
			case *square.CatalogDiscountFixedAmount:
				itemDiscount = &common.DollarDiscount{
					Amount: t.AmountMoney.Amount,
				}
			case *square.CatalogDiscountVariableAmount:
				itemDiscount = &common.DollarDiscount{
					Amount: t.AmountMoney.Amount,
				}
			case *square.CatalogDiscountFixedPercentage:
				itemDiscount = &common.PercentDiscount{
					Amount: t.Percentage,
				}
			case *square.CatalogDiscountVariablePercentage:
				itemDiscount = &common.PercentDiscount{
					Amount: t.Percentage,
				}
			default:
				err := errors.New("found unknown catalog discount type")
				s.logger.Error(err)
				return nil, err

			}
			discounts = append(discounts, &Discount{
				AppliedTo:    storeDiscount.AppliedTo,
				ItemDiscount: itemDiscount,
			})
			break
		}

	}
	if err := objects.Error(); err != nil {
		msg := "error getting discounts from list catalog"
		s.logger.WithError(err).Error(msg)
		return nil, errors.Wrap(err, msg)
	}
	if len(discounts) == 0 {
		err := errors.New("code %s exists in store, but not on square?")
		s.logger.Error(err)
		return nil, err
	}

	return discounts, nil
}
