package forms

import (
	"context"
	"time"
)

type FormsService interface {
	ContactUs(ctx context.Context, name, email, message, recaptchaResponse string) error
	SafetyReport(ctx context.Context, occurredOn time.Time, description string, severity int, issuesBefore bool, resolution, name, email, phone, recaptchaRespone string) error
}

type Server struct {
	service FormsService
}

func NewServer(service FormsService) *Server {
	return &Server{
		service: service,
	}
}
