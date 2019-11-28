package add

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/registration/common"
	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/utility"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sirupsen/logrus"
)

type tierData struct {
	tier int
	cost int
}

type MailClient interface {
	Send(email *mail.SGMailV3) (*rest.Response, error)
}

type Store interface {
	AddRegistration(context.Context, *StoreRegistration) (string, error)
	DeleteRegistration(context.Context, string) error
	GetDiscounts(context.Context, []string) ([]string, []*common.StoreDiscount, error)
}

type SquareClient interface {
	ListCatalog(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator
	ListLocations(ctx context.Context) ([]*square.Location, error)
	CreateCheckout(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error)
	UpdateOrder(ctx context.Context, locationId, orderId string, order *square.Order, fieldsToClear []string, idempotencyKey string) (*square.Order, error)
	BatchRetrieveInventoryCounts(ctx context.Context, catalogObjectIds, locationIds []string, updatedAfter *time.Time) square.BatchRetrieveInventoryCountsIterator
}

type Authorizer interface {
	Userinfo(ctx context.Context, accessToken string) (*authorizer.Userinfo, error)
}

type Service struct {
	client     SquareClient
	store      Store
	authorizer Authorizer
	logger     *logrus.Logger
	mailClient MailClient
	active     bool
}

func NewService(logger *logrus.Logger, store Store, client SquareClient, authorizer Authorizer, mailClient MailClient, active bool) *Service {
	return &Service{
		store:      store,
		logger:     logger,
		client:     client,
		authorizer: authorizer,
		mailClient: mailClient,
		active:     active,
	}
}

func containsPaidItems(r *Registration) bool {
	_, noPassOk := r.PassType.(*common.NoPass)
	return !noPassOk || r.MixAndMatch != nil || r.TeamCompetition != nil || r.TShirt != nil || r.SoloJazz
}

func (s *Service) Add(ctx context.Context, registration *Registration, redirectUrl, accessToken string) (string, error) {
	s.logger.Trace("in add registration service")
	if !s.active {
		s.logger.Error("registration found when service is not active")
		return "", errors.New("registration found when service is not active")
	}
	userid := ""
	if accessToken != "" {
		s.logger.Trace("found access token, calling userinfo endpoint")
		userinfo, err := s.authorizer.Userinfo(ctx, accessToken)
		if err != nil {
			msg := "error fetching userinfo"
			s.logger.WithError(err).Debug(msg)
			return "", errors.Wrap(err, msg)
		}
		userid = userinfo.UserId
	}
	s.logger.Trace("generating reference id")
	referenceId, err := uuid.NewV4()
	if err != nil {
		wrap := "error generating reference id"
		s.logger.WithError(err).Error(wrap)
		return "", errors.Wrap(err, wrap)
	}

	s.logger.Trace("looking up discount codes from database")
	discountKeys, discounts, err := s.store.GetDiscounts(ctx, registration.DiscountCodes)
	if err != nil {
		wrap := "error fetching discount codes from datastore"
		if _, ok := errors.Cause(err).(common.ErrDiscountDoesNotExist); ok {
			s.logger.WithError(err).Debug(wrap)
		} else {
			s.logger.WithError(err).Error(wrap)
		}
		return "", errors.Wrap(err, wrap)
	}

	storeRegistration := &StoreRegistration{
		FirstName:       registration.FirstName,
		LastName:        registration.LastName,
		StreetAddress:   registration.StreetAddress,
		City:            registration.City,
		State:           registration.State,
		ZipCode:         registration.ZipCode,
		Email:           registration.Email,
		HomeScene:       registration.HomeScene,
		IsStudent:       registration.IsStudent,
		PassType:        registration.PassType,
		MixAndMatch:     registration.MixAndMatch,
		SoloJazz:        registration.SoloJazz,
		TeamCompetition: registration.TeamCompetition,
		TShirt:          registration.TShirt,
		Housing:         registration.Housing,
		UserId:          userid,
		Discounts:       discountKeys,
	}
	returnerURL := redirectUrl
	if containsPaidItems(registration) {
		discountsMap := map[string]common.PurchaseItem{}
		for _, discount := range discounts {
			discountsMap[discount.Name] = discount.AppliedTo
		}

		purchaseItems := map[common.PurchaseItem]*square.OrderLineItem{}
		switch registration.PassType.(type) {
		case *common.WeekendPass:
			purchaseItems[common.FullWeekendPurchaseItem] = &square.OrderLineItem{Quantity: "1"}
		case *common.DanceOnlyPass:
			purchaseItems[common.DanceOnlyPurchaseItem] = &square.OrderLineItem{Quantity: "1"}
		default:
			//Do nothing
		}

		if registration.MixAndMatch != nil {
			purchaseItems[common.MixAndMatchPurchaseItem] = &square.OrderLineItem{Quantity: "1"}
		}

		if registration.SoloJazz {
			purchaseItems[common.SoloJazzPurchaseItem] = &square.OrderLineItem{Quantity: "1"}
		}

		if registration.TeamCompetition != nil {
			purchaseItems[common.TeamCompetitionPurchaseItem] = &square.OrderLineItem{Quantity: "1"}
		}

		if registration.TShirt != nil {
			purchaseItems[common.TShirtPurchaseItem] = &square.OrderLineItem{Quantity: "1"}
		}

		if registration.IsStudent {
			discountsMap[utility.StudentDiscountItem] = common.FullWeekendPurchaseItem
		}

		s.logger.Trace("Fetching all locations from square")
		locations, err := s.client.ListLocations(ctx)
		if err != nil {
			wrap := "error listing locations from square"
			utility.LogSquareError(s.logger, err, wrap)
			return "", errors.Wrap(err, wrap)
		}
		if len(locations) != 1 {
			err := fmt.Errorf("found wrong number of locations %v", len(locations))
			s.logger.Error(err)
			return "", err
		}

		s.logger.Trace("Fetching all items from square")
		objects := s.client.ListCatalog(ctx, []square.CatalogObjectType{square.CatalogObjectTypeItem, square.CatalogObjectTypeDiscount})

		idempotencyKey, err := uuid.NewV4()
		if err != nil {
			wrap := "error generating idempotency key"
			s.logger.WithError(err).Error(wrap)
			return "", errors.Wrap(err, wrap)
		}

		// We'll need to fetch all the weekend pass ids in case of out of stock
		tiers := map[string]tierData{}
		for objects.Next() {
			switch o := objects.Value().CatalogObjectType.(type) {
			case *square.CatalogItem:
				s.logger.Tracef("Comparing item name %s to legend", o.Name)
				switch o.Name {
				case utility.MixAndMatchItem, utility.TeamCompItem, utility.SoloJazzItem, utility.TShirtItem:
					if len(o.Variations) != 1 {
						err := fmt.Errorf("Found unexpected number of variations: %v", len(o.Variations))
						s.logger.Error(err)
						return "", err
					}
					v := o.Variations[0]
					_, ok := v.CatalogObjectType.(*square.CatalogItemVariation)
					if !ok {
						err := "Invalid response from square...item variation isn't a variation?"
						s.logger.Error(err)
						return "", errors.New(err)
					}
					var pi *square.OrderLineItem
					switch o.Name {
					case utility.MixAndMatchItem:
						pi, ok = purchaseItems[common.MixAndMatchPurchaseItem]
						if !ok {
							continue
						}
					case utility.TeamCompItem:
						pi, ok = purchaseItems[common.TeamCompetitionPurchaseItem]
						if !ok {
							continue
						}
					case utility.SoloJazzItem:
						pi, ok = purchaseItems[common.SoloJazzPurchaseItem]
						if !ok {
							continue
						}
					case utility.TShirtItem:
						pi, ok = purchaseItems[common.TShirtPurchaseItem]
						if !ok {
							continue
						}
					default:
						err := errors.New("Impossible code path...how did I get here")
						s.logger.Error(err)
						return "", err
					}
					pi.CatalogObjectId = v.Id
				case utility.DancePassItem:
					s.logger.Trace("Found dance pass item")
					pi, ok := purchaseItems[common.DanceOnlyPurchaseItem]
					if !ok {
						continue
					}
					for _, v := range o.Variations {
						variation, ok := v.CatalogObjectType.(*square.CatalogItemVariation)
						if !ok {
							err := "Invalid response from square...item variation isn't a variation?"
							s.logger.Error(err)
							return "", errors.New(err)
						}
						if variation.Name != "Presale" {
							s.logger.Tracef("Did not find dance pass variant Presale (found %s), moving on", variation.Name)
							continue
						}
						s.logger.Trace("Found dance pass variant Presale")
						pi.CatalogObjectId = v.Id
					}
				case utility.WeekendPassItem:
					s.logger.Trace("Found full weekend pass item")
					weekendPass, ok := registration.PassType.(*common.WeekendPass)
					if !ok {
						continue
					}
					pi, ok := purchaseItems[common.FullWeekendPurchaseItem]
					if !ok {
						continue
					}

					var tierString string
					switch weekendPass.Tier {
					case common.WeekendPassTier1:
						tierString = utility.WeekendPassTier1Name
					case common.WeekendPassTier2:
						tierString = utility.WeekendPassTier2Name
					case common.WeekendPassTier3:
						tierString = utility.WeekendPassTier3Name
					case common.WeekendPassTier4:
						tierString = utility.WeekendPassTier4Name
					case common.WeekendPassTier5:
						tierString = utility.WeekendPassTier5Name
					}
					for _, v := range o.Variations {
						variation, ok := v.CatalogObjectType.(*square.CatalogItemVariation)
						if !ok {
							err := "Invalid response from square...item variation isn't a variation?"
							s.logger.Error(err)
							return "", errors.New(err)
						}
						switch variation.Name {
						case utility.WeekendPassTier1Name:
							tiers[v.Id] = tierData{1, variation.PriceMoney.Amount}
						case utility.WeekendPassTier2Name:
							tiers[v.Id] = tierData{2, variation.PriceMoney.Amount}
						case utility.WeekendPassTier3Name:
							tiers[v.Id] = tierData{3, variation.PriceMoney.Amount}
						case utility.WeekendPassTier4Name:
							tiers[v.Id] = tierData{4, variation.PriceMoney.Amount}
						case utility.WeekendPassTier5Name:
							tiers[v.Id] = tierData{5, variation.PriceMoney.Amount}
						default: // Do nothing, we have other names that are allowable
						}
						if variation.Name == tierString {
							s.logger.Trace("Found weekend pass")
							pi.CatalogObjectId = v.Id
						}
					}
				}
			case *square.CatalogDiscount:
				appliedTo, ok := discountsMap[o.Name]
				if !ok {
					continue
				}
				delete(discountsMap, o.Name)
				pi, ok := purchaseItems[appliedTo]
				if !ok {
					continue
				}
				if pi.Discounts == nil {
					pi.Discounts = []*square.OrderLineItemDiscount{}
				}
				pi.Discounts = append(pi.Discounts, &square.OrderLineItemDiscount{
					CatalogObjectId: objects.Value().Id,
				})
			default:
				err := errors.New("found unknown catalog object type (should be impossible)")
				s.logger.Error(err)
				return "", err
			}
		}
		if objects.Error() != nil {
			wrap := "error fetching all items from square"
			utility.LogSquareError(s.logger, objects.Error(), wrap)
			return "", errors.Wrap(objects.Error(), wrap)
		}

		fullWeekendItem, ok := purchaseItems[common.FullWeekendPurchaseItem]
		if ok {
			// Check stock so that an open registration from day 1 doesn't still register for tier 1
			lowestTier := 999999
			lowestTierCost := 0
			outOfStock := false
			weekendPassIds := []string{}
			for weekendPassId, _ := range tiers {
				weekendPassIds = append(weekendPassIds, weekendPassId)
			}
			counts := s.client.BatchRetrieveInventoryCounts(ctx, weekendPassIds, nil, nil)
			for counts.Next() {
				quantity, err := strconv.ParseFloat(counts.Value().Quantity, 64)
				if err != nil {
					s.logger.WithField("quantity", counts.Value().Quantity).Error("could not convert quantity to float")
					return "", errors.Wrapf(err, "could not convert quantity %s to float", counts.Value().Quantity)
				}
				if counts.Value().CatalogObjectId == fullWeekendItem.CatalogObjectId {
					if quantity < 1 {
						outOfStock = true
					} else {
						break
					}
				}
				if quantity > 0 {
					thisTier := tiers[counts.Value().CatalogObjectId]
					s.logger.Tracef("tier %d", thisTier.tier)
					if thisTier.tier < lowestTier {
						s.logger.Tracef("new lowest tier found")
						lowestTier = thisTier.tier
						lowestTierCost = thisTier.cost
					}
				}
			}
			if outOfStock {
				return "", ErrOutOfStock{
					NextTier: lowestTier,
					NextCost: lowestTierCost,
				}
			}
		}

		if len(discountsMap) != 0 {
			keys := make([]string, len(discountsMap))
			i := 0
			for key, _ := range discountsMap {
				keys[i] = key
				i++
			}
			err := fmt.Errorf("disount names %v not found in square database", keys)
			s.logger.Error(err)
			return "", err
		}

		order := &square.CreateOrderRequest{
			IdempotencyKey: idempotencyKey.String(),
			Order: &square.Order{
				ReferenceId: referenceId.String(),
				LocationId:  locations[0].Id,
				LineItems:   make([]*square.OrderLineItem, len(purchaseItems)),
				//Version:     1,
			},
		}
		i := 0
		for itemName, purchaseItem := range purchaseItems {
			if purchaseItem.CatalogObjectId == "" {
				err := fmt.Errorf("could not find square item to purchase a %s", itemName)
				s.logger.Error(err)
				return "", err
			}

			order.Order.LineItems[i] = purchaseItem
			i++
		}

		s.logger.Trace("creating checkout with square")
		checkout, err := s.client.CreateCheckout(ctx, locations[0].Id, idempotencyKey.String(), order, false, utility.SmackdownEmail, registration.Email, nil, redirectUrl, nil, "")
		if err != nil {
			errorList, ok := err.(*square.ErrorList)

			// If this error is anything other than "can't create checkouts worth less than a dollar"
			if !ok || len(errorList.Errors) > 1 || errorList.Errors[0].Category != square.ErrorCategoryInvalidRequestError || errorList.Errors[0].Code != square.ErrorCodeValueTooLow || errorList.Errors[0].Field != "order.total_money.amount" {
				wrap := "error creating square checkout"
				utility.LogSquareError(s.logger, err, wrap)
				return "", errors.Wrap(err, wrap)
			}
			s.logger.Trace("registration with checkout amount less than a dollar found, not creating order or checkout")
		} else {
			storeRegistration.OrderIds = []string{checkout.Order.Id}
			returnerURL = checkout.CheckoutPageUrl
		}
	}

	s.logger.Trace("Adding registration to database")
	registrationId, err := s.store.AddRegistration(ctx, storeRegistration)
	if err != nil {
		wrap := "error adding registration to database"
		s.logger.WithError(err).Error(wrap)
		return "", errors.Wrap(err, wrap)
	}

	s.logger.Trace("sending registration email")
	from := mail.NewEmail("Dayton Swing Smackdown", "info@daytonswingsmackdown.com")
	to := mail.NewEmail(registration.FirstName+" "+registration.LastName, registration.Email)
	personalization := &mail.Personalization{
		DynamicTemplateData: map[string]interface{}{
			"first_name":       registration.FirstName,
			"last_name":        registration.LastName,
			"street_address":   registration.StreetAddress,
			"city":             registration.City,
			"state":            registration.State,
			"zip_code":         registration.ZipCode,
			"email":            registration.Email,
			"home_scene":       registration.HomeScene,
			"student":          registration.IsStudent,
			"registration_id":  registrationId,
			"mix_and_match":    registration.MixAndMatch != nil,
			"solo_jazz":        registration.SoloJazz,
			"team_competition": registration.TeamCompetition != nil,
			"tshirt":           registration.TShirt != nil,
		},
		To: []*mail.Email{to},
	}

	if registration.MixAndMatch != nil {
		personalization.SetDynamicTemplateData("mix_and_match_role", registration.MixAndMatch.Role)
	}
	if registration.TeamCompetition != nil {
		personalization.SetDynamicTemplateData("team_competition_name", registration.TeamCompetition.Name)
	}
	if registration.TShirt != nil {
		personalization.SetDynamicTemplateData("tshirt_style", registration.TShirt.Style)
	}
	switch p := registration.PassType.(type) {
	case *common.WeekendPass:
		personalization.SetDynamicTemplateData("weekend_pass_type", "Full Weekend Pass")
		personalization.SetDynamicTemplateData("workshop_level", p.Level)
	case *common.DanceOnlyPass:
		personalization.SetDynamicTemplateData("weekend_pass_type", "Dance Only Pass")
	default:
		personalization.SetDynamicTemplateData("weekend_pass_type", false)
	}

	switch h := registration.Housing.(type) {
	case *common.ProvideHousing:
		personalization.SetDynamicTemplateData("housing", "I can provide housing")
		personalization.SetDynamicTemplateData("provide_housing", map[string]interface{}{
			"guests":  h.Quantity,
			"pets":    h.Pets,
			"details": h.Details,
		})
	case *common.RequireHousing:
		personalization.SetDynamicTemplateData("housing", "I require housing")
		personalization.SetDynamicTemplateData("require_housing", map[string]interface{}{
			"pet_allergies": h.PetAllergies,
			"details":       h.Details,
		})
	default:
		personalization.SetDynamicTemplateData("housing", "I neither require nor can provide housing")
	}
	message := &mail.SGMailV3{
		From:             from,
		Personalizations: []*mail.Personalization{personalization},
		TemplateID:       "d-15759d9e2e3d4dfa9602dc7ec6512e8c",
	}
	mailResp, err := s.mailClient.Send(message)
	if err != nil {
		wrap := "error sending registration email"
		s.logger.WithError(err).Error(wrap)
		return "", errors.Wrap(err, wrap)
	}
	if mailResp.StatusCode != http.StatusAccepted {
		err := fmt.Errorf("received bad status code %v", mailResp.StatusCode)
		s.logger.WithField("sendgrid response", mailResp).Error(err)
		return "", err
	}
	return returnerURL, nil
}
