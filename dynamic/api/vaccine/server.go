package vaccine

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/vaccine"
)

type VaccineService interface {
	Get(ctx context.Context, token, id string) (vaccine.Info, error)
	Upload(ctx context.Context, token string, filesize int64, id string) (string, error)
	Approve(ctx context.Context, token string, id string) error
	Reject(ctx context.Context, token string, id string, reason string) error
}

type Server struct {
	service VaccineService
}

func NewServer(service VaccineService) *Server {
	return &Server{
		service: service,
	}
}
