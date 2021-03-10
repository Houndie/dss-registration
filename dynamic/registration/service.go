package registration

import (
	"context"
	"time"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/sendinblue"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/sirupsen/logrus"
)

type MailClient interface {
	SendSMTPEmail(ctx context.Context, params *sendinblue.SMTPEmailParams) (string, error)
}

type Store interface {
	AddRegistration(context.Context, *storage.Registration) (string, error)
	GetDiscount(context.Context, string) (*storage.Discount, error)
	GetRegistrationsByUser(ctx context.Context, userId string) ([]*storage.Registration, error)
	GetRegistration(ctx context.Context, id string) (*storage.Registration, error)
	IsAdmin(context.Context, string) (bool, error)
	UpdateRegistration(ctx context.Context, r *storage.Registration) error
}

type Service struct {
	client         common.SquareClient
	logger         *logrus.Logger
	active         bool
	useMailSandbox bool
	authorizer     Authorizer
	store          Store
	mailClient     MailClient
}

func NewService(active, useMailSandbox bool, logger *logrus.Logger, client common.SquareClient, authorizer Authorizer, store Store, mailClient MailClient) *Service {
	return &Service{
		active:         active,
		useMailSandbox: useMailSandbox,
		client:         client,
		logger:         logger,
		authorizer:     authorizer,
		store:          store,
		mailClient:     mailClient,
	}
}

type Authorizer interface {
	GetUserinfo(ctx context.Context, accessToken string) (authorizer.Userinfo, error)
}

type Summary struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	Paid      bool
}
