package registration

import (
	"context"
	"errors"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/Houndie/dss-registration/dynamic/square"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/test_utility"
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

	authorizer := &mockAuthorizer{
		UserinfoFunc: UserinfoFromIDCheck(t, expectedToken, expectedUserID),
	}
	co := commonCatalogObjects()
	expectedOrders := []*square.Order{
		{
			Id:    "order1",
			State: square.OrderStateCompleted,
			LineItems: []*square.OrderLineItem{
				{CatalogObjectId: co.weekendPassID[storage.Tier2]},
				{CatalogObjectId: co.soloJazzID},
				{CatalogObjectId: co.tShirtID},
			},
		},
	}

	userID := expectedUserID
	orderIDs := []string{expectedOrders[0].Id}

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
			Style: storage.TShirtStyleBellaM,
			Paid:  true,
		},
		Housing: &storage.ProvideHousing{
			Quantity: 11,
			Pets:     "Dogs",
			Details:  "whatever",
		},
		CreatedAt: time.Now(),
	}

	store := &mockStore{
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
				OrderIds:        orderIDs,
				UserId:          userID,
			}, nil
		},
	}

	client := &mockSquareClient{
		ListLocationsFunc: func(context.Context) ([]*square.Location, error) {
			return []*square.Location{{Id: expectedLocationID}}, nil
		},
		ListCatalogFunc:         listCatalogFuncFromSlice(co.catalog()),
		BatchRetrieveOrdersFunc: ordersFromSliceCheck(t, expectedLocationID, expectedOrders),
	}

	service := NewService(true, logger, client, authorizer, store, &mockMailClient{})
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

	authorizer := &mockAuthorizer{
		UserinfoFunc: UserinfoFromIDCheck(t, expectedToken, expectedUserID),
	}
	co := commonCatalogObjects()
	expectedOrders := []*square.Order{
		{
			Id:    "order1",
			State: square.OrderStateCompleted,
			LineItems: []*square.OrderLineItem{
				{CatalogObjectId: co.weekendPassID[storage.Tier2]},
				{CatalogObjectId: co.soloJazzID},
				{CatalogObjectId: co.tShirtID},
			},
		},
	}

	orderIDs := []string{expectedOrders[0].Id}

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
			Style: storage.TShirtStyleBellaM,
			Paid:  true,
		},
		Housing: &storage.ProvideHousing{
			Quantity: 11,
			Pets:     "Dogs",
			Details:  "whatever",
		},
		CreatedAt: time.Now(),
	}

	client := &mockSquareClient{
		ListLocationsFunc: func(context.Context) ([]*square.Location, error) {
			return []*square.Location{{Id: expectedLocationID}}, nil
		},
		ListCatalogFunc:         listCatalogFuncFromSlice(co.catalog()),
		BatchRetrieveOrdersFunc: ordersFromSliceCheck(t, expectedLocationID, expectedOrders),
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
			store := &mockStore{
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
						OrderIds:        orderIDs,
						UserId:          test.userID,
					}, nil
				},
			}

			service := NewService(true, logger, client, authorizer, store, &mockMailClient{})
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
