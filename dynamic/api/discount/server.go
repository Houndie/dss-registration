package discount

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/sirupsen/logrus"
)

type DiscountService interface {
	Get(ctx context.Context, code string) ([]*common.Discount, error)
}

type Server struct {
	service DiscountService
	logger  logrus.Logger
}

func NewServer(service DiscountService) *Server {
	return &Server{
		service: service,
	}
}
