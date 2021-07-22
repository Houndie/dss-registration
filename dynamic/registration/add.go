package registration

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/sendinblue"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/square-go/objects"
	"github.com/gofrs/uuid"
)

func containsPaidItems(r *Info) bool {
	_, noPassOk := r.PassType.(*NoPass)
	return !noPassOk || r.MixAndMatch != nil || r.TeamCompetition != nil || r.TShirt != nil || r.SoloJazz != nil
}

func makeLineItems(registration *Info, squareData *common.SquareData, paymentData *common.PaymentData, discounts map[storage.PurchaseItem][]string) ([]*objects.OrderLineItem, []*objects.OrderLineItemDiscount, error) {
	lineItems := []*objects.OrderLineItem{}
	lineDiscounts := []*objects.OrderLineItemDiscount{}
	switch t := registration.PassType.(type) {
	case *WeekendPass:
		if !paymentData.WeekendPassPaid {
			li, ld, err := makeLineItem(squareData.PurchaseItems.FullWeekend[t.Tier].ID, discounts[storage.FullWeekendPurchaseItem])
			if err != nil {
				return nil, nil, fmt.Errorf("error making full weekend line item: %w", err)
			}

			if registration.IsStudent {
				uid, err := uuid.NewV4()
				if err != nil {
					return nil, nil, fmt.Errorf("error creating new uuid for student discount uid: %w", err)
				}
				studentDiscount := &objects.OrderLineItemDiscount{
					CatalogObjectID: squareData.Discounts.StudentDiscount.ID,
					Scope:           objects.OrderLineItemDiscountScopeLineItem,
					UID:             uid.String(),
				}
				studentAppliedDiscount := &objects.OrderLineItemAppliedDiscount{
					DiscountUID: uid.String(),
				}

				if li.AppliedDiscounts == nil {
					li.AppliedDiscounts = []*objects.OrderLineItemAppliedDiscount{studentAppliedDiscount}
				} else {
					li.AppliedDiscounts = append(li.AppliedDiscounts, studentAppliedDiscount)
				}
				lineDiscounts = append(lineDiscounts, studentDiscount)
			}
			lineItems = append(lineItems, li)
			lineDiscounts = append(lineDiscounts, ld...)
		}
	case *DanceOnlyPass:
		if !paymentData.DanceOnlyPaid {
			li, ld, err := makeLineItem(squareData.PurchaseItems.DanceOnly.ID, discounts[storage.DanceOnlyPurchaseItem])
			if err != nil {
				return nil, nil, fmt.Errorf("error making dance only item: %w", err)
			}
			lineItems = append(lineItems, li)
			lineDiscounts = append(lineDiscounts, ld...)
		}
	}

	if registration.MixAndMatch != nil && !paymentData.MixAndMatchPaid {
		li, ld, err := makeLineItem(squareData.PurchaseItems.MixAndMatch[registration.MixAndMatch.Role].ID, discounts[storage.MixAndMatchPurchaseItem])
		if err != nil {
			return nil, nil, fmt.Errorf("error making mix and match line item: %w", err)
		}
		lineItems = append(lineItems, li)
		lineDiscounts = append(lineDiscounts, ld...)
	}

	if registration.SoloJazz != nil && !paymentData.SoloJazzPaid {
		li, ld, err := makeLineItem(squareData.PurchaseItems.SoloJazz.ID, discounts[storage.SoloJazzPurchaseItem])
		if err != nil {
			return nil, nil, fmt.Errorf("error making solo jazz line item: %w", err)
		}
		lineItems = append(lineItems, li)
		lineDiscounts = append(lineDiscounts, ld...)
	}

	if registration.TeamCompetition != nil && !paymentData.TeamCompetitionPaid {
		li, ld, err := makeLineItem(squareData.PurchaseItems.TeamCompetition.ID, discounts[storage.TeamCompetitionPurchaseItem])
		if err != nil {
			return nil, nil, fmt.Errorf("error making team competition line item: %w", err)
		}
		lineItems = append(lineItems, li)
		lineDiscounts = append(lineDiscounts, ld...)
	}

	if registration.TShirt != nil && !paymentData.TShirtPaid {
		li, ld, err := makeLineItem(squareData.PurchaseItems.TShirt[registration.TShirt.Style].ID, discounts[storage.TShirtPurchaseItem])
		if err != nil {
			return nil, nil, fmt.Errorf("error making t-shirt line item: %w", err)
		}
		lineItems = append(lineItems, li)
		lineDiscounts = append(lineDiscounts, ld...)
	}
	return lineItems, lineDiscounts, nil
}

func makeLineItem(catalogID string, discountIDs []string) (*objects.OrderLineItem, []*objects.OrderLineItemDiscount, error) {
	var orderDiscounts []*objects.OrderLineItemDiscount
	var appliedDiscounts []*objects.OrderLineItemAppliedDiscount
	if len(discountIDs) != 0 {
		orderDiscounts = make([]*objects.OrderLineItemDiscount, len(discountIDs))
		appliedDiscounts = make([]*objects.OrderLineItemAppliedDiscount, len(discountIDs))
		for i, d := range discountIDs {
			uid, err := uuid.NewV4()
			if err != nil {
				return nil, nil, fmt.Errorf("error creating uid for line item discount: %w", err)
			}

			orderDiscounts[i] = &objects.OrderLineItemDiscount{
				CatalogObjectID: d,
				UID:             uid.String(),
			}

			appliedDiscounts[i] = &objects.OrderLineItemAppliedDiscount{
				DiscountUID: uid.String(),
			}
		}
	}
	return &objects.OrderLineItem{
		Quantity:         "1",
		CatalogObjectID:  catalogID,
		AppliedDiscounts: appliedDiscounts,
	}, orderDiscounts, nil
}

func discountCodeMap(ctx context.Context, codeDiscounts map[string][]*common.Discount, discountCodes []string) (map[storage.PurchaseItem][]string, error) {
	discounts := map[storage.PurchaseItem][]string{}
	for _, code := range discountCodes {
		for _, d := range codeDiscounts[code] {
			if discounts[d.AppliedTo] == nil {
				discounts[d.AppliedTo] = []string{}
			}
			discounts[d.AppliedTo] = append(discounts[d.AppliedTo], d.ID)
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
		referenceID, err := uuid.NewV4()
		if err != nil {
			return "", fmt.Errorf("error generating reference id: %w", err)
		}

		discounts, err := discountCodeMap(ctx, s.squareData.Discounts.CodeDiscounts, registration.DiscountCodes)
		if err != nil {
			return "", err
		}

		s.logger.Trace("Fetching all locations from square")
		locationListRes, err := s.client.Locations.List(ctx, nil)
		if err != nil {
			return "", fmt.Errorf("error listing locations from square: %w", err)
		}
		if len(locationListRes.Locations) != 1 {
			return "", fmt.Errorf("found wrong number of locations %v", len(locationListRes.Locations))
		}

		myFullWeekend, ok := registration.PassType.(*WeekendPass)
		if ok {
			bestTier, _, err := common.LowestInStockTier(ctx, s.squareData.PurchaseItems.FullWeekend, s.client)
			if err != nil {
				return "", fmt.Errorf("error finding best tier and cost: %w", err)
			}
			if myFullWeekend.Tier < bestTier {
				return "", ErrOutOfStock{
					NextTier: bestTier,
				}
			}
		}

		lineItems, lineDiscounts, err := makeLineItems(registration, s.squareData, &common.PaymentData{}, discounts)
		if err != nil {
			return "", err
		}

		order := &objects.CreateOrderRequest{
			IdempotencyKey: idempotencyKey,
			Order: &objects.Order{
				ReferenceID: referenceID.String(),
				LocationID:  locationListRes.Locations[0].ID,
				LineItems:   lineItems,
				Discounts:   lineDiscounts,
			},
		}

		s.logger.Trace("creating checkout with square")
		returnerURL, orderID, err = common.CreateCheckout(ctx, s.client, locationListRes.Locations[0].ID, idempotencyKey, order, registration.Email, redirectUrl)
		if err != nil {
			return "", err
		}
	}

	s.logger.Trace("Adding registration to database")
	userid := ""
	if accessToken != "" {
		s.logger.Trace("found access token")
		userinfo, err := s.authorizer.GetUserinfo(ctx, accessToken)
		if err != nil {
			return "", fmt.Errorf("error fetching userinfo: %w", err)
		}
		userid = userinfo.UserID()
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
		UserID:          userid,
		DiscountCodes:   registration.DiscountCodes,
		OrderIDs:        orderIDs,
	}
	_, err := s.store.AddRegistration(ctx, storeRegistration)
	if err != nil {
		return "", fmt.Errorf("error adding registration to database: %w", err)
	}

	s.logger.Trace("sending registration email")
	mailParams, err := toMailParams(registration)
	if err != nil {
		return "", fmt.Errorf("error generating mail parameters")
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
		return "", fmt.Errorf("error sending registration email: %w", err)
	}
	return returnerURL, nil
}
