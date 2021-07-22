package registration

import (
	"context"
	"errors"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/commontest"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/test_utility"
	"github.com/Houndie/square-go"
	"github.com/Houndie/square-go/locations"
	"github.com/Houndie/square-go/objects"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

func TestGet(t *testing.T) {
	expectedToken := "token"
	expectedRegistrationID := "registrationID"
	expectedUserID := "user id"
	expectedLocationID := "location id"

	logger := logrus.New()
	devnull, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatalf("error opening null: %v", err)
	}
	logger.SetOutput(devnull)
	logger.AddHook(&test_utility.ErrorHook{T: t})

	authorizer := &commontest.MockAuthorizer{
		GetUserinfoFunc: commontest.UserinfoFromIDCheck(t, expectedToken, []authorizer.Permission{}, expectedUserID, []authorizer.Permission{}),
	}
	co := commontest.CommonCatalogObjects()
	style := storage.TShirtStyleBellaM
	expectedOrders := []*objects.Order{
		{
			ID:    "order1",
			State: objects.OrderStateCompleted,
			LineItems: []*objects.OrderLineItem{
				{CatalogObjectID: co.WeekendPassID[storage.Tier2]},
				{CatalogObjectID: co.SoloJazzID},
				{CatalogObjectID: co.TShirtID[style]},
			},
		},
	}

	userID := expectedUserID
	orderIDs := []string{expectedOrders[0].ID}

	registration := &Info{
		ID:            expectedRegistrationID,
		FirstName:     "Tony",
		LastName:      "Vlachos",
		StreetAddress: "123 In A Tree",
		City:          "Fiji",
		State:         "Fiji",
		ZipCode:       "13245",
		Email:         "king@survivor.com",
		HomeScene:     "Fiji",
		IsStudent:     true,
		PassType: &WeekendPass{
			Tier:  storage.Tier2,
			Level: storage.Level1,
			Paid:  true,
		},
		MixAndMatch: &MixAndMatch{
			Role: storage.MixAndMatchRoleLeader,
		},
		SoloJazz: &SoloJazz{
			Paid: true,
		},
		TeamCompetition: &TeamCompetition{
			Name: "Cops R Us",
		},
		TShirt: &TShirt{
			Style: style,
			Paid:  true,
		},
		Housing: &storage.ProvideHousing{
			Quantity: 11,
			Pets:     "Dogs",
			Details:  "whatever",
		},
		CreatedAt: time.Now(),
	}

	store := &commontest.MockStore{
		GetRegistrationFunc: func(ctx context.Context, registrationID string) (*storage.Registration, error) {
			if registrationID != expectedRegistrationID {
				t.Fatalf("found unexpected registration id %s, expected %s", registrationID, expectedRegistrationID)
			}

			return &storage.Registration{
				ID:              registration.ID,
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
				CreatedAt:       registration.CreatedAt,
				OrderIDs:        orderIDs,
				UserID:          userID,
			}, nil
		},
	}

	client := &square.Client{
		Locations: &commontest.MockSquareLocationsClient{
			ListFunc: func(context.Context, *locations.ListRequest) (*locations.ListResponse, error) {
				return &locations.ListResponse{
					Locations: []*objects.Location{{ID: expectedLocationID}},
				}, nil
			},
		},
		Orders: &commontest.MockSquareOrdersClient{
			BatchRetrieveFunc: commontest.OrdersFromSliceCheck(t, expectedLocationID, expectedOrders),
		},
	}

	service := NewService(true, false, logger, client, commontest.CommonCatalogObjects().SquareData(), authorizer, store, &commontest.MockMailClient{})
	r, err := service.Get(context.Background(), expectedToken, expectedRegistrationID)
	if err != nil {
		t.Fatalf("error in get registration call: %v", err)
	}
	if !reflect.DeepEqual(r, registration) {
		t.Fatalf("found registration %s, expected registration %s", spew.Sdump(r), spew.Sdump(registration))
	}
}

func TestGetWrongUser(t *testing.T) {
	expectedToken := "token"
	expectedRegistrationID := "registrationID"
	expectedUserID := "user id"
	expectedLocationID := "location id"

	logger := logrus.New()
	devnull, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		t.Fatalf("error opening null: %v", err)
	}
	logger.SetOutput(devnull)
	logger.AddHook(&test_utility.ErrorHook{T: t})

	authorizer := &commontest.MockAuthorizer{
		GetUserinfoFunc: commontest.UserinfoFromIDCheck(t, expectedToken, []authorizer.Permission{}, expectedUserID, []authorizer.Permission{}),
	}
	style := storage.TShirtStyleBellaM
	co := commontest.CommonCatalogObjects()
	expectedOrders := []*objects.Order{
		{
			ID:    "order1",
			State: objects.OrderStateCompleted,
			LineItems: []*objects.OrderLineItem{
				{CatalogObjectID: co.WeekendPassID[storage.Tier2]},
				{CatalogObjectID: co.SoloJazzID},
				{CatalogObjectID: co.TShirtID[style]},
			},
		},
	}

	orderIDs := []string{expectedOrders[0].ID}

	registration := &Info{
		ID:            expectedRegistrationID,
		FirstName:     "Tony",
		LastName:      "Vlachos",
		StreetAddress: "123 In A Tree",
		City:          "Fiji",
		State:         "Fiji",
		ZipCode:       "13245",
		Email:         "king@survivor.com",
		HomeScene:     "Fiji",
		IsStudent:     true,
		PassType: &WeekendPass{
			Tier:  storage.Tier2,
			Level: storage.Level1,
			Paid:  true,
		},
		MixAndMatch: &MixAndMatch{
			Role: storage.MixAndMatchRoleLeader,
		},
		SoloJazz: &SoloJazz{
			Paid: true,
		},
		TeamCompetition: &TeamCompetition{
			Name: "Cops R Us",
		},
		TShirt: &TShirt{
			Style: style,
			Paid:  true,
		},
		Housing: &storage.ProvideHousing{
			Quantity: 11,
			Pets:     "Dogs",
			Details:  "whatever",
		},
		CreatedAt: time.Now(),
	}

	client := &square.Client{
		Locations: &commontest.MockSquareLocationsClient{
			ListFunc: func(context.Context, *locations.ListRequest) (*locations.ListResponse, error) {
				return &locations.ListResponse{
					[]*objects.Location{{ID: expectedLocationID}},
				}, nil
			},
		},
		Orders: &commontest.MockSquareOrdersClient{
			BatchRetrieveFunc: commontest.OrdersFromSliceCheck(t, expectedLocationID, expectedOrders),
		},
	}

	for _, test := range []struct {
		name              string
		registrationFound bool
		userID            string
	}{
		{
			name:              "wrong user",
			registrationFound: true,
			userID:            "some different user",
		},
		{
			name:              "no registration",
			registrationFound: false,
			userID:            expectedUserID,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			store := &commontest.MockStore{
				GetRegistrationFunc: func(ctx context.Context, registrationID string) (*storage.Registration, error) {
					if !test.registrationFound {
						return nil, storage.ErrNotFound{}
					}
					return &storage.Registration{
						ID:              registration.ID,
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
						CreatedAt:       registration.CreatedAt,
						OrderIDs:        orderIDs,
						UserID:          test.userID,
					}, nil
				},
			}

			service := NewService(true, false, logger, client, commontest.CommonCatalogObjects().SquareData(), authorizer, store, &commontest.MockMailClient{})
			_, err = service.Get(context.Background(), expectedToken, expectedRegistrationID)
			if err == nil {
				t.Fatalf("expected error, found none")
			}
			nferr := storage.ErrNotFound{}
			if !errors.As(err, &nferr) {
				t.Fatalf("expected not found error, found: %v", err)
			}
		})
	}
}
