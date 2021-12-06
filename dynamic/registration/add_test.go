package registration

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/commontest"
	"github.com/Houndie/dss-registration/dynamic/sendinblue"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/test_utility"
	"github.com/Houndie/square-go"
	"github.com/Houndie/square-go/checkout"
	"github.com/Houndie/square-go/locations"
	"github.com/Houndie/square-go/objects"
	"github.com/sirupsen/logrus"
)

var testPermissionConfig = &PermissionConfig{
	List: "list:registrations",
}

type itemCheck struct {
	found bool
	id    string
}

/*func registrationCheck(t *testing.T, testRegistration, controlRegistration *Info, controlID string) {
	t.Helper()

	if testRegistration.ID != controlID {
		t.Fatalf("expected test id %v, found %v", testRegistration.ID, controlID)
	}

	if testRegistration.FirstName != controlRegistration.FirstName {
		t.Fatalf("expected first name %v, found %v", testRegistration.FirstName, controlRegistration.FirstName)
	}

	if testRegistration.LastName != controlRegistration.LastName {
		t.Fatalf("expected last name %v, found %v", testRegistration.LastName, controlRegistration.LastName)
	}

	if testRegistration.StreetAddress != controlRegistration.StreetAddress {
		t.Fatalf("expected street address %v, found %v", testRegistration.StreetAddress, controlRegistration.StreetAddress)
	}

	if testRegistration.City != controlRegistration.City {
		t.Fatalf("expected city %v, found %v", testRegistration.City, controlRegistration.City)
	}

	if testRegistration.State != controlRegistration.State {
		t.Fatalf("expected state %v, found %v", testRegistration.State, controlRegistration.State)
	}

	if testRegistration.ZipCode != controlRegistration.ZipCode {
		t.Fatalf("expected zip code %v, found %v", testRegistration.ZipCode, controlRegistration.ZipCode)
	}

	if testRegistration.Email != controlRegistration.Email {
		t.Fatalf("expected email %v, found %v", testRegistration.Email, controlRegistration.Email)
	}

	if testRegistration.HomeScene != controlRegistration.HomeScene {
		t.Fatalf("expected home scene %v, found %v", testRegistration.HomeScene, controlRegistration.HomeScene)
	}

	if testRegistration.IsStudent != controlRegistration.IsStudent {
		t.Fatalf("expected student status %v, found %v", testRegistration.IsStudent, controlRegistration.IsStudent)
	}

	switch t := testRegistration.PassType.(type) {
	case *WeekendPass:
	}

}*/

func discountCheck(t *testing.T, discountArray []*objects.OrderLineItemDiscount, appliedDiscounts []*objects.OrderLineItemAppliedDiscount, discountID string) {
	t.Helper()
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
	expectedAccessToken := "12345"
	expectedUserID := "67890"
	active := true
	registrationID := "some key"

	logger := logrus.New()
	devnull, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatalf("error opening null: %v", err)
	}
	logger.SetOutput(devnull)
	logger.AddHook(&test_utility.ErrorHook{T: t})

	co := commontest.CommonCatalogObjects()

	inventoryCounts := make([]*objects.InventoryCount, len(co.WeekendPassID))
	idx := 0
	for _, id := range co.WeekendPassID {
		inventoryCounts[idx] = &objects.InventoryCount{
			CatalogObjectID: id,
			Quantity:        "25",
		}
		idx++
	}

	tests := []struct {
		name         string
		registration *Info
	}{
		{
			name: "all_items",
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
				DiscountCodes: []string{co.FullWeekendDiscountName, co.MixAndMatchDiscountName},
			},
		},
		{
			name: "dance_only_provide_housing",
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
		GetUserinfoFunc: commontest.UserinfoFromIDCheck(t, expectedAccessToken, []string{}, expectedUserID, []string{}),
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			client := &square.Client{}

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
					if test.registration.SoloJazz != nil {
						if r.SoloJazz == nil {
							t.Fatalf("expected registration solo jazz purchase status %v, found %v", test.registration.SoloJazz, r.SoloJazz)
						}

						if test.registration.SoloJazz.AdminPaymentOverride != r.SoloJazz.ManuallyPaid {
							t.Fatalf("expected registration solo jazz purchase admin override %v, found %v", test.registration.SoloJazz.AdminPaymentOverride, r.SoloJazz.ManuallyPaid)

						}
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
					if len(r.DiscountCodes) != 2 && len(r.DiscountCodes) != 0 {
						t.Fatalf("expected 2 or 0 registration discount codes, found %d", len(r.DiscountCodes))
					}
					if len(r.OrderIDs) != 0 {
						t.Fatalf("expected no order, found one")
					}
					return registrationID, nil
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

			service := NewService(active, false, logger, client, commontest.CommonCatalogObjects().SquareData(), authorizer, store, mailClient, testPermissionConfig)

			outputRegistration, err := service.Add(context.Background(), test.registration, expectedAccessToken)
			if err != nil {
				t.Fatalf("error found in call to add: %v", err)
			}

			if outputRegistration.ID != registrationID {
				t.Fatalf("expected ID %s, found %s", registrationID, outputRegistration.ID)
			}

			outputRegistration.ID = test.registration.ID
			if !reflect.DeepEqual(test.registration, outputRegistration) {
				t.Fatalf("expected registration %v, found %v", outputRegistration, test.registration)
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

	service := NewService(active, false, logger, &square.Client{}, commontest.CommonCatalogObjects().SquareData(), &commontest.MockAuthorizer{}, &commontest.MockStore{}, &commontest.MockMailClient{}, testPermissionConfig)

	registration := &Info{
		FirstName: "John",
		LastName:  "Smith",
		Email:     "benzejaa@gmail.com",
		PassType:  &NoPass{},
		Housing:   &storage.NoHousing{},
	}

	_, err = service.Add(context.Background(), registration, "7")
	if err == nil {
		t.Fatalf("service not active, expected error, found none")
	}
}

func TestAddCostNothing(t *testing.T) {
	registrationID := "key"
	active := true
	logger := logrus.New()
	devnull, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatalf("error opening null: %v", err)
	}
	logger.SetOutput(devnull)

	co := commontest.CommonCatalogObjects()

	inventoryCounts := make([]*objects.InventoryCount, len(co.WeekendPassID))
	idx := 0
	for _, id := range co.WeekendPassID {
		inventoryCounts[idx] = &objects.InventoryCount{
			CatalogObjectID: id,
			Quantity:        "25",
		}
		idx++
	}

	client := &square.Client{
		Inventory: &commontest.MockSquareInventoryClient{
			BatchRetrieveCountsFunc: commontest.InventoryCountsFromSlice(inventoryCounts),
		},
		Locations: &commontest.MockSquareLocationsClient{
			ListFunc: func(context.Context, *locations.ListRequest) (*locations.ListResponse, error) {
				return &locations.ListResponse{
					Locations: []*objects.Location{{ID: "7"}},
				}, nil
			},
		},
		Checkout: &commontest.MockSquareCheckoutClient{
			CreateFunc: func(ctx context.Context, req *checkout.CreateRequest) (*checkout.CreateResponse, error) {
				return nil, &objects.ErrorList{
					Errors: []*objects.Error{
						{
							Category: objects.ErrorCategoryInvalidRequestError,
							Code:     objects.ErrorCodeValueTooLow,
							Field:    "order.total_money.amount",
						},
					},
				}
			},
		},
	}

	authorizer := &commontest.MockAuthorizer{
		GetUserinfoFunc: commontest.UserinfoFromID("12345", []string{}),
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
			return registrationID, nil
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

	service := NewService(active, false, logger, client, commontest.CommonCatalogObjects().SquareData(), authorizer, store, mailClient, testPermissionConfig)

	outputRegistration, err := service.Add(context.Background(), registration, "7")
	if err != nil {
		t.Fatalf("error found in call to add: %v", err)
	}
	if outputRegistration.ID != registrationID {
		t.Fatalf("expected ID %s, found %s", registrationID, outputRegistration.ID)
	}

	outputRegistration.ID = ""
	if !reflect.DeepEqual(registration, outputRegistration) {
		t.Fatalf("expected registration %v, found %v", outputRegistration, registration)
	}
	if !mailSent {
		t.Fatalf("Mail was not sent to user")
	}
	if !storeAdded {
		t.Fatalf("Registration not added to store")
	}
}
