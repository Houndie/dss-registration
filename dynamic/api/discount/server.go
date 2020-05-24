package discount

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/discount"
	"github.com/sirupsen/logrus"
)

type DiscountService interface {
	Add(ctx context.Context, token string, discount *discount.Bundle) error
	Get(ctx context.Context, code string) (*discount.Bundle, error)
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
