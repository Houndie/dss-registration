package registration

import (
	"context"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/commontest"
	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/test_utility"
	"github.com/Houndie/dss-registration/dynamic/utility"
	"github.com/gofrs/uuid"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/sirupsen/logrus"
)

type itemCheck struct {
	found bool
	id    string
}

func discountCheck(t *testing.T, discountArray []*square.OrderLineItemDiscount, discountID string) {
	found := false
	for _, d := range discountArray {
		if d.CatalogObjectId == discountID {
			found = true
		}
	}
	if !found {
		t.Fatalf("looking for discount code %q, not found in registration", discountID)
	}

}

func TestAdd(t *testing.T) {
	expectedCheckoutUrl := "https://squareup.com/some_checkout"
	expectedOrderID := "some order id"
	expectedRedirectUrl := "https://daytonswingsmackdown.com/landing"
	expectedAccessToken := "12345"
	expectedUserID := "67890"
	expectedDiscountCode := "DJ"
	active := true

	expectedIdempotencyKey, err := uuid.NewV4()
	if err != nil {
		t.Fatalf("error generating idempotency key for test: %v", err)
	}

	logger := logrus.New()
	devnull, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatalf("error opening null: %v", err)
	}
	logger.SetOutput(devnull)
	logger.AddHook(&test_utility.ErrorHook{T: t})

	co := commontest.CommonCatalogObjects()

	inventoryCounts := make([]*square.InventoryCount, len(co.WeekendPassID))
	idx := 0
	for _, id := range co.WeekendPassID {
		inventoryCounts[idx] = &square.InventoryCount{
			CatalogObjectId: id,
			Quantity:        "25",
		}
		idx++
	}

	expectedLocationID := "here"

	tests := []struct {
		name         string
		registration *Info
		makeOrder    bool
	}{
		{
			name:      "all_items",
			makeOrder: true,
			registration: &Info{
				FirstName:       "John",
				LastName:        "Smith",
				StreetAddress:   "123 Any Street",
				City:            "Anytown",
				State:           "WA",
				ZipCode:         "12345",
				Email:           "benzejaa@gmail.com",
				HomeScene:       "Anytown Swing Cats",
				IsStudent:       true,
				PassType:        &WeekendPass{Tier: storage.Tier1, Level: storage.Level3},
				MixAndMatch:     &MixAndMatch{Role: storage.MixAndMatchRoleLeader},
				SoloJazz:        &SoloJazz{},
				TeamCompetition: &TeamCompetition{Name: "Anytown Anywhere"},
				TShirt:          &TShirt{Style: storage.TShirtStyleBellaS},
				Housing: &storage.RequireHousing{
					PetAllergies: "cats",
					Details:      "please house me with my 100 closest friends",
				},
				DiscountCodes: []string{expectedDiscountCode},
			},
		},
		{
			name:      "dance_only_provide_housing",
			makeOrder: true,
			registration: &Info{
				FirstName:     "John",
				LastName:      "Smith",
				StreetAddress: "123 Any Street",
				City:          "Anytown",
				State:         "WA",
				ZipCode:       "12345",
				Email:         "benzejaa@gmail.com",
				HomeScene:     "Anytown Swing Cats",
				IsStudent:     false,
				PassType:      &DanceOnlyPass{},
				Housing: &storage.ProvideHousing{
					Quantity: 100,
					Pets:     "inifinite dogs",
					Details:  "I don't believe in beds, all guests sleep outside",
				},
			},
		},
		{
			name: "no_items",
			registration: &Info{
				FirstName:     "John",
				LastName:      "Smith",
				StreetAddress: "123 Any Street",
				City:          "Anytown",
				State:         "WA",
				ZipCode:       "12345",
				Email:         "benzejaa@gmail.com",
				HomeScene:     "Anytown Swing Cats",
				IsStudent:     false,
				PassType:      &NoPass{},
				Housing:       &storage.NoHousing{},
			},
		},
	}

	authorizer := &commontest.MockAuthorizer{
		UserinfoFunc: commontest.UserinfoFromIDCheck(t, expectedAccessToken, expectedUserID),
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			client := &commontest.MockSquareClient{
				ListCatalogFunc: func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
					if !test.makeOrder {
						t.Fatalf("no orderable items found, square should not be called")
					}
					return commontest.ListCatalogFuncFromSlice(co.Catalog())(ctx, types)
				},
				BatchRetrieveInventoryCountsFunc: func(ctx context.Context, catalogObjectIds, locationIds []string, updatedAfter *time.Time) square.BatchRetrieveInventoryCountsIterator {
					if !test.makeOrder {
						t.Fatalf("no orderable items found, square should not be called")
					}
					return commontest.InventoryCountsFromSliceCheck(t, co.WeekendPassID, inventoryCounts)(ctx, catalogObjectIds, locationIds, updatedAfter)
				},
				ListLocationsFunc: func(context.Context) ([]*square.Location, error) {
					if !test.makeOrder {
						t.Fatalf("no orderable items found, square should not be called")
					}
					return []*square.Location{{Id: expectedLocationID}}, nil
				},
				CreateCheckoutFunc: func(ctx context.Context, locationID, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error) {
					if !test.makeOrder {
						t.Fatalf("no orderable items found, square should not be called")
					}
					if locationID != expectedLocationID {
						t.Fatalf("expected location ID %s, found %s", expectedLocationID, locationID)
					}
					if idempotencyKey != expectedIdempotencyKey.String() {
						t.Fatalf("expected idempotencyKey %v, found %s", expectedIdempotencyKey, idempotencyKey)
					}
					if merchantSupportEmail != utility.SmackdownEmail {
						t.Fatalf("expected merchant email %s, found %s", utility.SmackdownEmail, merchantSupportEmail)
					}
					if prePopulateBuyerEmail != test.registration.Email {
						t.Fatalf("expected user email %s, found %s", test.registration.Email, prePopulateBuyerEmail)
					}
					if redirectUrl != expectedRedirectUrl {
						t.Fatalf("expected redirectUrl %s, found %s", expectedRedirectUrl, redirectUrl)
					}

					itemChecks := []*itemCheck{}
					switch p := test.registration.PassType.(type) {
					case *WeekendPass:
						itemChecks = append(itemChecks, &itemCheck{id: co.WeekendPassID[p.Tier]})
					case *DanceOnlyPass:
						itemChecks = append(itemChecks, &itemCheck{id: co.DancePassID})
						// default, do nothing
					}
					if test.registration.MixAndMatch != nil {
						itemChecks = append(itemChecks, &itemCheck{id: co.MixAndMatchID})
					}
					if test.registration.TeamCompetition != nil {
						itemChecks = append(itemChecks, &itemCheck{id: co.TeamCompetitionID})
					}
					if test.registration.TShirt != nil {
						itemChecks = append(itemChecks, &itemCheck{id: co.TShirtID})
					}
					if test.registration.SoloJazz != nil {
						itemChecks = append(itemChecks, &itemCheck{id: co.SoloJazzID})
					}
					for _, lineItem := range order.Order.LineItems {
						if lineItem.Quantity != "1" {
							t.Fatalf("found unknown quantity of items %s, expected \"1\"", lineItem.Quantity)
						}
						found := false
						for _, itemCheck := range itemChecks {
							if itemCheck.id == lineItem.CatalogObjectId {
								if itemCheck.found {
									t.Fatalf("order item with id %q found twice", itemCheck.id)
								}
								itemCheck.found = true
								found = true

								if p, ok := test.registration.PassType.(*WeekendPass); ok && lineItem.CatalogObjectId == co.WeekendPassID[p.Tier] {
									if len(test.registration.DiscountCodes) > 0 {
										discountCheck(t, lineItem.Discounts, co.FullWeekendDiscountID)
									}
									if test.registration.IsStudent {
										discountCheck(t, lineItem.Discounts, co.StudentDiscountID)
									}
								} else if test.registration.MixAndMatch != nil && lineItem.CatalogObjectId == co.MixAndMatchID && len(test.registration.DiscountCodes) > 0 {
									discountCheck(t, lineItem.Discounts, co.MixAndMatchDiscountID)
								}
								break
							}
						}
						if !found {
							t.Fatalf("found order for unexpected item id %q", lineItem.CatalogObjectId)
						}

					}
					for _, itemCheck := range itemChecks {
						if !itemCheck.found {
							t.Fatalf("item with id %q not found", itemCheck.id)
						}
					}
					return &square.Checkout{
						CheckoutPageUrl: expectedCheckoutUrl,
						Order: &square.Order{
							Id: expectedOrderID,
						},
					}, nil
				},
			}

			storeAdded := false
			store := &commontest.MockStore{
				AddRegistrationFunc: func(ctx context.Context, r *storage.Registration) (string, error) {
					storeAdded = true
					if r.FirstName != test.registration.FirstName {
						t.Fatalf("expected registration first name %s, found %s", test.registration.FirstName, r.FirstName)
					}
					if r.LastName != test.registration.LastName {
						t.Fatalf("expected registration last name %s, found %s", test.registration.LastName, r.LastName)
					}
					if r.StreetAddress != test.registration.StreetAddress {
						t.Fatalf("expected registration street address %s, found %s", test.registration.StreetAddress, r.StreetAddress)
					}
					if r.City != test.registration.City {
						t.Fatalf("expected registration city %s, found %s", test.registration.City, r.City)
					}
					if r.State != test.registration.State {
						t.Fatalf("expected registration state %s, found %s", test.registration.State, r.State)
					}
					if r.ZipCode != test.registration.ZipCode {
						t.Fatalf("expected registration zip code %s, found %s", test.registration.ZipCode, r.ZipCode)
					}
					if r.Email != test.registration.Email {
						t.Fatalf("expected registration email %s, found %s", test.registration.Email, r.Email)
					}
					if r.HomeScene != test.registration.HomeScene {
						t.Fatalf("expected registration home scene %s, found %s", test.registration.HomeScene, r.HomeScene)
					}
					if r.IsStudent != test.registration.IsStudent {
						t.Fatalf("expected registration student status %v, found %v", test.registration.IsStudent, r.IsStudent)
					}
					if (r.SoloJazz) != (test.registration.SoloJazz != nil) {
						t.Fatalf("expected registration solo jazz purchase status %v, found %v", test.registration.SoloJazz, r.SoloJazz)
					}
					switch pt := test.registration.PassType.(type) {
					case *WeekendPass:
						exp, ok := r.PassType.(*storage.WeekendPass)
						if !ok {
							t.Fatalf("expected full weekend pass type, didn't find it.")
						}
						if exp.Tier != pt.Tier {
							t.Fatalf("found full weekend pass tier %v, expected %v", exp.Tier, pt.Tier)
						}
						if exp.Level != pt.Level {
							t.Fatalf("found full weekend pass level %v, expected %v", exp.Level, pt.Level)
						}
					case *DanceOnlyPass:
						_, ok := r.PassType.(*storage.DanceOnlyPass)
						if !ok {
							t.Fatalf("expected full dance only pass type, didn't find it.")
						}
					case *NoPass:
						_, ok := r.PassType.(*storage.NoPass)
						if !ok {
							t.Fatalf("expected full no pass type, didn't find it.")
						}

					}
					if test.registration.MixAndMatch != nil {
						if r.MixAndMatch == nil {
							t.Fatalf("found no mix and match data when expected")
						}
						if test.registration.MixAndMatch.Role != r.MixAndMatch.Role {
							t.Fatalf("found unexpected role %v, expected %v", r.MixAndMatch.Role, test.registration.MixAndMatch.Role)
						}
					} else if r.MixAndMatch != nil {
						t.Fatalf("found mix and match when not expected")
					}

					if test.registration.TeamCompetition != nil {
						if r.TeamCompetition == nil {
							t.Fatalf("found no team competition data when expected")
						}
						if test.registration.TeamCompetition.Name != r.TeamCompetition.Name {
							t.Fatalf("found unexpected name %s, expected %s", r.TeamCompetition.Name, test.registration.TeamCompetition.Name)
						}
					} else if r.TeamCompetition != nil {
						t.Fatalf("found team competition when not expected")
					}

					if test.registration.TShirt != nil {
						if r.TShirt == nil {
							t.Fatalf("found no tshirt data when expected")
						}
						if test.registration.TShirt.Style != r.TShirt.Style {
							t.Fatalf("found unexpected style %s, expected %s", r.TShirt.Style, test.registration.TShirt.Style)
						}
					} else if r.TShirt != nil {
						t.Fatalf("found tshirt when not expected")
					}
					if !reflect.DeepEqual(r.Housing, test.registration.Housing) {
						t.Fatalf("expected registration housing %#q, found %#q", test.registration.Housing, r.Housing)
					}
					if r.UserId != expectedUserID {
						t.Fatalf("expected registration user id %s, found %s", r.UserId, expectedUserID)
					}
					if len(r.DiscountCodes) != len(test.registration.DiscountCodes) {
						t.Fatalf("expected %d discounts, found %d", len(r.DiscountCodes), len(test.registration.DiscountCodes))
					}
					if len(r.DiscountCodes) == 1 && r.DiscountCodes[0] != expectedDiscountCode {
						t.Fatalf("expected registration discount codes  %#q, found %#q", []string{expectedDiscountCode}, r.DiscountCodes)
					}
					if test.makeOrder {
						if len(r.OrderIds) != 1 || r.OrderIds[0] != expectedOrderID {
							t.Fatalf("expected registration order id %#q, found %#q", []string{expectedOrderID}, r.OrderIds)
						}
					} else if len(r.OrderIds) != 0 {
						t.Fatalf("expected no order, found one")
					}
					return "some key", nil
				},
				GetDiscountFunc: func(ctx context.Context, code string) (*storage.Discount, error) {
					if code != test.registration.DiscountCodes[0] {
						t.Fatalf("expected discount code %s, found %s", test.registration.DiscountCodes[0], code)
					}
					return &storage.Discount{
						ID:   "123534",
						Code: code,
						Discounts: []*storage.SingleDiscount{
							{
								Name:      co.FullWeekendDiscountName,
								AppliedTo: storage.FullWeekendPurchaseItem,
							},
							{
								Name:      co.MixAndMatchDiscountName,
								AppliedTo: storage.MixAndMatchPurchaseItem,
							},
						},
					}, nil
				},
			}

			mailSent := false
			mailClient := &commontest.MockMailClient{
				SendFunc: func(msg *mail.SGMailV3) (*rest.Response, error) {
					mailSent = true
					if len(msg.Personalizations) != 1 {
						t.Fatalf("xpected 1 personalization, found %d", len(msg.Personalizations))
					}
					dt := msg.Personalizations[0].DynamicTemplateData
					if dt[mailFirstNameKey] != test.registration.FirstName {
						t.Fatalf("expected registration first name %s, found %s", test.registration.FirstName, dt[mailFirstNameKey])
					}
					if dt[mailLastNameKey] != test.registration.LastName {
						t.Fatalf("expected registration last name %s, found %s", test.registration.LastName, dt[mailLastNameKey])
					}
					if dt[mailStreetAddressKey] != test.registration.StreetAddress {
						t.Fatalf("expected registration street address %s, found %s", test.registration.StreetAddress, dt[mailStreetAddressKey])
					}
					if dt[mailCityKey] != test.registration.City {
						t.Fatalf("expected registration city %s, found %s", test.registration.City, dt[mailCityKey])
					}
					if dt[mailStateKey] != test.registration.State {
						t.Fatalf("expected registration state %s, found %s", test.registration.State, dt[mailStateKey])
					}
					if dt[mailZipCodeKey] != test.registration.ZipCode {
						t.Fatalf("expected registration zip code %s, found %s", test.registration.ZipCode, dt[mailZipCodeKey])
					}
					if dt[mailEmailKey] != test.registration.Email {
						t.Fatalf("expected registration email %s, found %s", test.registration.Email, dt[mailEmailKey])
					}
					if dt[mailHomeSceneKey] != test.registration.HomeScene {
						t.Fatalf("expected registration home scene %s, found %s", test.registration.HomeScene, dt[mailHomeSceneKey])
					}
					if dt[mailStudentKey] != test.registration.IsStudent {
						t.Fatalf("expected registration student status %v, found %v", test.registration.IsStudent, dt[mailStudentKey])
					}
					if dt[mailSoloJazzKey] != test.registration.SoloJazz {
						t.Fatalf("expected registration solo jazz purchase status %v, found %v", test.registration.SoloJazz, dt[mailSoloJazzKey])
					}
					if dt[mailMixAndMatchKey] != (test.registration.MixAndMatch != nil) {
						t.Fatalf("expected registration mix and match purchase status %v, found %v", test.registration.MixAndMatch != nil, dt[mailMixAndMatchKey])
					}
					if test.registration.MixAndMatch != nil && dt[mailMixAndMatchRoleKey] != test.registration.MixAndMatch.Role {
						t.Fatalf("expected registration mix and match role %s, found %s", test.registration.MixAndMatch.Role, dt[mailMixAndMatchRoleKey])
					}
					if dt[mailTeamCompKey] != (test.registration.TeamCompetition != nil) {
						t.Fatalf("expected registration team competition purchase status %v, found %v", test.registration.TeamCompetition != nil, dt[mailTeamCompKey])
					}
					if test.registration.TeamCompetition != nil && dt[mailTeamCompNameKey] != test.registration.TeamCompetition.Name {
						t.Fatalf("expected registration team competition name %s, found %s", test.registration.TeamCompetition.Name, dt[mailTeamCompNameKey])
					}
					if dt[mailTShirtKey] != (test.registration.TShirt != nil) {
						t.Fatalf("expected registration tshirt purchase status %v, found %v", test.registration.TShirt != nil, dt[mailTShirtKey])
					}
					if test.registration.TShirt != nil && dt[mailTShirtStyleKey] != test.registration.TShirt.Style {
						t.Fatalf("expected registration t-shirt style %s, found %s", test.registration.TShirt.Style, dt[mailTShirtStyleKey])
					}

					var expectedWeekendPassValue interface{}
					switch p := test.registration.PassType.(type) {
					case *WeekendPass:
						expectedWeekendPassValue = mailFullWeekendValue
						if dt[mailWorkshopLevelKey] != p.Level {
							t.Fatalf("expected pass level %d, found %d", p.Level, dt[mailWorkshopLevelKey])
						}
					case *DanceOnlyPass:
						expectedWeekendPassValue = mailDanceOnlyValue
					default:
						expectedWeekendPassValue = mailNoPassValue
					}
					if dt[mailWeekendPassKey] != expectedWeekendPassValue {
						t.Fatalf("expected weekend pass value %q, found %q", expectedWeekendPassValue, dt[mailWeekendPassKey])
					}

					var expectedHousingValue interface{}
					switch h := test.registration.Housing.(type) {
					case *storage.ProvideHousing:
						expectedHousingValue = mailProvideHousingValue
						ph, ok := dt[mailProvideHousingKey]
						if !ok {
							t.Fatalf("No provide housing data found in mail")
						}
						phmap, ok := ph.(map[string]interface{})
						if !ok {
							t.Fatalf("could not convert provide housing data to map")
						}
						if phmap[mailProvideHousingGuestsKey] != h.Quantity {
							t.Fatalf("expected guests quantity %d, found %d", h.Quantity, phmap[mailProvideHousingGuestsKey])
						}
						if phmap[mailProvideHousingPetsKey] != h.Pets {
							t.Fatalf("expected pet information %s, found %s", h.Pets, phmap[mailProvideHousingPetsKey])
						}
						if phmap[mailProvideHousingDetailsKey] != h.Details {
							t.Fatalf("expected provide housing detail information %s, found %s", h.Details, phmap[mailProvideHousingDetailsKey])
						}
					case *storage.RequireHousing:
						expectedHousingValue = mailRequireHousingValue
						rh, ok := dt[mailRequireHousingKey]
						if !ok {
							t.Fatalf("No provide housing data found in mail")
						}
						rhmap, ok := rh.(map[string]interface{})
						if !ok {
							t.Fatalf("could not convert require housing data to map")
						}
						if rhmap[mailRequireHousingAllergiesKey] != h.PetAllergies {
							t.Fatalf("expected allergy information %s, found %s", h.PetAllergies, rhmap[mailRequireHousingAllergiesKey])
						}
						if rhmap[mailRequireHousingDetailsKey] != h.Details {
							t.Fatalf("expected require housing detail information %s, found %s", h.Details, rhmap[mailRequireHousingDetailsKey])
						}
					default:
						expectedHousingValue = mailNoHousingValue
					}
					if dt[mailHousingKey] != expectedHousingValue {
						t.Fatalf("expected housing value %q, found %q", expectedHousingValue, dt[mailHousingKey])
					}
					return &rest.Response{
						StatusCode: http.StatusAccepted,
					}, nil
				},
			}

			service := NewService(active, false, logger, client, authorizer, store, mailClient)

			checkoutUrl, err := service.Add(context.Background(), test.registration, expectedRedirectUrl, expectedIdempotencyKey.String(), expectedAccessToken)
			if err != nil {
				t.Fatalf("error found in call to add: %v", err)
			}
			if test.makeOrder {
				if checkoutUrl != expectedCheckoutUrl {
					t.Fatalf("expected checkout url %s, found %s", expectedCheckoutUrl, checkoutUrl)
				}
			} else if checkoutUrl != expectedRedirectUrl {
				t.Fatalf("expected checkout url %s, found %s", expectedCheckoutUrl, expectedRedirectUrl)
			}
			if !mailSent {
				t.Fatalf("Mail was not sent to user")
			}
			if !storeAdded {
				t.Fatalf("Registration not added to store")
			}
		})
	}
}

func TestAddNotActive(t *testing.T) {
	active := false
	logger := logrus.New()
	devnull, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatalf("error opening null: %v", err)
	}
	logger.SetOutput(devnull)

	service := NewService(active, false, logger, &commontest.MockSquareClient{}, &commontest.MockAuthorizer{}, &commontest.MockStore{}, &commontest.MockMailClient{})

	registration := &Info{
		FirstName: "John",
		LastName:  "Smith",
		Email:     "benzejaa@gmail.com",
		PassType:  &NoPass{},
		Housing:   &storage.NoHousing{},
	}

	_, err = service.Add(context.Background(), registration, "https://smackdown.com", "7", "7")
	if err == nil {
		t.Fatalf("service not active, expected error, found none")
	}
}

func TestAddCostNothing(t *testing.T) {
	active := true
	logger := logrus.New()
	devnull, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatalf("error opening null: %v", err)
	}
	logger.SetOutput(devnull)

	co := commontest.CommonCatalogObjects()

	inventoryCounts := make([]*square.InventoryCount, len(co.WeekendPassID))
	idx := 0
	for _, id := range co.WeekendPassID {
		inventoryCounts[idx] = &square.InventoryCount{
			CatalogObjectId: id,
			Quantity:        "25",
		}
		idx++
	}

	client := &commontest.MockSquareClient{
		ListCatalogFunc:                  commontest.ListCatalogFuncFromSlice(co.Catalog()),
		BatchRetrieveInventoryCountsFunc: commontest.InventoryCountsFromSlice(inventoryCounts),
		ListLocationsFunc: func(context.Context) ([]*square.Location, error) {
			return []*square.Location{{Id: "7"}}, nil
		},
		CreateCheckoutFunc: func(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error) {
			return nil, &square.ErrorList{
				Errors: []*square.Error{
					{
						Category: square.ErrorCategoryInvalidRequestError,
						Code:     square.ErrorCodeValueTooLow,
						Field:    "order.total_money.amount",
					},
				},
			}
		},
	}

	authorizer := &commontest.MockAuthorizer{
		UserinfoFunc: func(ctx context.Context, accessToken string) (*authorizer.Userinfo, error) {
			return &authorizer.Userinfo{
				UserId: "1235",
			}, nil
		},
	}

	registration := &Info{
		FirstName: "John",
		LastName:  "Smith",
		Email:     "benzejaa@gmail.com",
		PassType:  &DanceOnlyPass{},
		Housing:   &storage.NoHousing{},
	}

	storeAdded := false
	store := &commontest.MockStore{
		AddRegistrationFunc: func(ctx context.Context, r *storage.Registration) (string, error) {
			storeAdded = true
			if r.FirstName != registration.FirstName {
				t.Fatalf("expected registration first name %s, found %s", registration.FirstName, r.FirstName)
			}
			if r.LastName != registration.LastName {
				t.Fatalf("expected registration last name %s, found %s", registration.LastName, r.LastName)
			}
			if r.Email != registration.Email {
				t.Fatalf("expected registration email %s, found %s", registration.Email, r.Email)
			}
			return "key", nil
		},
	}

	mailSent := false
	mailClient := &commontest.MockMailClient{
		SendFunc: func(msg *mail.SGMailV3) (*rest.Response, error) {
			mailSent = true
			if len(msg.Personalizations) != 1 {
				t.Fatalf("xpected 1 personalization, found %d", len(msg.Personalizations))
			}
			dt := msg.Personalizations[0].DynamicTemplateData
			if dt[mailFirstNameKey] != registration.FirstName {
				t.Fatalf("expected registration first name %s, found %s", registration.FirstName, dt[mailFirstNameKey])
			}
			if dt[mailLastNameKey] != registration.LastName {
				t.Fatalf("expected registration last name %s, found %s", registration.LastName, dt[mailLastNameKey])
			}
			if dt[mailEmailKey] != registration.Email {
				t.Fatalf("expected registration email %s, found %s", registration.Email, dt[mailEmailKey])
			}
			return &rest.Response{
				StatusCode: http.StatusAccepted,
			}, nil
		},
	}

	service := NewService(active, false, logger, client, authorizer, store, mailClient)

	redirectUrl := "https://smackdown.com"
	checkoutUrl, err := service.Add(context.Background(), registration, redirectUrl, "7", "7")
	if err != nil {
		t.Fatalf("error found in call to add: %v", err)
	}
	if checkoutUrl != redirectUrl {
		t.Fatalf("expected checkout url %s, found %s", checkoutUrl, redirectUrl)
	}
	if !mailSent {
		t.Fatalf("Mail was not sent to user")
	}
	if !storeAdded {
		t.Fatalf("Registration not added to store")
	}
}
