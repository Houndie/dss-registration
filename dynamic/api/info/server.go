package info

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/info"
)

type Service interface {
	Health(ctx context.Context) info.Healthiness
	Version() string
}

type Server struct {
	service Service
}

func NewServer(service Service) *Server {
	return &Server{
		service: service,
	}
}
