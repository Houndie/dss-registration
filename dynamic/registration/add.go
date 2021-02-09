package registration

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/common"
	"github.com/Houndie/dss-registration/dynamic/sendinblue"
	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/gofrs/uuid"
)

func containsPaidItems(r *Info) bool {
	_, noPassOk := r.PassType.(*NoPass)
	return !noPassOk || r.MixAndMatch != nil || r.TeamCompetition != nil || r.TShirt != nil || r.SoloJazz != nil
}

func makeLineItems(registration *Info, squareData *common.SquareData, paymentData *common.PaymentData, discounts map[storage.PurchaseItem][]string) ([]*square.OrderLineItem, []*square.OrderLineItemDiscount, error) {
	lineItems := []*square.OrderLineItem{}
	lineDiscounts := []*square.OrderLineItemDiscount{}
	switch t := registration.PassType.(type) {
	case *WeekendPass:
		if !paymentData.WeekendPassPaid {
			li, ld, err := makeLineItem(squareData.FullWeekend[t.Tier].VariationID, discounts[storage.FullWeekendPurchaseItem], squareData.Discounts)
			if err != nil {
				return nil, nil, fmt.Errorf("error making full weekend line item: %w", err)
			}

			if registration.IsStudent {
				uid, err := uuid.NewV4()
				if err != nil {
					return nil, nil, fmt.Errorf("error creating new uuid for student discount uid: %w", err)
				}
				studentDiscount := &square.OrderLineItemDiscount{
					CatalogObjectID: squareData.StudentDiscount.ID,
					Scope:           square.OrderLineItemDiscountScopeLineItem,
					UID:             uid.String(),
				}
				studentAppliedDiscount := &square.OrderLineItemAppliedDiscount{
					DiscountUID: uid.String(),
				}

				if li.AppliedDiscounts == nil {
					li.AppliedDiscounts = []*square.OrderLineItemAppliedDiscount{studentAppliedDiscount}
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
			li, ld, err := makeLineItem(squareData.DanceOnly.VariationID, discounts[storage.DanceOnlyPurchaseItem], squareData.Discounts)
			if err != nil {
				return nil, nil, fmt.Errorf("error making dance only item: %w", err)
			}
			lineItems = append(lineItems, li)
			lineDiscounts = append(lineDiscounts, ld...)
		}
	}

	if registration.MixAndMatch != nil && !paymentData.MixAndMatchPaid {
		li, ld, err := makeLineItem(squareData.MixAndMatch.VariationID, discounts[storage.MixAndMatchPurchaseItem], squareData.Discounts)
		if err != nil {
			return nil, nil, fmt.Errorf("error making mix and match line item: %w", err)
		}
		lineItems = append(lineItems, li)
		lineDiscounts = append(lineDiscounts, ld...)
	}

	if registration.SoloJazz != nil && !paymentData.SoloJazzPaid {
		li, ld, err := makeLineItem(squareData.SoloJazz.VariationID, discounts[storage.SoloJazzPurchaseItem], squareData.Discounts)
		if err != nil {
			return nil, nil, fmt.Errorf("error making solo jazz line item: %w", err)
		}
		lineItems = append(lineItems, li)
		lineDiscounts = append(lineDiscounts, ld...)
	}

	if registration.TeamCompetition != nil && !paymentData.TeamCompetitionPaid {
		li, ld, err := makeLineItem(squareData.TeamCompetition.VariationID, discounts[storage.TeamCompetitionPurchaseItem], squareData.Discounts)
		if err != nil {
			return nil, nil, fmt.Errorf("error making team competition line item: %w", err)
		}
		lineItems = append(lineItems, li)
		lineDiscounts = append(lineDiscounts, ld...)
	}

	if registration.TShirt != nil && !paymentData.TShirtPaid {
		li, ld, err := makeLineItem(squareData.TShirt.VariationID, discounts[storage.TShirtPurchaseItem], squareData.Discounts)
		if err != nil {
			return nil, nil, fmt.Errorf("error making t-shirt line item: %w", err)
		}
		lineItems = append(lineItems, li)
		lineDiscounts = append(lineDiscounts, ld...)
	}
	return lineItems, lineDiscounts, nil
}

func makeLineItem(catalogID string, discountNames []string, discounts map[string]*common.Discount) (*square.OrderLineItem, []*square.OrderLineItemDiscount, error) {
	var orderDiscounts []*square.OrderLineItemDiscount
	var appliedDiscounts []*square.OrderLineItemAppliedDiscount
	if len(discountNames) != 0 {
		orderDiscounts = make([]*square.OrderLineItemDiscount, len(discountNames))
		appliedDiscounts = make([]*square.OrderLineItemAppliedDiscount, len(discountNames))
		for i, name := range discountNames {
			d, ok := discounts[name]
			if !ok {
				return nil, nil, fmt.Errorf("discount name %v not found in square data", name)
			}

			uid, err := uuid.NewV4()
			if err != nil {
				return nil, nil, fmt.Errorf("error creating uid for line item discount: %w", err)
			}

			orderDiscounts[i] = &square.OrderLineItemDiscount{
				CatalogObjectID: d.ID,
				UID:             uid.String(),
			}

			appliedDiscounts[i] = &square.OrderLineItemAppliedDiscount{
				DiscountUID: uid.String(),
			}
		}
	}
	return &square.OrderLineItem{
		Quantity:         "1",
		CatalogObjectID:  catalogID,
		AppliedDiscounts: appliedDiscounts,
	}, orderDiscounts, nil
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
	fmt.Println(redirectUrl)
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

		lineItems, lineDiscounts, err := makeLineItems(registration, squareData, &common.PaymentData{}, discounts)
		if err != nil {
			return "", err
		}

		order := &square.CreateOrderRequest{
			IdempotencyKey: idempotencyKey,
			Order: &square.Order{
				ReferenceID: referenceID.String(),
				LocationID:  locations[0].ID,
				LineItems:   lineItems,
				Discounts:   lineDiscounts,
			},
		}

		s.logger.Trace("creating checkout with square")
		returnerURL, orderID, err = common.CreateCheckout(ctx, s.client, locations[0].ID, idempotencyKey, order, registration.Email, redirectUrl)
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
		userid = userinfo.UserID
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
