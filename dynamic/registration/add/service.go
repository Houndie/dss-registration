package add

import (
	"context"
	"fmt"
	"net/http"

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

type MailClient interface {
	Send(email *mail.SGMailV3) (*rest.Response, error)
}

type Store interface {
	AddRegistration(context.Context, *StoreRegistration) (string, error)
	DeleteRegistration(context.Context, string) error
}

type SquareClient interface {
	ListCatalog(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator
	ListLocations(ctx context.Context) ([]*square.Location, error)
	CreateCheckout(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error)
	UpdateOrder(ctx context.Context, locationId, orderId string, order *square.Order, fieldsToClear []string, idempotencyKey string) (*square.Order, error)
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
}

func NewService(logger *logrus.Logger, store Store, client SquareClient, authorizer Authorizer, mailClient MailClient) *Service {
	return &Service{
		store:      store,
		logger:     logger,
		client:     client,
		authorizer: authorizer,
		mailClient: mailClient,
	}
}

func containsPaidItems(r *Registration) bool {
	_, noPassOk := r.PassType.(*common.NoPass)
	return !noPassOk || r.MixAndMatch != nil || r.TeamCompetition != nil || r.TShirt != nil || r.SoloJazz
}

func (s *Service) Add(ctx context.Context, registration *Registration, redirectUrl, accessToken string) (string, error) {
	s.logger.Trace("in add registration service")
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
	}
	returnerURL := redirectUrl
	if containsPaidItems(registration) {
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
		objects := s.client.ListCatalog(ctx, nil)

		idempotencyKey, err := uuid.NewV4()
		if err != nil {
			wrap := "error generating idempotency key"
			s.logger.WithError(err).Error(wrap)
			return "", errors.Wrap(err, wrap)
		}

		order := &square.CreateOrderRequest{
			IdempotencyKey: idempotencyKey.String(),
			Order: &square.Order{
				ReferenceId: referenceId.String(),
				LocationId:  locations[0].Id,
				LineItems:   []*square.OrderLineItem{},
				Version:     1,
			},
		}

		for objects.Next() {
			item, ok := objects.Value().CatalogObjectType.(*square.CatalogItem)
			if !ok {
				s.logger.Trace("Square object was not of type catalog item")
				continue
			}
			s.logger.Tracef("Comparing item name %s to legend", item.Name)
			switch item.Name {
			case utility.MixAndMatchItem, utility.TeamCompItem, utility.SoloJazzItem, utility.TShirtItem:
				if len(item.Variations) != 1 {
					err := fmt.Errorf("Found unexpected number of variations: %v", len(item.Variations))
					s.logger.Error(err)
					return "", err
				}
				v := item.Variations[0]
				_, ok := v.CatalogObjectType.(*square.CatalogItemVariation)
				if !ok {
					err := "Invalid response from square...item variation isn't a variation?"
					s.logger.Error(err)
					return "", errors.New(err)
				}
				switch item.Name {
				case utility.MixAndMatchItem:
					if registration.MixAndMatch == nil {
						continue
					}
					order.Order.LineItems = append(order.Order.LineItems, &square.OrderLineItem{
						Quantity:        "1",
						CatalogObjectId: v.Id,
					})
				case utility.TeamCompItem:
					if registration.TeamCompetition == nil {
						continue
					}
					order.Order.LineItems = append(order.Order.LineItems, &square.OrderLineItem{
						Quantity:        "1",
						CatalogObjectId: v.Id,
					})
				case utility.SoloJazzItem:
					if !registration.SoloJazz {
						continue
					}
					order.Order.LineItems = append(order.Order.LineItems, &square.OrderLineItem{
						Quantity:        "1",
						CatalogObjectId: v.Id,
					})
				case utility.TShirtItem:
					if registration.TShirt == nil {
						continue
					}
					order.Order.LineItems = append(order.Order.LineItems, &square.OrderLineItem{
						Quantity:        "1",
						CatalogObjectId: v.Id,
					})
				default:
					err := errors.New("Impossible code path...how did I get here")
					s.logger.Error(err)
					return "", err
				}
			case utility.DancePassItem:
				s.logger.Trace("Found dance pass item")
				if _, ok := registration.PassType.(*common.DanceOnlyPass); !ok {
					continue
				}
				for _, v := range item.Variations {
					variation, ok := v.CatalogObjectType.(*square.CatalogItemVariation)
					if !ok {
						err := "Invalid response from square...item variation isn't a variation?"
						s.logger.Error(err)
						return "", errors.New(err)
					}
					if variation.Name == "Presale" {
						s.logger.Trace("Found dance pass variant Presale")
						order.Order.LineItems = append(order.Order.LineItems, &square.OrderLineItem{
							Quantity:        "1",
							CatalogObjectId: v.Id,
						})
						break
					}
					s.logger.Tracef("Did not find dance pass variant Presale (found %s), moving on", variation.Name)
				}
			case utility.WeekendPassItem:
				s.logger.Trace("Found full weekend pass item")
				weekendPass, ok := registration.PassType.(*common.WeekendPass)
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
				for _, v := range item.Variations {
					variation, ok := v.CatalogObjectType.(*square.CatalogItemVariation)
					if !ok {
						err := "Invalid response from square...item variation isn't a variation?"
						s.logger.Error(err)
						return "", errors.New(err)
					}
					if variation.Name == tierString {
						s.logger.Trace("Found weekend pass")
						order.Order.LineItems = append(order.Order.LineItems, &square.OrderLineItem{
							Quantity:        "1",
							CatalogObjectId: v.Id,
						})
						break
					}
				}
			}
		}
		if objects.Error() != nil {
			wrap := "error fetching all items from square"
			utility.LogSquareError(s.logger, objects.Error(), wrap)
			return "", errors.Wrap(objects.Error(), wrap)
		}

		s.logger.Trace("creating checkout with square")
		checkout, err := s.client.CreateCheckout(ctx, locations[0].Id, idempotencyKey.String(), order, false, utility.SmackdownEmail, registration.Email, nil, redirectUrl, nil, "")
		if err != nil {
			wrap := "error creating square checkout"
			utility.LogSquareError(s.logger, err, wrap)
			return "", errors.Wrap(err, wrap)
		}

		storeRegistration.OrderIds = []string{checkout.Order.Id}
		returnerURL = checkout.CheckoutPageUrl
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
