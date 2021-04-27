package forms

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Houndie/dss-registration/dynamic/sendinblue"
)

type MailClient interface {
	SendSMTPEmail(ctx context.Context, params *sendinblue.SMTPEmailParams) (string, error)
}

type RecaptchaClient interface {
	VerifySite(ctx context.Context, recaptchaResponse string) (bool, error)
}

type Service struct {
	mailClient      MailClient
	recaptchaClient RecaptchaClient
	location        *time.Location
}

func NewService(mailClient MailClient, recaptchaClient RecaptchaClient) (*Service, error) {
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		return nil, fmt.Errorf("error constructing location for time presentation: %w", err)
	}

	return &Service{
		mailClient:      mailClient,
		location:        location,
		recaptchaClient: recaptchaClient,
	}, nil
}

var ErrRecaptchaFailed = errors.New("recaptcha response is not valid")
