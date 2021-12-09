package vaccine

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/sendinblue"
	"github.com/Houndie/dss-registration/dynamic/storage"
)

func (s *Service) Approve(ctx context.Context, token string, id string) error {
	s.logger.Trace("vaccine upload service")
	userinfo, err := s.authorizer.GetUserinfo(ctx, token)
	if err != nil {
		return fmt.Errorf("could not authorize user: %w", err)
	}

	if !userinfo.IsAllowed(s.permissionConfig.Approve) {
		s.logger.WithError(err).Debug("permission denied")
		return storage.ErrNotFound{Key: id}
	}

	s.logger.Tracef("fetching registrations for user %s", userinfo.UserID())
	registration, err := s.store.GetRegistration(ctx, id)
	if err != nil {
		return fmt.Errorf("error getting registration: %w", err)
	}
	s.logger.Trace("found registration")

	if err := s.store.ApproveVaccine(ctx, id, true); err != nil {
		return fmt.Errorf("error approving vaccine in store: %w", err)
	}

	if err := s.objectClient.Delete(ctx, id); err != nil {
		return fmt.Errorf("error deleting vaccine proof: %w", err)
	}

	_, err = s.mailClient.SendSMTPEmail(ctx, &sendinblue.SMTPEmailParams{
		To: []*sendinblue.EmailPerson{
			{
				Name:  fmt.Sprintf("%s %s", registration.FirstName, registration.LastName),
				Email: registration.Email,
			},
		},
		/*BCC: []*sendinblue.EmailPerson{
			{
				Name:  "Dayton Swing Smackdown",
				Email: "info@daytonswingsmackdown.com",
			},
		},*/
		TemplateID: 4,
	})
	if err != nil {
		return fmt.Errorf("error sending vaccine approved email: %w", err)
	}

	return nil
}
