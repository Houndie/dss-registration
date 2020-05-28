package registration

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/gofrs/uuid"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func containsPaidItems(r *Info) bool {
	_, noPassOk := r.PassType.(*NoPass)
	return !noPassOk || r.MixAndMatch != nil || r.TeamCompetition != nil || r.TShirt != nil || r.SoloJazz != nil
}

func makeLineItems(registration *Info, squareData *common.SquareData, paymentData *common.PaymentData, discounts map[storage.PurchaseItem][]string) ([]*square.OrderLineItem, error) {
	lineItems := []*square.OrderLineItem{}
	switch t := registration.PassType.(type) {
	case *WeekendPass:
		if !paymentData.WeekendPassPaid {
			li, err := makeLineItem(squareData.FullWeekend[t.Tier].VariationID, discounts[storage.FullWeekendPurchaseItem], squareData.Discounts)
			if err != nil {
				return nil, fmt.Errorf("error making full weekend line item: %w", err)
			}

			if registration.IsStudent {
				studentDiscount := &square.OrderLineItemDiscount{
					CatalogObjectId: squareData.StudentDiscount.ID,
				}

				if li.Discounts == nil {
					li.Discounts = []*square.OrderLineItemDiscount{studentDiscount}
				} else {
					li.Discounts = append(li.Discounts, studentDiscount)
				}
			}
			lineItems = append(lineItems, li)
		}
	case *DanceOnlyPass:
		if !paymentData.DanceOnlyPaid {
			li, err := makeLineItem(squareData.DanceOnly.VariationID, discounts[storage.DanceOnlyPurchaseItem], squareData.Discounts)
			if err != nil {
				return nil, fmt.Errorf("error making dance only item: %w", err)
			}
			lineItems = append(lineItems, li)
		}
	}

	if registration.MixAndMatch != nil && !paymentData.MixAndMatchPaid {
		li, err := makeLineItem(squareData.MixAndMatch.VariationID, discounts[storage.MixAndMatchPurchaseItem], squareData.Discounts)
		if err != nil {
			return nil, fmt.Errorf("error making mix and match line item: %w", err)
		}
		lineItems = append(lineItems, li)
	}

	if registration.SoloJazz != nil && !paymentData.SoloJazzPaid {
		li, err := makeLineItem(squareData.SoloJazz.VariationID, discounts[storage.SoloJazzPurchaseItem], squareData.Discounts)
		if err != nil {
			return nil, fmt.Errorf("error making solo jazz line item: %w", err)
		}
		lineItems = append(lineItems, li)
	}

	if registration.TeamCompetition != nil && !paymentData.TeamCompetitionPaid {
		li, err := makeLineItem(squareData.TeamCompetition.VariationID, discounts[storage.TeamCompetitionPurchaseItem], squareData.Discounts)
		if err != nil {
			return nil, fmt.Errorf("error making team competition line item: %w", err)
		}
		lineItems = append(lineItems, li)
	}

	if registration.TShirt != nil && !paymentData.TShirtPaid {
		li, err := makeLineItem(squareData.TShirt.VariationID, discounts[storage.TShirtPurchaseItem], squareData.Discounts)
		if err != nil {
			return nil, fmt.Errorf("error making t-shirt line item: %w", err)
		}
		lineItems = append(lineItems, li)
	}
	return lineItems, nil
}

func makeLineItem(catalogID string, discountNames []string, discounts map[string]*common.Discount) (*square.OrderLineItem, error) {
	var orderDiscounts []*square.OrderLineItemDiscount
	if len(discountNames) != 0 {
		orderDiscounts = make([]*square.OrderLineItemDiscount, len(discountNames))
		for i, name := range discountNames {
			d, ok := discounts[name]
			if !ok {
				return nil, fmt.Errorf("discount name %v not found in square data", name)
			}

			orderDiscounts[i] = &square.OrderLineItemDiscount{
				CatalogObjectId: d.ID,
			}
		}
	}
	return &square.OrderLineItem{
		Quantity:        "1",
		CatalogObjectId: catalogID,
		Discounts:       orderDiscounts,
	}, nil
}

func discountCodeMap(ctx context.Context, store Store, discountCodes []string) (map[storage.PurchaseItem][]string, error) {
	discounts := map[storage.PurchaseItem][]string{}
	for _, d := range discountCodes {
		storedDiscount, err := store.GetDiscount(ctx, d)
		if err != nil {
			return nil, fmt.Errorf("error fetching discount from store: %w", err)
		}
		for _, sd := range storedDiscount.Discounts {
			if discounts[sd.AppliedTo] == nil {
				discounts[sd.AppliedTo] = []string{}
			}
			discounts[sd.AppliedTo] = append(discounts[sd.AppliedTo], sd.Name)
		}
	}
	return discounts, nil
}

func (s *Service) Add(ctx context.Context, registration *Info, redirectUrl, idempotencyKey, accessToken string) (string, error) {
	s.logger.Trace("in add registration service")
	if !s.active {
		return "", ErrRegistrationDisabled
	}

	returnerURL := redirectUrl
	orderID := ""
	if containsPaidItems(registration) {

		s.logger.Trace("generating reference id")
		referenceId, err := uuid.NewV4()
		if err != nil {
			return "", fmt.Errorf("error generating reference id: %w", err)
		}

		discounts, err := discountCodeMap(ctx, s.store, registration.DiscountCodes)
		if err != nil {
			return "", err
		}

		s.logger.Trace("Fetching all locations from square")
		locations, err := s.client.ListLocations(ctx)
		if err != nil {
			return "", fmt.Errorf("error listing locations from square: %w", err)
		}
		if len(locations) != 1 {
			return "", fmt.Errorf("found wrong number of locations %v", len(locations))
		}

		s.logger.Trace("Fetching all items from square")
		squareData, err := common.GetSquareCatalog(ctx, s.client)
		if err != nil {
			return "", fmt.Errorf("error fetching all items from square: %w", err)
		}

		myFullWeekend, ok := registration.PassType.(*WeekendPass)
		if ok {
			bestTier, bestCost, err := common.LowestInStockTier(ctx, squareData, s.client)
			if err != nil {
				return "", fmt.Errorf("error finding best tier and cost: %w", err)
			}
			if myFullWeekend.Tier < bestTier {
				return "", ErrOutOfStock{
					NextTier: bestTier,
					NextCost: bestCost,
				}
			}
		}

		lineItems, err := makeLineItems(registration, squareData, &common.PaymentData{}, discounts)
		if err != nil {
			return "", err
		}

		order := &square.CreateOrderRequest{
			IdempotencyKey: idempotencyKey,
			Order: &square.Order{
				ReferenceId: referenceId.String(),
				LocationId:  locations[0].Id,
				LineItems:   lineItems,
			},
		}

		s.logger.Trace("creating checkout with square")
		returnerURL, orderID, err = common.CreateCheckout(ctx, s.client, locations[0].Id, idempotencyKey, order, registration.Email, redirectUrl)
		if err != nil {
			return "", err
		}
	}

	s.logger.Trace("Adding registration to database")
	userid := ""
	if accessToken != "" {
		s.logger.Trace("found access token, calling userinfo endpoint")
		userinfo, err := s.authorizer.Userinfo(ctx, accessToken)
		if err != nil {
			return "", fmt.Errorf("error fetching userinfo: %w", err)
		}
		userid = userinfo.UserId
	}

	var orderIDs []string
	if orderID != "" {
		orderIDs = []string{orderID}
	}

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
		UserId:          userid,
		DiscountCodes:   registration.DiscountCodes,
		OrderIds:        orderIDs,
	}
	registrationId, err := s.store.AddRegistration(ctx, storeRegistration)
	if err != nil {
		return "", fmt.Errorf("error adding registration to database: %w", err)
	}

	s.logger.Trace("sending registration email")
	from := mail.NewEmail("Dayton Swing Smackdown", "info@daytonswingsmackdown.com")
	to := mail.NewEmail(registration.FirstName+" "+registration.LastName, registration.Email)
	personalization := &mail.Personalization{
		DynamicTemplateData: map[string]interface{}{
			mailFirstNameKey:      registration.FirstName,
			mailLastNameKey:       registration.LastName,
			mailStreetAddressKey:  registration.StreetAddress,
			mailCityKey:           registration.City,
			mailStateKey:          registration.State,
			mailZipCodeKey:        registration.ZipCode,
			mailEmailKey:          registration.Email,
			mailHomeSceneKey:      registration.HomeScene,
			mailStudentKey:        registration.IsStudent,
			mailRegistrationIDKey: registrationId,
			mailMixAndMatchKey:    registration.MixAndMatch != nil,
			mailSoloJazzKey:       registration.SoloJazz,
			mailTeamCompKey:       registration.TeamCompetition != nil,
			mailTShirtKey:         registration.TShirt != nil,
		},
		To: []*mail.Email{to},
	}

	if registration.MixAndMatch != nil {
		personalization.SetDynamicTemplateData(mailMixAndMatchRoleKey, registration.MixAndMatch.Role)
	}
	if registration.TeamCompetition != nil {
		personalization.SetDynamicTemplateData(mailTeamCompNameKey, registration.TeamCompetition.Name)
	}
	if registration.TShirt != nil {
		personalization.SetDynamicTemplateData(mailTShirtStyleKey, registration.TShirt.Style)
	}
	switch p := registration.PassType.(type) {
	case *WeekendPass:
		personalization.SetDynamicTemplateData(mailWeekendPassKey, mailFullWeekendValue)
		personalization.SetDynamicTemplateData(mailWorkshopLevelKey, p.Level)
	case *DanceOnlyPass:
		personalization.SetDynamicTemplateData(mailWeekendPassKey, mailDanceOnlyValue)
	default:
		personalization.SetDynamicTemplateData(mailWeekendPassKey, mailNoPassValue)
	}

	switch h := registration.Housing.(type) {
	case *storage.ProvideHousing:
		personalization.SetDynamicTemplateData(mailHousingKey, mailProvideHousingValue)
		personalization.SetDynamicTemplateData(mailProvideHousingKey, map[string]interface{}{
			mailProvideHousingGuestsKey:  h.Quantity,
			mailProvideHousingPetsKey:    h.Pets,
			mailProvideHousingDetailsKey: h.Details,
		})
	case *storage.RequireHousing:
		personalization.SetDynamicTemplateData(mailHousingKey, mailRequireHousingValue)
		personalization.SetDynamicTemplateData(mailRequireHousingKey, map[string]interface{}{
			mailRequireHousingAllergiesKey: h.PetAllergies,
			mailRequireHousingDetailsKey:   h.Details,
		})
	default:
		personalization.SetDynamicTemplateData(mailHousingKey, mailNoHousingValue)
	}

	var mailSettings *mail.MailSettings
	if s.useMailSandbox {
		mailSettings = &mail.MailSettings{
			SandboxMode: mail.NewSetting(true),
		}

	}
	message := &mail.SGMailV3{
		From:             from,
		Personalizations: []*mail.Personalization{personalization},
		TemplateID:       "d-15759d9e2e3d4dfa9602dc7ec6512e8c",
		MailSettings:     mailSettings,
	}
	mailResp, err := s.mailClient.Send(message)
	if err != nil {
		return "", fmt.Errorf("error sending registration email: %w", err)
	}
	okCode := http.StatusAccepted
	if s.useMailSandbox {
		okCode = http.StatusOK
	}
	if mailResp.StatusCode != okCode {
		return "", fmt.Errorf("received bad status code from mailserver %v", mailResp.StatusCode)
	}
	return returnerURL, nil
}
