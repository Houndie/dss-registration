package discount

import (
	"context"
	"errors"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/square-go"
	"github.com/sirupsen/logrus"
)

type Store interface {
	GetDiscount(context.Context, string) (*storage.Discount, error)
	AddDiscount(context.Context, *storage.Discount) error
	IsAdmin(context.Context, string) (bool, error)
	ListDiscounts(context.Context) ([]*storage.Discount, error)
	UpdateDiscount(ctx context.Context, oldCode string, newDiscount *storage.Discount) error
	DeleteDiscount(ctx context.Context, code string) error
}

type Authorizer interface {
	GetUserinfo(ctx context.Context, accessToken string) (authorizer.Userinfo, error)
}

func NewService(store Store, client *square.Client, logger *logrus.Logger, authorizer Authorizer) *Service {
	return &Service{
		store:      store,
		client:     client,
		logger:     logger,
		authorizer: authorizer,
	}
}

type Service struct {
	store      Store
	client     *square.Client
	logger     *logrus.Logger
	authorizer Authorizer
}

type Single struct {
	Amount    DiscountAmount
	Name      string
	AppliedTo storage.PurchaseItem
}

type Bundle struct {
	Code      string
	Discounts []*Single
}

type DollarDiscount int

type PercentDiscount string

type SquareNotFound struct{}

type DiscountAmount interface {
	isDiscountAmount()
}

func (DollarDiscount) isDiscountAmount()  {}
func (PercentDiscount) isDiscountAmount() {}
func (SquareNotFound) isDiscountAmount()  {}

func amountFromSquare(name string, squareData *common.SquareData) (DiscountAmount, error) {
	squareDiscount, ok := squareData.Discounts[name]
	if !ok {
		return SquareNotFound{}, nil
	}

	switch sd := squareDiscount.Amount.(type) {
	case common.DollarDiscount:
		return DollarDiscount(int(sd)), nil
	case common.PercentDiscount:
		return PercentDiscount(string(sd)), nil
	default:
		return nil, errors.New("unknown discount type from square data")
	}
}

func fromStore(b *storage.Discount, squareData *common.SquareData) (*Bundle, error) {
	singleDiscounts := make([]*Single, len(b.Discounts))
	for i, singleDiscount := range b.Discounts {
		amount, err := amountFromSquare(singleDiscount.Name, squareData)
		if err != nil {
			return nil, err
		}
		singleDiscounts[i] = &Single{
			Amount:    amount,
			Name:      singleDiscount.Name,
			AppliedTo: singleDiscount.AppliedTo,
		}
	}

	return &Bundle{
		Code:      b.Code,
		Discounts: singleDiscounts,
	}, nil
}

func toStore(b *Bundle) *storage.Discount {
	singleDiscounts := make([]*storage.SingleDiscount, len(b.Discounts))
	for i, sd := range b.Discounts {
		singleDiscounts[i] = &storage.SingleDiscount{
			Name:      sd.Name,
			AppliedTo: sd.AppliedTo,
		}
	}
	return &storage.Discount{
		Code:      b.Code,
		Discounts: singleDiscounts,
	}
}
