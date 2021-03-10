package registration

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/commontest"
	"github.com/Houndie/dss-registration/dynamic/sendinblue"
	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/test_utility"
	"github.com/Houndie/dss-registration/dynamic/utility"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

type itemCheck struct {
	found bool
	id    string
}

func discountCheck(t *testing.T, discountArray []*square.OrderLineItemDiscount, appliedDiscounts []*square.OrderLineItemAppliedDiscount, discountID string) {
	found := false
	for _, d := range discountArray {
		if d.CatalogObjectID == discountID {
			found = true
			appliedFound := false
			for _, ad := range appliedDiscounts {
				if ad.DiscountUID == d.UID {
					appliedFound = true
					break
				}
			}
			if !appliedFound {
				t.Fatalf("discount %q found, but not it was not applied to this item", discountID)
			}
			break
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
			CatalogObjectID: id,
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
		GetUserinfoFunc: commontest.UserinfoFromIDCheck(t, expectedAccessToken, []authorizer.Permission{}, expectedUserID, []authorizer.Permission{}),
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
				BatchRetrieveInventoryCountsFunc: func(ctx context.Context, catalogObjectIDs, locationIDs []string, updatedAfter *time.Time) square.BatchRetrieveInventoryCountsIterator {
					if !test.makeOrder {
						t.Fatalf("no orderable items found, square should not be called")
					}
					return commontest.InventoryCountsFromSliceCheck(t, co.WeekendPassID, inventoryCounts)(ctx, catalogObjectIDs, locationIDs, updatedAfter)
				},
				ListLocationsFunc: func(context.Context) ([]*square.Location, error) {
					if !test.makeOrder {
						t.Fatalf("no orderable items found, square should not be called")
					}
					return []*square.Location{{ID: expectedLocationID}}, nil
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
							if itemCheck.id == lineItem.CatalogObjectID {
								if itemCheck.found {
									t.Fatalf("order item with id %q found twice", itemCheck.id)
								}
								itemCheck.found = true
								found = true

								if p, ok := test.registration.PassType.(*WeekendPass); ok && lineItem.CatalogObjectID == co.WeekendPassID[p.Tier] {
									if len(test.registration.DiscountCodes) > 0 {
										discountCheck(t, order.Order.Discounts, lineItem.AppliedDiscounts, co.FullWeekendDiscountID)
									}
									if test.registration.IsStudent {
										discountCheck(t, order.Order.Discounts, lineItem.AppliedDiscounts, co.StudentDiscountID)
									}
								} else if test.registration.MixAndMatch != nil && lineItem.CatalogObjectID == co.MixAndMatchID && len(test.registration.DiscountCodes) > 0 {
									discountCheck(t, order.Order.Discounts, lineItem.AppliedDiscounts, co.MixAndMatchDiscountID)
								}
								break
							}
						}
						if !found {
							t.Fatalf("found order for unexpected item id %q", lineItem.CatalogObjectID)
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
							ID: expectedOrderID,
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
					if r.UserID != expectedUserID {
						t.Fatalf("expected registration user id %s, found %s", r.UserID, expectedUserID)
					}
					if len(r.DiscountCodes) != len(test.registration.DiscountCodes) {
						t.Fatalf("expected %d discounts, found %d", len(r.DiscountCodes), len(test.registration.DiscountCodes))
					}
					if len(r.DiscountCodes) == 1 && r.DiscountCodes[0] != expectedDiscountCode {
						t.Fatalf("expected registration discount codes  %#q, found %#q", []string{expectedDiscountCode}, r.DiscountCodes)
					}
					if test.makeOrder {
						if len(r.OrderIDs) != 1 || r.OrderIDs[0] != expectedOrderID {
							t.Fatalf("expected registration order id %#q, found %#q", []string{expectedOrderID}, r.OrderIDs)
						}
					} else if len(r.OrderIDs) != 0 {
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
				SendSMTPEmailFunc: func(ctx context.Context, params *sendinblue.SMTPEmailParams) (string, error) {
					mailSent = true
					p, ok := params.Params.(*mailParams)
					if !ok {
						t.Fatal("unexpected type for mail parameters")
					}
					if p.FirstName != test.registration.FirstName {
						t.Fatalf("expected registration first name %s, found %s", test.registration.FirstName, p.FirstName)
					}
					if p.LastName != test.registration.LastName {
						t.Fatalf("expected registration last name %s, found %s", test.registration.LastName, p.LastName)
					}
					if p.StreetAddress != test.registration.StreetAddress {
						t.Fatalf("expected registration street address %s, found %s", test.registration.StreetAddress, p.StreetAddress)
					}
					if p.City != test.registration.City {
						t.Fatalf("expected registration city %s, found %s", test.registration.City, p.City)
					}
					if p.State != test.registration.State {
						t.Fatalf("expected registration state %s, found %s", test.registration.State, p.State)
					}
					if p.ZipCode != test.registration.ZipCode {
						t.Fatalf("expected registration zip code %s, found %s", test.registration.ZipCode, p.ZipCode)
					}
					if p.HomeScene != test.registration.HomeScene {
						t.Fatalf("expected registration home scene %s, found %s", test.registration.HomeScene, p.HomeScene)
					}
					if p.IsStudent != test.registration.IsStudent {
						t.Fatalf("expected registration student status %v, found %v", test.registration.IsStudent, p.IsStudent)
					}
					if p.SoloJazz.Purchased != (test.registration.SoloJazz != nil) {
						t.Fatalf("expected registration solo jazz purchase status %v, found %v", test.registration.SoloJazz != nil, p.SoloJazz.Purchased)
					}
					if test.registration.MixAndMatch != nil {
						if !p.MixAndMatch.Purchased {
							t.Fatalf("expected mix and match purchase, found none")
						}
						switch test.registration.MixAndMatch.Role {
						case storage.MixAndMatchRoleLeader:
							if p.MixAndMatch.Role != mailLeader {
								t.Fatalf("expected leader role, found %s", p.MixAndMatch.Role)
							}
						case storage.MixAndMatchRoleFollower:
							if p.MixAndMatch.Role != mailFollower {
								t.Fatalf("expected follower role, found %s", p.MixAndMatch.Role)
							}
						default:
							t.Fatalf("unknown control mix and match role")
						}
					}

					if test.registration.TeamCompetition != nil {
						if !p.TeamCompetition.Purchased {
							t.Fatalf("expected team competition, found none")
						}
						if p.TeamCompetition.Name != test.registration.TeamCompetition.Name {
							t.Fatalf("expected registration team competition name %s, found %s", test.registration.TeamCompetition.Name, p.TeamCompetition.Name)
						}
					}

					if test.registration.TShirt != nil {
						if !p.TShirt.Purchased {
							t.Fatalf("expected t-shirt purchase, found none")
						}
						switch test.registration.TShirt.Style {
						case storage.TShirtStyleUnisexS:
							if p.TShirt.Style != mailUnisexS {
								t.Fatalf("expected small unisex, found %s", p.TShirt.Style)
							}
						case storage.TShirtStyleUnisexM:
							if p.TShirt.Style != mailUnisexM {
								t.Fatalf("expected medium unisex, found %s", p.TShirt.Style)
							}
						case storage.TShirtStyleUnisexL:
							if p.TShirt.Style != mailUnisexL {
								t.Fatalf("expected large unisex, found %s", p.TShirt.Style)
							}
						case storage.TShirtStyleUnisexXL:
							if p.TShirt.Style != mailUnisexXL {
								t.Fatalf("expected extra large unisex, found %s", p.TShirt.Style)
							}
						case storage.TShirtStyleUnisex2XL:
							if p.TShirt.Style != mailUnisex2XL {
								t.Fatalf("expected 2XL unisex, found %s", p.TShirt.Style)
							}
						case storage.TShirtStyleUnisex3XL:
							if p.TShirt.Style != mailUnisex3XL {
								t.Fatalf("expected 3XL unisex, found %s", p.TShirt.Style)
							}
						case storage.TShirtStyleBellaS:
							if p.TShirt.Style != mailBellaS {
								t.Fatalf("expected small bella, found %s", p.TShirt.Style)
							}
						case storage.TShirtStyleBellaM:
							if p.TShirt.Style != mailBellaM {
								t.Fatalf("expected medium bella, found %s", p.TShirt.Style)
							}
						case storage.TShirtStyleBellaL:
							if p.TShirt.Style != mailBellaL {
								t.Fatalf("expected large bella, found %s", p.TShirt.Style)
							}
						case storage.TShirtStyleBellaXL:
							if p.TShirt.Style != mailBellaXL {
								t.Fatalf("expected extra large bella, found %s", p.TShirt.Style)
							}
						case storage.TShirtStyleBella2XL:
							if p.TShirt.Style != mailBella2XL {
								t.Fatalf("expected 2XL bella, found %s", p.TShirt.Style)
							}
						}
					}

					switch rp := test.registration.PassType.(type) {
					case *WeekendPass:
						if p.PassType.Type != mailWeekendPass {
							t.Fatalf("expected weekend pass found %v", p.PassType.Type)
						}
						if p.PassType.WeekendPass.Level != int(rp.Level) {
							t.Fatalf("expected pass level %d, found %d", rp.Level, p.PassType.WeekendPass.Level)
						}
					case *DanceOnlyPass:
						if p.PassType.Type != mailDanceOnlyPass {
							t.Fatalf("expected dance only pass found %v", p.PassType.Type)
						}
					default:
						if p.PassType.Type != mailNoPass {
							t.Fatalf("expected no pass found %v", p.PassType.Type)
						}
					}

					switch h := test.registration.Housing.(type) {
					case *storage.ProvideHousing:
						if p.Housing.Type != mailProvideHousing {
							t.Fatalf("expected providing housing found %v", p.PassType.Type)
						}
						if p.Housing.Provide.Quantity != h.Quantity {
							t.Fatalf("expected guests quantity %d, found %d", h.Quantity, p.Housing.Provide.Quantity)
						}
						if p.Housing.Provide.Pets != h.Pets {
							t.Fatalf("expected pet information %s, found %s", h.Pets, p.Housing.Provide.Pets)
						}
						if p.Housing.Provide.Details != h.Details {
							t.Fatalf("expected provide housing detail information %s, found %s", h.Details, p.Housing.Provide.Details)
						}
					case *storage.RequireHousing:
						if p.Housing.Type != mailRequireHousing {
							t.Fatalf("expected requiring housing found %v", p.PassType.Type)
						}
						if p.Housing.Require.PetAllergies != h.PetAllergies {
							t.Fatalf("expected allergy information %s, found %s", h.PetAllergies, p.Housing.Require.PetAllergies)
						}
						if p.Housing.Require.Details != h.Details {
							t.Fatalf("expected require housing detail information %s, found %s", h.Details, p.Housing.Require.Details)
						}
					case *storage.NoHousing:
						if p.Housing.Type != mailNoHousing {
							t.Fatalf("expected no housing found %v", p.PassType.Type)
						}
					}
					return "", nil
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
			CatalogObjectID: id,
			Quantity:        "25",
		}
		idx++
	}

	client := &commontest.MockSquareClient{
		ListCatalogFunc:                  commontest.ListCatalogFuncFromSlice(co.Catalog()),
		BatchRetrieveInventoryCountsFunc: commontest.InventoryCountsFromSlice(inventoryCounts),
		ListLocationsFunc: func(context.Context) ([]*square.Location, error) {
			return []*square.Location{{ID: "7"}}, nil
		},
		CreateCheckoutFunc: func(ctx context.Context, locationID, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error) {
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
		GetUserinfoFunc: commontest.UserinfoFromID("12345", []authorizer.Permission{}),
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
		SendSMTPEmailFunc: func(ctx context.Context, params *sendinblue.SMTPEmailParams) (string, error) {
			mailSent = true
			p, ok := params.Params.(*mailParams)
			if !ok {
				t.Fatal("unexpected type for mail parameters")
			}
			if p.FirstName != registration.FirstName {
				t.Fatalf("expected registration first name %s, found %s", registration.FirstName, p.FirstName)
			}
			if p.LastName != registration.LastName {
				t.Fatalf("expected registration last name %s, found %s", registration.LastName, p.LastName)
			}
			return "", nil
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
