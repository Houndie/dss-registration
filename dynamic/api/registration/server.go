package registration

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/registration"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/sirupsen/logrus"
)

type RegistrationService interface {
	Add(ctx context.Context, registration *registration.Info, redirectUrl, idempotencyKey, accessToken string) (string, error)
	Get(ctx context.Context, token, registrationID string) (*registration.Info, error)
	Populate(ctx context.Context) (storage.WeekendPassTier, error)
	SummaryByUser(ctx context.Context, token string) ([]*registration.Summary, error)
	Update(ctx context.Context, token string, idempotencyKey string, registration *registration.Info, redirectUrl string) (string, error)
}

type Server struct {
	service RegistrationService
	logger  logrus.Logger
}

func NewServer(service RegistrationService) *Server {
	return &Server{
		service: service,
	}
}
