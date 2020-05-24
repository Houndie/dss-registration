package discount

import (
	"context"
	"errors"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/sirupsen/logrus"
)

type Store interface {
	GetDiscount(context.Context, string) (*storage.Discount, error)
	AddDiscount(context.Context, *storage.Discount) error
	IsAdmin(context.Context, string) (bool, error)
}

type Authorizer interface {
	Userinfo(ctx context.Context, accessToken string) (*authorizer.Userinfo, error)
}

func NewService(store Store, client common.SquareClient, logger *logrus.Logger, authorizer Authorizer) *Service {
	return &Service{
		store:      store,
		client:     client,
		logger:     logger,
		authorizer: authorizer,
	}
}

type Service struct {
	store      Store
	client     common.SquareClient
	logger     *logrus.Logger
	authorizer Authorizer
}

type Single struct {
	Amount    common.DiscountAmount
	Name      string
	AppliedTo storage.PurchaseItem
}

type Bundle struct {
	Code      string
	Discounts []*Single
}

var ErrUnauthorized = errors.New("User is not authorized for this operation")
