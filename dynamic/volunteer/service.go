package volunteer

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/sirupsen/logrus"
)

type Store interface {
	InsertVolunteer(ctx context.Context, submission *storage.Volunteer) error
	VolunteerExists(ctx context.Context, userId string) (bool, error)
}

type Authorizer interface {
	Userinfo(ctx context.Context, accessToken string) (authorizer.Userinfo, error)
}

type Service struct {
	logger     *logrus.Logger
	store      Store
	authorizer Authorizer
}

func NewService(logger *logrus.Logger, store Store, authorizer Authorizer) *Service {
	return &Service{
		logger:     logger,
		store:      store,
		authorizer: authorizer,
	}
}
