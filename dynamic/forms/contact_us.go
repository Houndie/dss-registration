package forms

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/sendinblue"
)

func (s *Service) ContactUs(ctx context.Context, name, email, message, recaptchaResponse string) error {
	success, err := s.recaptchaClient.VerifySite(ctx, recaptchaResponse)
	if err != nil {
		return fmt.Errorf("error verifying recaptcha: %w", err)
	}
	if !success {
		return ErrRecaptchaFailed
	}
	_, err = s.mailClient.SendSMTPEmail(ctx, &sendinblue.SMTPEmailParams{
		To: []*sendinblue.EmailPerson{
			{
				Name:  "Dayton Swing Smackdown",
				Email: "info@daytonswingsmackdown.com",
			},
		},
		ReplyTo: &sendinblue.EmailPerson{
			Name:  name,
			Email: email,
		},
		Params: struct {
			Name    string `json:"name"`
			Email   string `json:"email"`
			Message string `json:"message"`
		}{
			Name:    name,
			Email:   email,
			Message: message,
		},
		TemplateID: 1,
	})
	if err != nil {
		return fmt.Errorf("error sending contact us email: %w", err)
	}
	return nil
}
