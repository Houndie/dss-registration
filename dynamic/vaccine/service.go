package vaccine

import (
	"context"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/object"
	"github.com/Houndie/dss-registration/dynamic/sendinblue"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/sirupsen/logrus"
)

type Store interface {
	GetVaccine(context.Context, string) (bool, error)
	ApproveVaccine(context.Context, string, bool) error
	GetRegistration(ctx context.Context, id string) (*storage.Registration, error)
}

type Service struct {
	logger           *logrus.Logger
	authorizer       Authorizer
	store            Store
	objectClient     object.Client
	permissionConfig *PermissionConfig
	mailClient       MailClient
}

type PermissionConfig struct {
	Approve string
	Get     string
	Upload  string
}

type MailClient interface {
	SendSMTPEmail(ctx context.Context, params *sendinblue.SMTPEmailParams) (string, error)
}

func NewService(logger *logrus.Logger, authorizer Authorizer, store Store, objectClient object.Client, mailClient MailClient, permissionConfig *PermissionConfig) *Service {
	return &Service{
		logger:           logger,
		authorizer:       authorizer,
		store:            store,
		objectClient:     objectClient,
		permissionConfig: permissionConfig,
		mailClient:       mailClient,
	}
}

type Authorizer interface {
	GetUserinfo(ctx context.Context, accessToken string) (authorizer.Userinfo, error)
}
