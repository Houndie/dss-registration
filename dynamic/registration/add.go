package registration

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/sendinblue"
	"github.com/Houndie/dss-registration/dynamic/storage"
)

func paymentCheck(registration *Info, isAdmin bool, storeRegistration *storage.Registration, pd *common.PaymentData) error {
	switch p := registration.PassType.(type) {
	case *WeekendPass:
		if storePass, ok := storeRegistration.PassType.(*storage.WeekendPass); p.AdminPaymentOverride && !isAdmin && (!ok || !storePass.ManuallyPaid) {
			return ErrHasAdminOverride{
				"weekend_pass",
			}
		}

		if p.SquarePaid && !pd.WeekendPassPaid {
			return ErrHasSquarePayment{
				"weekend_pass",
			}
		}

	case *DanceOnlyPass:
		if storePass, ok := storeRegistration.PassType.(*storage.DanceOnlyPass); p.AdminPaymentOverride && !isAdmin && (!ok || !storePass.ManuallyPaid) {
			return ErrHasAdminOverride{
				"dance_only_pass",
			}
		}

		if p.SquarePaid && !pd.DanceOnlyPaid {
			return ErrHasSquarePayment{
				"dance_only_pass",
			}
		}
	}

	if registration.MixAndMatch != nil {
		if registration.MixAndMatch.AdminPaymentOverride && !isAdmin && (storeRegistration.MixAndMatch == nil || !storeRegistration.MixAndMatch.ManuallyPaid) {
			return ErrHasAdminOverride{
				"mix_and_match",
			}
		}

		if registration.MixAndMatch.SquarePaid && !pd.MixAndMatchPaid {
			return ErrHasSquarePayment{
				"mix_and_match",
			}
		}
	}

	if registration.SoloJazz != nil {
		if registration.SoloJazz.AdminPaymentOverride && !isAdmin && (storeRegistration.SoloJazz == nil || !storeRegistration.SoloJazz.ManuallyPaid) {
			return ErrHasAdminOverride{
				"solo_jazz",
			}
		}

		if registration.SoloJazz.SquarePaid && !pd.SoloJazzPaid {
			return ErrHasSquarePayment{
				"solo_jazz",
			}
		}
	}

	if registration.TeamCompetition != nil {
		if registration.TeamCompetition.AdminPaymentOverride && !isAdmin && (storeRegistration.TeamCompetition == nil || !storeRegistration.TeamCompetition.ManuallyPaid) {
			return ErrHasAdminOverride{
				"team_competition",
			}
		}

		if registration.TeamCompetition.SquarePaid && !pd.TeamCompetitionPaid {
			return ErrHasSquarePayment{
				"team_competition",
			}
		}
	}

	if registration.TShirt != nil {
		if registration.TShirt.AdminPaymentOverride && !isAdmin && (storeRegistration.TShirt == nil || !storeRegistration.TShirt.ManuallyPaid) {
			return ErrHasAdminOverride{
				"tshirt",
			}
		}

		if registration.TShirt.SquarePaid && !pd.TShirtPaid {
			return ErrHasSquarePayment{
				"tshirt",
			}
		}
	}

	return nil
}

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

	if err := paymentCheck(registration, false, &storage.Registration{}, &common.PaymentData{}); err != nil {
		return nil, err
	}

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
		/*BCC: []*sendinblue.EmailPerson{
			{
				Name:  "Dayton Swing Smackdown",
				Email: "info@daytonswingsmackdown.com",
			},
		},*/
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
