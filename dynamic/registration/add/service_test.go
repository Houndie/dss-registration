package add

import (
	"context"
	"testing"
	"time"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/test_utility"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func TestAdd(t *testing.T) {

	registration := &Registration{
		FirstName:       "John",
		LastName:        "Doe",
		StreetAddress:   "123 Any St.",
		City:            "New York",
		State:           "NY",
		ZipCode:         "12345",
		Email:           "John.Doe@example.com",
		HomeScene:       "Frim Fram",
		IsStudent:       true,
		PassType:        &NoPass{},
		MixAndMatch:     nil,
		SoloJazz:        true,
		TeamCompetition: nil,
		TShirt:          nil,
		Housing:         &NoHousing{},
	}
	testRedirectUrl := "https://daytonswingsmackdown.com/registration-complete"
	expectedCheckoutUrl := "https://squareup.com/98734987349873498345"

	catalogObjectsIdx := -1

	testAccessToken := "some.access.token"
	testUserId := "12356"

	squareClient := &mockSquareClient{
		ListCatalogFunc: func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
			return &mockListCatalogIterator{
				ValueFunc: func() *square.CatalogObject {
					return catalogObjects[catalogObjectsIdx]
				},
				ErrorFunc: func() error {
					return nil
				},
				NextFunc: func() bool {
					catalogObjectsIdx++
					return catalogObjectsIdx < len(catalogObjects)
				},
			}
		},
		ListLocationsFunc: func(ctx context.Context) ([]*square.Location, error) {
			return locations, nil
		},
		CreateCheckoutFunc: func(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error) {
			createdAt := time.Now()
			if locationId != locations[0].Id {
				t.Fatalf("found location id %s, expected %s", locationId, locations[0].Id)
			}
			if order.Order.LocationId != locations[0].Id {
				t.Fatalf("found order location id %s, expected %s", order.Order.LocationId, locations[0].Id)
			}
			if len(order.Order.LineItems) != 1 {
				t.Fatalf("Found unexpected number of line items %d (expected 1)", len(order.Order.LineItems))
			}
			if order.Order.LineItems[0].Quantity != "1" {
				t.Fatalf("Ordering incorrect number of solo jazz tickets %s (expected 1)", order.Order.LineItems[0].Quantity)
			}
			if order.Order.LineItems[0].CatalogObjectId != soloJazzVariationId {
				t.Fatalf("Found incorrect item id %s in order, expected %s", order.Order.LineItems[0].CatalogObjectId, soloJazzVariationId)
			}

			if merchantSupportEmail != smackdownEmail {
				t.Fatalf("found incorrect merchant email %s, expected %s", merchantSupportEmail, smackdownEmail)
			}
			if prePopulateBuyerEmail != registration.Email {
				t.Fatalf("found incorrect buyer email %s, expected %s", prePopulateBuyerEmail, registration.Email)
			}
			if askForShippingAddress {
				t.Fatalf("unnecessarily asking user for shipping address")
			}
			if redirectUrl != testRedirectUrl {
				t.Fatalf("found incorrect redirect url %s, expected %s", redirectUrl, testRedirectUrl)
			}
			if len(additionalRecipients) != 0 {
				t.Fatalf("found additional %d recipients", len(additionalRecipients))
			}

			return &square.Checkout{
				Id:                         "checkout id",
				CheckoutPageUrl:            expectedCheckoutUrl,
				AskForShippingAddress:      askForShippingAddress,
				MerchantSupportEmail:       merchantSupportEmail,
				PrePopulateBuyerEmail:      prePopulateBuyerEmail,
				PrePopulateShippingAddress: prePopulateShippingAddress,
				RedirectUrl:                redirectUrl,
				Order:                      order.Order,
				CreatedAt:                  &createdAt,
				AdditionalRecipients:       nil,
			}, nil
		},
	}

	added := false
	store := &mockStore{
		AddRegistrationFunc: func(ctx context.Context, r *StoreRegistration) (string, error) {
			added = true
			if registration.FirstName != r.FirstName {
				t.Fatalf("found first name %s, expected %s", r.FirstName, registration.FirstName)
			}
			if registration.LastName != r.LastName {
				t.Fatalf("found last name %s, expected %s", r.LastName, registration.LastName)
			}
			if registration.StreetAddress != r.StreetAddress {
				t.Fatalf("found street address %s, expected %s", r.StreetAddress, registration.StreetAddress)
			}
			if registration.City != r.City {
				t.Fatalf("found city %s, expected %s", r.City, registration.City)
			}
			if registration.State != r.State {
				t.Fatalf("found state %s, expected %s", r.State, registration.State)
			}
			if registration.ZipCode != r.ZipCode {
				t.Fatalf("found zip code %s, expected %s", r.ZipCode, registration.ZipCode)
			}
			if registration.Email != r.Email {
				t.Fatalf("found email %s, expected %s", r.Email, registration.Email)
			}
			if registration.HomeScene != r.HomeScene {
				t.Fatalf("found home scene %s, expected %s", r.HomeScene, registration.HomeScene)
			}
			if registration.IsStudent != r.IsStudent {
				t.Fatalf("found student status %v, expected %v", r.IsStudent, registration.IsStudent)
			}
			if registration.SoloJazz != r.SoloJazz {
				t.Fatalf("found solo jazz ticket %v, expected %v", r.SoloJazz, registration.SoloJazz)
			}
			_, ok := r.PassType.(*NoPass)
			if !ok {
				t.Fatalf("found incorrect pass type")
			}
			_, ok = r.Housing.(*NoHousing)
			if !ok {
				t.Fatalf("found incorrect housing type")
			}
			if r.MixAndMatch != nil {
				t.Fatalf("found unexpected mix and match registration")
			}
			if r.TeamCompetition != nil {
				t.Fatalf("found unexpected team competition registration")
			}
			if r.TShirt != nil {
				t.Fatalf("found unexpected tshirt registration")
			}
			if r.UserId != testUserId {
				t.Fatalf("found incorrect userid %s, expected %s", r.UserId, testUserId)
			}
			return "store key", nil
		},
		DeleteRegistrationFunc: func(ctx context.Context, id string) error {
			t.Errorf("called delete registration")
			return nil
		},
	}

	logger := logrus.New()
	logger.SetOutput(&test_utility.ErrorWriter{t})
	logger.AddHook(&test_utility.ErrorHook{t})
	logger.SetLevel(logrus.TraceLevel)

	authorizer := &MockAuthorizer{
		UserinfoFunc: func(ctx context.Context, accessToken string) (*authorizer.Userinfo, error) {
			if accessToken != testAccessToken {
				t.Fatalf("found incorrect access token %s, expected %s", accessToken, testAccessToken)
			}
			return &authorizer.Userinfo{
				UserId: testUserId,
			}, nil
		},
	}

	url, err := NewService(logger, store, squareClient, authorizer).Add(context.Background(), registration, testRedirectUrl, testAccessToken)
	if err != nil {
		t.Fatalf("error adding new registration: %v", err)
	}
	if url != expectedCheckoutUrl {
		t.Fatalf("found url %s, expected %s", url, expectedCheckoutUrl)
	}
	if !added {
		t.Fatalf("registration not added to database")
	}
}

func TestAddPassTypes(t *testing.T) {
	testCases := []struct {
		passType PassType
		name     string
		itemId   string
	}{
		{
			passType: &WeekendPass{
				Level: WeekendPassLevel1,
				Tier:  WeekendPassTier3,
			},
			name:   "weekendpass",
			itemId: weekendPassTier3VariationId,
		},
		{
			passType: &DanceOnlyPass{},
			name:     "danceonly",
			itemId:   dancePassVariationId,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			testAccessToken := "some.access.token"
			testUserId := "12356"

			registration := &Registration{
				FirstName:       "John",
				LastName:        "Doe",
				StreetAddress:   "123 Any St.",
				City:            "New York",
				State:           "NY",
				ZipCode:         "12345",
				Email:           "John.Doe@example.com",
				HomeScene:       "Frim Fram",
				IsStudent:       true,
				PassType:        testCase.passType,
				MixAndMatch:     nil,
				SoloJazz:        false,
				TeamCompetition: nil,
				TShirt:          nil,
				Housing:         &NoHousing{},
			}
			testRedirectUrl := "https://daytonswingsmackdown.com/registration-complete"
			expectedCheckoutUrl := "https://squareup.com/98734987349873498345"

			catalogObjectsIdx := -1

			squareClient := &mockSquareClient{
				ListCatalogFunc: func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
					return &mockListCatalogIterator{
						ValueFunc: func() *square.CatalogObject {
							return catalogObjects[catalogObjectsIdx]
						},
						ErrorFunc: func() error {
							return nil
						},
						NextFunc: func() bool {
							catalogObjectsIdx++
							return catalogObjectsIdx < len(catalogObjects)
						},
					}
				},
				ListLocationsFunc: func(ctx context.Context) ([]*square.Location, error) {
					return locations, nil
				},
				CreateCheckoutFunc: func(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error) {
					if len(order.Order.LineItems) != 1 {
						t.Fatalf("found incorrect number of line items %d", len(order.Order.LineItems))
					}
					if order.Order.LineItems[0].CatalogObjectId != testCase.itemId {
						t.Fatalf("found incorrect square item id %s, expected %s", order.Order.LineItems[0].CatalogObjectId, testCase.itemId)
					}
					if order.Order.LineItems[0].Quantity != "1" {
						t.Fatalf("found incorrect number of pass orders %s, expected 1", order.Order.LineItems[0].Quantity)
					}
					createdAt := time.Now()
					return &square.Checkout{
						Id:                         "checkout id",
						CheckoutPageUrl:            expectedCheckoutUrl,
						AskForShippingAddress:      askForShippingAddress,
						MerchantSupportEmail:       merchantSupportEmail,
						PrePopulateBuyerEmail:      prePopulateBuyerEmail,
						PrePopulateShippingAddress: prePopulateShippingAddress,
						RedirectUrl:                redirectUrl,
						Order:                      order.Order,
						CreatedAt:                  &createdAt,
						AdditionalRecipients:       nil,
					}, nil
				},
			}

			store := &mockStore{
				AddRegistrationFunc: func(ctx context.Context, r *StoreRegistration) (string, error) {
					if !cmp.Equal(r.PassType, testCase.passType, cmpopts.IgnoreUnexported()) {
						t.Fatalf("found incorrect pass type data %s, expected %s", spew.Sdump(r.PassType), spew.Sdump(testCase.passType))
					}
					return "store key", nil
				},
				DeleteRegistrationFunc: func(ctx context.Context, id string) error {
					t.Error("delete registration called")
					return nil
				},
			}

			logger := logrus.New()
			logger.SetOutput(&test_utility.ErrorWriter{t})
			logger.AddHook(&test_utility.ErrorHook{t})
			logger.SetLevel(logrus.TraceLevel)

			authorizer := &MockAuthorizer{
				UserinfoFunc: func(ctx context.Context, accessToken string) (*authorizer.Userinfo, error) {
					return &authorizer.Userinfo{
						UserId: testUserId,
					}, nil
				},
			}

			_, err := NewService(logger, store, squareClient, authorizer).Add(context.Background(), registration, testRedirectUrl, testAccessToken)
			if err != nil {
				t.Fatalf("error adding new registration: %v", err)
			}
		})
	}
}

func TestAddHousing(t *testing.T) {
	testCases := []struct {
		housing Housing
		name    string
	}{
		{
			housing: &ProvideHousing{
				Pets:     "I have sevaral cute dogs",
				Quantity: 9,
				Details:  "I have a barn with lots of stalls for sleeping",
			},
			name: "providehousing",
		},
		{
			housing: &RequireHousing{
				PetAllergies: "If I am within 100 yards of a dog I will die",
				Details:      "Please house me with all 300 of my friends in a bed, thanks!",
			},
			name: "requirehousing",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			testAccessToken := "some.access.token"
			testUserId := "12356"

			registration := &Registration{
				FirstName:       "John",
				LastName:        "Doe",
				StreetAddress:   "123 Any St.",
				City:            "New York",
				State:           "NY",
				ZipCode:         "12345",
				Email:           "John.Doe@example.com",
				HomeScene:       "Frim Fram",
				IsStudent:       true,
				PassType:        &NoPass{},
				MixAndMatch:     nil,
				SoloJazz:        true,
				TeamCompetition: nil,
				TShirt:          nil,
				Housing:         testCase.housing,
			}
			testRedirectUrl := "https://daytonswingsmackdown.com/registration-complete"
			expectedCheckoutUrl := "https://squareup.com/98734987349873498345"

			catalogObjectsIdx := -1

			squareClient := &mockSquareClient{
				ListCatalogFunc: func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
					return &mockListCatalogIterator{
						ValueFunc: func() *square.CatalogObject {
							return catalogObjects[catalogObjectsIdx]
						},
						ErrorFunc: func() error {
							return nil
						},
						NextFunc: func() bool {
							catalogObjectsIdx++
							return catalogObjectsIdx < len(catalogObjects)
						},
					}
				},
				ListLocationsFunc: func(ctx context.Context) ([]*square.Location, error) {
					return locations, nil
				},
				CreateCheckoutFunc: func(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error) {
					createdAt := time.Now()
					return &square.Checkout{
						Id:                         "checkout id",
						CheckoutPageUrl:            expectedCheckoutUrl,
						AskForShippingAddress:      askForShippingAddress,
						MerchantSupportEmail:       merchantSupportEmail,
						PrePopulateBuyerEmail:      prePopulateBuyerEmail,
						PrePopulateShippingAddress: prePopulateShippingAddress,
						RedirectUrl:                redirectUrl,
						Order:                      order.Order,
						CreatedAt:                  &createdAt,
						AdditionalRecipients:       nil,
					}, nil
				},
			}

			store := &mockStore{
				AddRegistrationFunc: func(ctx context.Context, r *StoreRegistration) (string, error) {
					if !cmp.Equal(r.Housing, testCase.housing, cmpopts.IgnoreUnexported()) {
						t.Fatalf("found incorrect housing data %s, expected %s", spew.Sdump(r.Housing), spew.Sdump(testCase.housing))
					}
					return "store key", nil
				},
				DeleteRegistrationFunc: func(ctx context.Context, id string) error {
					t.Error("delete registration called")
					return nil
				},
			}

			logger := logrus.New()
			logger.SetOutput(&test_utility.ErrorWriter{t})
			logger.AddHook(&test_utility.ErrorHook{t})
			logger.SetLevel(logrus.TraceLevel)

			authorizer := &MockAuthorizer{
				UserinfoFunc: func(ctx context.Context, accessToken string) (*authorizer.Userinfo, error) {
					return &authorizer.Userinfo{
						UserId: testUserId,
					}, nil
				},
			}

			_, err := NewService(logger, store, squareClient, authorizer).Add(context.Background(), registration, testRedirectUrl, testAccessToken)
			if err != nil {
				t.Fatalf("error adding new registration: %v", err)
			}
		})
	}
}

func TestAddMixAndMatch(t *testing.T) {
	testAccessToken := "some.access.token"
	testUserId := "12356"

	registration := &Registration{
		FirstName:     "John",
		LastName:      "Doe",
		StreetAddress: "123 Any St.",
		City:          "New York",
		State:         "NY",
		ZipCode:       "12345",
		Email:         "John.Doe@example.com",
		HomeScene:     "Frim Fram",
		IsStudent:     true,
		PassType:      &NoPass{},
		MixAndMatch: &MixAndMatch{
			Role: MixAndMatchRoleLeader,
		},
		SoloJazz:        false,
		TeamCompetition: nil,
		TShirt:          nil,
		Housing:         &NoHousing{},
	}
	testRedirectUrl := "https://daytonswingsmackdown.com/registration-complete"
	expectedCheckoutUrl := "https://squareup.com/98734987349873498345"

	catalogObjectsIdx := -1

	squareClient := &mockSquareClient{
		ListCatalogFunc: func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
			return &mockListCatalogIterator{
				ValueFunc: func() *square.CatalogObject {
					return catalogObjects[catalogObjectsIdx]
				},
				ErrorFunc: func() error {
					return nil
				},
				NextFunc: func() bool {
					catalogObjectsIdx++
					return catalogObjectsIdx < len(catalogObjects)
				},
			}
		},
		ListLocationsFunc: func(ctx context.Context) ([]*square.Location, error) {
			return locations, nil
		},
		CreateCheckoutFunc: func(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error) {
			createdAt := time.Now()
			if len(order.Order.LineItems) != 1 {
				t.Fatalf("Found unexpected number of line items %d (expected 1)", len(order.Order.LineItems))
			}
			if order.Order.LineItems[0].CatalogObjectId != mixAndMatchVariationId {
				t.Fatalf("Found incorrect item id %s in order, expected %s", order.Order.LineItems[0].CatalogObjectId, soloJazzVariationId)
			}
			if order.Order.LineItems[0].Quantity != "1" {
				t.Fatalf("Ordering incorrect number of mix and match tickets %s (expected 1)", order.Order.LineItems[0].Quantity)
			}
			return &square.Checkout{
				Id:                         "checkout id",
				CheckoutPageUrl:            expectedCheckoutUrl,
				AskForShippingAddress:      askForShippingAddress,
				MerchantSupportEmail:       merchantSupportEmail,
				PrePopulateBuyerEmail:      prePopulateBuyerEmail,
				PrePopulateShippingAddress: prePopulateShippingAddress,
				RedirectUrl:                redirectUrl,
				Order:                      order.Order,
				CreatedAt:                  &createdAt,
				AdditionalRecipients:       nil,
			}, nil
		},
	}

	store := &mockStore{
		AddRegistrationFunc: func(ctx context.Context, r *StoreRegistration) (string, error) {
			if !cmp.Equal(r.MixAndMatch, registration.MixAndMatch, cmpopts.IgnoreUnexported()) {
				t.Fatalf("found incorrect mix and match data %s, expected %s", spew.Sdump(r.MixAndMatch), spew.Sdump(registration.MixAndMatch))
			}
			return "store key", nil
		},
		DeleteRegistrationFunc: func(ctx context.Context, id string) error {
			t.Error("delete registration called")
			return nil
		},
	}

	logger := logrus.New()
	logger.SetOutput(&test_utility.ErrorWriter{t})
	logger.AddHook(&test_utility.ErrorHook{t})
	logger.SetLevel(logrus.TraceLevel)

	authorizer := &MockAuthorizer{
		UserinfoFunc: func(ctx context.Context, accessToken string) (*authorizer.Userinfo, error) {
			return &authorizer.Userinfo{
				UserId: testUserId,
			}, nil
		},
	}

	_, err := NewService(logger, store, squareClient, authorizer).Add(context.Background(), registration, testRedirectUrl, testAccessToken)
	if err != nil {
		t.Fatalf("error adding new registration: %v", err)
	}
}

func TestAddTeamCompetition(t *testing.T) {
	testAccessToken := "some.access.token"
	testUserId := "12356"

	registration := &Registration{
		FirstName:     "John",
		LastName:      "Doe",
		StreetAddress: "123 Any St.",
		City:          "New York",
		State:         "NY",
		ZipCode:       "12345",
		Email:         "John.Doe@example.com",
		HomeScene:     "Frim Fram",
		IsStudent:     true,
		PassType:      &NoPass{},
		MixAndMatch:   nil,
		SoloJazz:      false,
		TeamCompetition: &TeamCompetition{
			Name: "Super duper dancing troopers",
		},
		TShirt:  nil,
		Housing: &NoHousing{},
	}
	testRedirectUrl := "https://daytonswingsmackdown.com/registration-complete"
	expectedCheckoutUrl := "https://squareup.com/98734987349873498345"

	catalogObjectsIdx := -1

	squareClient := &mockSquareClient{
		ListCatalogFunc: func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
			return &mockListCatalogIterator{
				ValueFunc: func() *square.CatalogObject {
					return catalogObjects[catalogObjectsIdx]
				},
				ErrorFunc: func() error {
					return nil
				},
				NextFunc: func() bool {
					catalogObjectsIdx++
					return catalogObjectsIdx < len(catalogObjects)
				},
			}
		},
		ListLocationsFunc: func(ctx context.Context) ([]*square.Location, error) {
			return locations, nil
		},
		CreateCheckoutFunc: func(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error) {
			createdAt := time.Now()
			if len(order.Order.LineItems) != 1 {
				t.Fatalf("Found unexpected number of line items %d (expected 1)", len(order.Order.LineItems))
			}
			if order.Order.LineItems[0].CatalogObjectId != teamCompVariationId {
				t.Fatalf("Found incorrect item id %s in order, expected %s", order.Order.LineItems[0].CatalogObjectId, teamCompVariationId)
			}
			if order.Order.LineItems[0].Quantity != "1" {
				t.Fatalf("Ordering incorrect number of team comp tickets %s (expected 1)", order.Order.LineItems[0].Quantity)
			}
			return &square.Checkout{
				Id:                         "checkout id",
				CheckoutPageUrl:            expectedCheckoutUrl,
				AskForShippingAddress:      askForShippingAddress,
				MerchantSupportEmail:       merchantSupportEmail,
				PrePopulateBuyerEmail:      prePopulateBuyerEmail,
				PrePopulateShippingAddress: prePopulateShippingAddress,
				RedirectUrl:                redirectUrl,
				Order:                      order.Order,
				CreatedAt:                  &createdAt,
				AdditionalRecipients:       nil,
			}, nil
		},
	}

	store := &mockStore{
		AddRegistrationFunc: func(ctx context.Context, r *StoreRegistration) (string, error) {
			if !cmp.Equal(r.TeamCompetition, registration.TeamCompetition, cmpopts.IgnoreUnexported()) {
				t.Fatalf("found incorrect mix and match data %s, expected %s", spew.Sdump(r.TeamCompetition), spew.Sdump(registration.TeamCompetition))
			}
			return "store key", nil
		},
		DeleteRegistrationFunc: func(ctx context.Context, id string) error {
			t.Error("delete registration called")
			return nil
		},
	}

	logger := logrus.New()
	logger.SetOutput(&test_utility.ErrorWriter{t})
	logger.AddHook(&test_utility.ErrorHook{t})
	logger.SetLevel(logrus.TraceLevel)

	authorizer := &MockAuthorizer{
		UserinfoFunc: func(ctx context.Context, accessToken string) (*authorizer.Userinfo, error) {
			return &authorizer.Userinfo{
				UserId: testUserId,
			}, nil
		},
	}

	_, err := NewService(logger, store, squareClient, authorizer).Add(context.Background(), registration, testRedirectUrl, testAccessToken)
	if err != nil {
		t.Fatalf("error adding new registration: %v", err)
	}
}

func TestAddTShirt(t *testing.T) {
	testAccessToken := "some.access.token"
	testUserId := "12356"

	registration := &Registration{
		FirstName:       "John",
		LastName:        "Doe",
		StreetAddress:   "123 Any St.",
		City:            "New York",
		State:           "NY",
		ZipCode:         "12345",
		Email:           "John.Doe@example.com",
		HomeScene:       "Frim Fram",
		IsStudent:       true,
		PassType:        &NoPass{},
		MixAndMatch:     nil,
		SoloJazz:        false,
		TeamCompetition: nil,
		TShirt: &TShirt{
			Style: TShirtStyleBellaL,
		},
		Housing: &NoHousing{},
	}
	testRedirectUrl := "https://daytonswingsmackdown.com/registration-complete"
	expectedCheckoutUrl := "https://squareup.com/98734987349873498345"

	catalogObjectsIdx := -1

	squareClient := &mockSquareClient{
		ListCatalogFunc: func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
			return &mockListCatalogIterator{
				ValueFunc: func() *square.CatalogObject {
					return catalogObjects[catalogObjectsIdx]
				},
				ErrorFunc: func() error {
					return nil
				},
				NextFunc: func() bool {
					catalogObjectsIdx++
					return catalogObjectsIdx < len(catalogObjects)
				},
			}
		},
		ListLocationsFunc: func(ctx context.Context) ([]*square.Location, error) {
			return locations, nil
		},
		CreateCheckoutFunc: func(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error) {
			createdAt := time.Now()
			if len(order.Order.LineItems) != 1 {
				t.Fatalf("Found unexpected number of line items %d (expected 1)", len(order.Order.LineItems))
			}
			if order.Order.LineItems[0].CatalogObjectId != tShirtVariationId {
				t.Fatalf("Found incorrect item id %s in order, expected %s", order.Order.LineItems[0].CatalogObjectId, tShirtVariationId)
			}
			if order.Order.LineItems[0].Quantity != "1" {
				t.Fatalf("Ordering incorrect number of tshirts %s (expected 1)", order.Order.LineItems[0].Quantity)
			}
			return &square.Checkout{
				Id:                         "checkout id",
				CheckoutPageUrl:            expectedCheckoutUrl,
				AskForShippingAddress:      askForShippingAddress,
				MerchantSupportEmail:       merchantSupportEmail,
				PrePopulateBuyerEmail:      prePopulateBuyerEmail,
				PrePopulateShippingAddress: prePopulateShippingAddress,
				RedirectUrl:                redirectUrl,
				Order:                      order.Order,
				CreatedAt:                  &createdAt,
				AdditionalRecipients:       nil,
			}, nil
		},
	}

	store := &mockStore{
		AddRegistrationFunc: func(ctx context.Context, r *StoreRegistration) (string, error) {
			if !cmp.Equal(r.TShirt, registration.TShirt, cmpopts.IgnoreUnexported()) {
				t.Fatalf("found incorrect mix and match data %s, expected %s", spew.Sdump(r.TShirt), spew.Sdump(registration.TShirt))
			}
			return "store key", nil
		},
		DeleteRegistrationFunc: func(ctx context.Context, id string) error {
			t.Error("delete registration called")
			return nil
		},
	}

	logger := logrus.New()
	logger.SetOutput(&test_utility.ErrorWriter{t})
	logger.AddHook(&test_utility.ErrorHook{t})
	logger.SetLevel(logrus.TraceLevel)

	authorizer := &MockAuthorizer{
		UserinfoFunc: func(ctx context.Context, accessToken string) (*authorizer.Userinfo, error) {
			return &authorizer.Userinfo{
				UserId: testUserId,
			}, nil
		},
	}

	_, err := NewService(logger, store, squareClient, authorizer).Add(context.Background(), registration, testRedirectUrl, testAccessToken)
	if err != nil {
		t.Fatalf("error adding new registration: %v", err)
	}
}

func TestAddCatalogError(t *testing.T) {
	testAccessToken := "some.access.token"
	testUserId := "12356"

	registration := &Registration{
		FirstName:       "John",
		LastName:        "Doe",
		StreetAddress:   "123 Any St.",
		City:            "New York",
		State:           "NY",
		ZipCode:         "12345",
		Email:           "John.Doe@example.com",
		HomeScene:       "Frim Fram",
		IsStudent:       true,
		PassType:        &NoPass{},
		MixAndMatch:     nil,
		SoloJazz:        true,
		TeamCompetition: nil,
		TShirt:          nil,
		Housing:         &NoHousing{},
	}
	testRedirectUrl := "https://daytonswingsmackdown.com/registration-complete"

	squareClient := &mockSquareClient{
		ListCatalogFunc: func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
			return &mockListCatalogIterator{
				ValueFunc: func() *square.CatalogObject {
					return nil
				},
				ErrorFunc: func() error {
					return errors.New("some error")
				},
				NextFunc: func() bool {
					return false
				},
			}
		},
		ListLocationsFunc: func(ctx context.Context) ([]*square.Location, error) {
			return locations, nil
		},
		CreateCheckoutFunc: func(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error) {
			t.Fatalf("Create checkout called when we couldn't fetch catalog items")
			return nil, nil
		},
	}

	adds, deletes := 0, 0
	addKey := "add key"

	store := &mockStore{
		AddRegistrationFunc: func(ctx context.Context, r *StoreRegistration) (string, error) {
			adds++
			return addKey, nil
		},
		DeleteRegistrationFunc: func(ctx context.Context, id string) error {
			if id != addKey {
				t.Fatalf("deleting wrong registration item")
			}
			deletes++
			return nil
		},
	}

	logger := logrus.New()
	logger.SetOutput(&test_utility.ErrorWriter{t})
	logger.SetLevel(logrus.TraceLevel)

	authorizer := &MockAuthorizer{
		UserinfoFunc: func(ctx context.Context, accessToken string) (*authorizer.Userinfo, error) {
			return &authorizer.Userinfo{
				UserId: testUserId,
			}, nil
		},
	}

	_, err := NewService(logger, store, squareClient, authorizer).Add(context.Background(), registration, testRedirectUrl, testAccessToken)
	if err == nil {
		t.Fatalf("No error returned when square could not list catalog items")
	}
	if adds != deletes {
		t.Fatalf("different number of adds and deletes?")
	}
}

func TestAddLocationsError(t *testing.T) {
	testAccessToken := "some.access.token"
	testUserId := "12356"

	registration := &Registration{
		FirstName:       "John",
		LastName:        "Doe",
		StreetAddress:   "123 Any St.",
		City:            "New York",
		State:           "NY",
		ZipCode:         "12345",
		Email:           "John.Doe@example.com",
		HomeScene:       "Frim Fram",
		IsStudent:       true,
		PassType:        &NoPass{},
		MixAndMatch:     nil,
		SoloJazz:        true,
		TeamCompetition: nil,
		TShirt:          nil,
		Housing:         &NoHousing{},
	}
	testRedirectUrl := "https://daytonswingsmackdown.com/registration-complete"
	catalogObjectsIdx := -1

	squareClient := &mockSquareClient{
		ListCatalogFunc: func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
			return &mockListCatalogIterator{
				ValueFunc: func() *square.CatalogObject {
					return catalogObjects[catalogObjectsIdx]
				},
				ErrorFunc: func() error {
					return nil
				},
				NextFunc: func() bool {
					catalogObjectsIdx++
					return catalogObjectsIdx < len(catalogObjects)
				},
			}
		},
		ListLocationsFunc: func(ctx context.Context) ([]*square.Location, error) {
			return nil, errors.New("some error")
		},
		CreateCheckoutFunc: func(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error) {
			t.Fatalf("Create checkout called when we couldn't fetch catalog items")
			return nil, nil
		},
	}

	adds, deletes := 0, 0
	addKey := "add key"

	store := &mockStore{
		AddRegistrationFunc: func(ctx context.Context, r *StoreRegistration) (string, error) {
			adds++
			return addKey, nil
		},
		DeleteRegistrationFunc: func(ctx context.Context, id string) error {
			if id != addKey {
				t.Fatalf("deleting wrong registration item")
			}
			deletes++
			return nil
		},
	}

	logger := logrus.New()
	logger.SetOutput(&test_utility.ErrorWriter{t})
	logger.SetLevel(logrus.TraceLevel)

	authorizer := &MockAuthorizer{
		UserinfoFunc: func(ctx context.Context, accessToken string) (*authorizer.Userinfo, error) {
			return &authorizer.Userinfo{
				UserId: testUserId,
			}, nil
		},
	}

	_, err := NewService(logger, store, squareClient, authorizer).Add(context.Background(), registration, testRedirectUrl, testAccessToken)
	if err == nil {
		t.Fatalf("No error returned when square could not list catalog items")
	}
	if adds != deletes {
		t.Fatalf("different number of adds and deletes?")
	}
}

func TestAddCheckoutError(t *testing.T) {
	testAccessToken := "some.access.token"
	testUserId := "12356"

	registration := &Registration{
		FirstName:       "John",
		LastName:        "Doe",
		StreetAddress:   "123 Any St.",
		City:            "New York",
		State:           "NY",
		ZipCode:         "12345",
		Email:           "John.Doe@example.com",
		HomeScene:       "Frim Fram",
		IsStudent:       true,
		PassType:        &NoPass{},
		MixAndMatch:     nil,
		SoloJazz:        true,
		TeamCompetition: nil,
		TShirt:          nil,
		Housing:         &NoHousing{},
	}
	testRedirectUrl := "https://daytonswingsmackdown.com/registration-complete"
	catalogObjectsIdx := -1

	squareClient := &mockSquareClient{
		ListCatalogFunc: func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
			return &mockListCatalogIterator{
				ValueFunc: func() *square.CatalogObject {
					return catalogObjects[catalogObjectsIdx]
				},
				ErrorFunc: func() error {
					return nil
				},
				NextFunc: func() bool {
					catalogObjectsIdx++
					return catalogObjectsIdx < len(catalogObjects)
				},
			}
		},
		ListLocationsFunc: func(ctx context.Context) ([]*square.Location, error) {
			return locations, nil
		},
		CreateCheckoutFunc: func(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error) {
			return nil, errors.New("some error")
		},
	}

	adds, deletes := 0, 0
	addKey := "add key"

	store := &mockStore{
		AddRegistrationFunc: func(ctx context.Context, r *StoreRegistration) (string, error) {
			adds++
			return addKey, nil
		},
		DeleteRegistrationFunc: func(ctx context.Context, id string) error {
			if id != addKey {
				t.Fatalf("deleting wrong registration item")
			}
			deletes++
			return nil
		},
	}

	logger := logrus.New()
	logger.SetOutput(&test_utility.ErrorWriter{t})
	logger.SetLevel(logrus.TraceLevel)

	authorizer := &MockAuthorizer{
		UserinfoFunc: func(ctx context.Context, accessToken string) (*authorizer.Userinfo, error) {
			return &authorizer.Userinfo{
				UserId: testUserId,
			}, nil
		},
	}

	_, err := NewService(logger, store, squareClient, authorizer).Add(context.Background(), registration, testRedirectUrl, testAccessToken)
	if err == nil {
		t.Fatalf("No error returned when square could not create checkout")
	}
	if adds != deletes {
		t.Fatalf("different number of adds and deletes?")
	}
}

func TestAddAddRegistrationError(t *testing.T) {
	testAccessToken := "some.access.token"
	testUserId := "12356"

	registration := &Registration{
		FirstName:       "John",
		LastName:        "Doe",
		StreetAddress:   "123 Any St.",
		City:            "New York",
		State:           "NY",
		ZipCode:         "12345",
		Email:           "John.Doe@example.com",
		HomeScene:       "Frim Fram",
		IsStudent:       true,
		PassType:        &NoPass{},
		MixAndMatch:     nil,
		SoloJazz:        true,
		TeamCompetition: nil,
		TShirt:          nil,
		Housing:         &NoHousing{},
	}
	testRedirectUrl := "https://daytonswingsmackdown.com/registration-complete"
	catalogObjectsIdx := -1

	squareClient := &mockSquareClient{
		ListCatalogFunc: func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
			return &mockListCatalogIterator{
				ValueFunc: func() *square.CatalogObject {
					return catalogObjects[catalogObjectsIdx]
				},
				ErrorFunc: func() error {
					return nil
				},
				NextFunc: func() bool {
					catalogObjectsIdx++
					return catalogObjectsIdx < len(catalogObjects)
				},
			}
		},
		ListLocationsFunc: func(ctx context.Context) ([]*square.Location, error) {
			return locations, nil
		},
		CreateCheckoutFunc: func(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error) {
			t.Fatalf("created checkout when registration could not be stored in db")
			return nil, nil
		},
	}

	store := &mockStore{
		AddRegistrationFunc: func(ctx context.Context, r *StoreRegistration) (string, error) {
			return "", errors.New("some error")
		},
		DeleteRegistrationFunc: func(ctx context.Context, id string) error {
			t.Fatalf("deleting registration that wasn't inserted")
			return nil
		},
	}

	logger := logrus.New()
	logger.SetOutput(&test_utility.ErrorWriter{t})
	logger.SetLevel(logrus.TraceLevel)

	authorizer := &MockAuthorizer{
		UserinfoFunc: func(ctx context.Context, accessToken string) (*authorizer.Userinfo, error) {
			return &authorizer.Userinfo{
				UserId: testUserId,
			}, nil
		},
	}

	_, err := NewService(logger, store, squareClient, authorizer).Add(context.Background(), registration, testRedirectUrl, testAccessToken)
	if err == nil {
		t.Fatalf("No error returned when square could not create checkout")
	}
}

func TestAddRegistrationNoPayments(t *testing.T) {
	testAccessToken := "some.access.token"
	testUserId := "12356"

	registration := &Registration{
		FirstName:       "John",
		LastName:        "Doe",
		StreetAddress:   "123 Any St.",
		City:            "New York",
		State:           "NY",
		ZipCode:         "12345",
		Email:           "John.Doe@example.com",
		HomeScene:       "Frim Fram",
		IsStudent:       true,
		PassType:        &NoPass{},
		MixAndMatch:     nil,
		SoloJazz:        false,
		TeamCompetition: nil,
		TShirt:          nil,
		Housing:         &NoHousing{},
	}
	testRedirectUrl := "https://daytonswingsmackdown.com/registration-complete"

	squareClient := &mockSquareClient{
		ListCatalogFunc: func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
			t.Fatalf("list catalog call made unnecessarily")
			return &mockListCatalogIterator{
				ValueFunc: func() *square.CatalogObject {
					return nil
				},
				ErrorFunc: func() error {
					return nil
				},
				NextFunc: func() bool {
					return false
				},
			}
		},
		ListLocationsFunc: func(ctx context.Context) ([]*square.Location, error) {
			t.Fatalf("list locations call made unnecessarily")
			return nil, nil
		},
		CreateCheckoutFunc: func(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error) {
			t.Fatalf("create checkout call made unnecessarily")
			return nil, nil
		},
	}

	store := &mockStore{
		AddRegistrationFunc: func(ctx context.Context, r *StoreRegistration) (string, error) {
			return "key", nil
		},
		DeleteRegistrationFunc: func(ctx context.Context, id string) error {
			t.Fatalf("deleting registration?")
			return nil
		},
	}

	logger := logrus.New()
	logger.SetOutput(&test_utility.ErrorWriter{t})
	logger.SetLevel(logrus.TraceLevel)
	logger.AddHook(&test_utility.ErrorHook{t})

	authorizer := &MockAuthorizer{
		UserinfoFunc: func(ctx context.Context, accessToken string) (*authorizer.Userinfo, error) {
			return &authorizer.Userinfo{
				UserId: testUserId,
			}, nil
		},
	}

	url, err := NewService(logger, store, squareClient, authorizer).Add(context.Background(), registration, testRedirectUrl, testAccessToken)
	if err != nil {
		t.Fatalf("error returned when creating registration: %v", err)
	}
	if url != testRedirectUrl {
		t.Fatalf("found redirect url %s, expected %s", url, testRedirectUrl)
	}
}

func TestAddRegistrationNoUserId(t *testing.T) {
	registration := &Registration{
		FirstName:       "John",
		LastName:        "Doe",
		StreetAddress:   "123 Any St.",
		City:            "New York",
		State:           "NY",
		ZipCode:         "12345",
		Email:           "John.Doe@example.com",
		HomeScene:       "Frim Fram",
		IsStudent:       true,
		PassType:        &NoPass{},
		MixAndMatch:     nil,
		SoloJazz:        true,
		TeamCompetition: nil,
		TShirt:          nil,
		Housing:         &NoHousing{},
	}
	testRedirectUrl := "https://daytonswingsmackdown.com/registration-complete"

	catalogObjectsIdx := -1

	squareClient := &mockSquareClient{
		ListCatalogFunc: func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
			return &mockListCatalogIterator{
				ValueFunc: func() *square.CatalogObject {
					return catalogObjects[catalogObjectsIdx]
				},
				ErrorFunc: func() error {
					return nil
				},
				NextFunc: func() bool {
					catalogObjectsIdx++
					return catalogObjectsIdx < len(catalogObjects)
				},
			}
		},
		ListLocationsFunc: func(ctx context.Context) ([]*square.Location, error) {
			return locations, nil
		},
		CreateCheckoutFunc: func(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error) {
			createdAt := time.Now()
			return &square.Checkout{
				Id:                         "checkout id",
				CheckoutPageUrl:            "some url",
				AskForShippingAddress:      askForShippingAddress,
				MerchantSupportEmail:       merchantSupportEmail,
				PrePopulateBuyerEmail:      prePopulateBuyerEmail,
				PrePopulateShippingAddress: prePopulateShippingAddress,
				RedirectUrl:                redirectUrl,
				Order:                      order.Order,
				CreatedAt:                  &createdAt,
				AdditionalRecipients:       nil,
			}, nil
		},
	}

	store := &mockStore{
		AddRegistrationFunc: func(ctx context.Context, r *StoreRegistration) (string, error) {
			if r.UserId != "" {
				t.Fatalf("expected no user id when no access token given, found %s", r.UserId)
			}
			return "key", nil
		},
		DeleteRegistrationFunc: func(ctx context.Context, id string) error {
			t.Fatalf("deleting registration?")
			return nil
		},
	}

	logger := logrus.New()
	logger.SetOutput(&test_utility.ErrorWriter{t})
	logger.SetLevel(logrus.TraceLevel)
	logger.AddHook(&test_utility.ErrorHook{t})

	authorizer := &MockAuthorizer{
		UserinfoFunc: func(ctx context.Context, accessToken string) (*authorizer.Userinfo, error) {
			t.Fatalf("found call to userinfo endpoint when no access token given")
			return nil, nil
		},
	}

	_, err := NewService(logger, store, squareClient, authorizer).Add(context.Background(), registration, testRedirectUrl, "")
	if err != nil {
		t.Fatalf("error returned when creating registration: %v", err)
	}
}

func TestAddRegistrationUserInfoError(t *testing.T) {
	registration := &Registration{
		FirstName:       "John",
		LastName:        "Doe",
		StreetAddress:   "123 Any St.",
		City:            "New York",
		State:           "NY",
		ZipCode:         "12345",
		Email:           "John.Doe@example.com",
		HomeScene:       "Frim Fram",
		IsStudent:       true,
		PassType:        &NoPass{},
		MixAndMatch:     nil,
		SoloJazz:        true,
		TeamCompetition: nil,
		TShirt:          nil,
		Housing:         &NoHousing{},
	}
	testRedirectUrl := "https://daytonswingsmackdown.com/registration-complete"

	catalogObjectsIdx := -1

	squareClient := &mockSquareClient{
		ListCatalogFunc: func(ctx context.Context, types []square.CatalogObjectType) square.ListCatalogIterator {
			return &mockListCatalogIterator{
				ValueFunc: func() *square.CatalogObject {
					return catalogObjects[catalogObjectsIdx]
				},
				ErrorFunc: func() error {
					return nil
				},
				NextFunc: func() bool {
					catalogObjectsIdx++
					return catalogObjectsIdx < len(catalogObjects)
				},
			}
		},
		ListLocationsFunc: func(ctx context.Context) ([]*square.Location, error) {
			return locations, nil
		},
		CreateCheckoutFunc: func(ctx context.Context, locationId, idempotencyKey string, order *square.CreateOrderRequest, askForShippingAddress bool, merchantSupportEmail, prePopulateBuyerEmail string, prePopulateShippingAddress *square.Address, redirectUrl string, additionalRecipients []*square.ChargeRequestAdditionalRecipient, note string) (*square.Checkout, error) {
			t.Fatalf("expected no checkout created when userinfo could not be attained")
			return nil, nil
		},
	}

	store := &mockStore{
		AddRegistrationFunc: func(ctx context.Context, r *StoreRegistration) (string, error) {
			t.Fatalf("expected no registraiton data created when userinfo could not be attained")
			return "", nil
		},
		DeleteRegistrationFunc: func(ctx context.Context, id string) error {
			t.Fatalf("deleting registration?")
			return nil
		},
	}

	logger := logrus.New()
	logger.SetOutput(&test_utility.ErrorWriter{t})
	logger.SetLevel(logrus.TraceLevel)
	logger.AddHook(&test_utility.ErrorHook{t})

	authorizer := &MockAuthorizer{
		UserinfoFunc: func(ctx context.Context, accessToken string) (*authorizer.Userinfo, error) {
			return nil, errors.New("some error")
		},
	}

	_, err := NewService(logger, store, squareClient, authorizer).Add(context.Background(), registration, testRedirectUrl, "some.access.token")
	if err == nil {
		t.Fatalf("expected error to be returned when userinfo endpoint fails")
	}
}
