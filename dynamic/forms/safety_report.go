package forms

import (
	"context"
	"fmt"
	"time"

	"github.com/Houndie/dss-registration/dynamic/sendinblue"
)

func (s *Service) SafetyReport(ctx context.Context, occurredOn time.Time, description string, severity int, issuesBefore bool, resolution, name, email, phone, recaptchaResponse string) error {
	success, err := s.recaptchaClient.VerifySite(ctx, recaptchaResponse)
	if err != nil {
		return fmt.Errorf("error verifying recaptcha: %w", err)
	}
	if !success {
		return ErrRecaptchaFailed
	}

	var replyTo *sendinblue.EmailPerson
	if email != "" {
		replyTo = &sendinblue.EmailPerson{
			Email: email,
			Name:  name,
		}
	}
	_, err = s.mailClient.SendSMTPEmail(ctx, &sendinblue.SMTPEmailParams{
		To: []*sendinblue.EmailPerson{
			{
				Name:  "Dayton Swing Smackdown",
				Email: "info@daytonswingsmackdown.com",
			},
		},
		ReplyTo: replyTo,
		Params: map[string]interface{}{
			"name":          name,
			"email":         email,
			"occurred_on":   occurredOn.In(s.location).String(),
			"description":   description,
			"severity":      severity,
			"issues_before": issuesBefore,
			"resolution":    resolution,
			"phone":         phone,
		},
		TemplateID: 2,
	})
	if err != nil {
		return fmt.Errorf("error sending safety report email: %w", err)
	}
	return nil
}
