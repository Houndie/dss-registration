package registration

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/sendinblue"
	"github.com/Houndie/dss-registration/dynamic/storage"
)

func (s *Service) Add(ctx context.Context, registration *Info, accessToken string) (*Info, error) {
	s.logger.Trace("in add registration service")
	if !s.active {
		return nil, ErrRegistrationDisabled
	}

	s.logger.Trace("found access token")
	userinfo, err := s.authorizer.GetUserinfo(ctx, accessToken)
	if err != nil {
		return nil, fmt.Errorf("error fetching userinfo: %w", err)
	}
	userid := userinfo.UserID()

	s.logger.Trace("Adding registration to database")
	storeRegistration := &storage.Registration{
		FirstName:       registration.FirstName,
		LastName:        registration.LastName,
		StreetAddress:   registration.StreetAddress,
		City:            registration.City,
		State:           registration.State,
		ZipCode:         registration.ZipCode,
		Email:           registration.Email,
		HomeScene:       registration.HomeScene,
		IsStudent:       registration.IsStudent,
		PassType:        toStoragePassType(registration.PassType),
		MixAndMatch:     toStorageMixAndMatch(registration.MixAndMatch),
		SoloJazz:        toStorageSoloJazz(registration.SoloJazz),
		TeamCompetition: toStorageTeamCompetition(registration.TeamCompetition),
		TShirt:          toStorageTShirt(registration.TShirt),
		Housing:         registration.Housing,
		UserID:          userid,
		DiscountCodes:   registration.DiscountCodes,
	}

	id, err := s.store.AddRegistration(ctx, storeRegistration)
	if err != nil {
		return nil, fmt.Errorf("error adding registration to database: %w", err)
	}

	s.logger.Trace("sending registration email")
	mailParams, err := toMailParams(registration)
	if err != nil {
		return nil, fmt.Errorf("error generating mail parameters")
	}
	_, err = s.mailClient.SendSMTPEmail(ctx, &sendinblue.SMTPEmailParams{
		To: []*sendinblue.EmailPerson{
			{
				Name:  fmt.Sprintf("%s %s", registration.FirstName, registration.LastName),
				Email: registration.Email,
			},
		},
		BCC: []*sendinblue.EmailPerson{
			{
				Name:  "Dayton Swing Smackdown",
				Email: "info@daytonswingsmackdown.com",
			},
		},
		Params:     mailParams,
		TemplateID: 3,
	})
	if err != nil {
		return nil, fmt.Errorf("error sending registration email: %w", err)
	}
	returnInfo := &Info{
		ID:            id,
		FirstName:     registration.FirstName,
		LastName:      registration.LastName,
		StreetAddress: registration.StreetAddress,
		City:          registration.City,
		State:         registration.State,
		ZipCode:       registration.ZipCode,
		Email:         registration.Email,
		HomeScene:     registration.HomeScene,
		IsStudent:     registration.IsStudent,
		Housing:       registration.Housing,
		DiscountCodes: registration.DiscountCodes,
	}

	switch p := registration.PassType.(type) {
	case *WeekendPass:
		returnInfo.PassType = &WeekendPass{
			Level: p.Level,
			Tier:  p.Tier,
		}
	case *DanceOnlyPass:
		returnInfo.PassType = &DanceOnlyPass{}
	case *NoPass:
		returnInfo.PassType = &NoPass{}
	}

	if registration.MixAndMatch != nil {
		returnInfo.MixAndMatch = &MixAndMatch{Role: registration.MixAndMatch.Role}
	}

	if registration.SoloJazz != nil {
		returnInfo.SoloJazz = &SoloJazz{}
	}

	if registration.TeamCompetition != nil {
		returnInfo.TeamCompetition = &TeamCompetition{Name: registration.TeamCompetition.Name}
	}

	if registration.TShirt != nil {
		returnInfo.TShirt = &TShirt{Style: registration.TShirt.Style}
	}

	return returnInfo, nil
}
