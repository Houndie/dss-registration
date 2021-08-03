package registration

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/registration"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/sirupsen/logrus"
)

type RegistrationService interface {
	Add(ctx context.Context, registration *registration.Info, accessToken string) (*registration.Info, error)
	Get(ctx context.Context, token, registrationID string) (*registration.Info, error)
	Populate(ctx context.Context) (storage.WeekendPassTier, error)
	SummaryByUser(ctx context.Context, token string) ([]*registration.Summary, error)
	Update(ctx context.Context, token string, registration *registration.Info) (*registration.Info, error)
	Pay(ctx context.Context, id, redirectURL, idempotencyKey, token string) (string, error)
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
