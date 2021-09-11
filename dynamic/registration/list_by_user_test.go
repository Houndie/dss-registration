package registration

import (
	"context"
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

func TestListByUser(t *testing.T) {
	registration1ID := "id1"
	registration2ID := "id2"
	expectedRegistrations := map[string]*Info{
		registration1ID: &Info{
			ID:        registration1ID,
			FirstName: "Joe",
			LastName:  "Dirt",
			Email:     "joedirt@verizon.net",
			CreatedAt: time.Now(),
			PassType:  &NoPass{},
		},
		registration2ID: &Info{
			ID:        registration2ID,
			FirstName: "John",
			LastName:  "Deer",
			Email:     "iliketractors@aol.com",
			CreatedAt: time.Now(),
			PassType:  &NoPass{},
		},
		"registration3": &Info{
			ID:        "registration3",
			FirstName: "Joe",
			LastName:  "Blow",
			Email:     "gameshows@mtv.com",
			CreatedAt: time.Now(),
			PassType:  &NoPass{},
		},
	}
	expectedToken := "iamauser"
	expectedUserID := "userid"
	expectedLocationID := "here"

	expectedOrders := []*objects.Order{
		{
			ID:    "order1",
			State: objects.OrderStateCompleted,
		},
		{
			ID:    "order2",
			State: objects.OrderStateCompleted,
		},
		{
			ID:    "order3",
			State: objects.OrderStateOpen,
		},
	}

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

	squareClient := &square.Client{
		Locations: &commontest.MockSquareLocationsClient{
			ListFunc: func(context.Context, *locations.ListRequest) (*locations.ListResponse, error) {
				return &locations.ListResponse{
					Locations: []*objects.Location{
						{
							ID: expectedLocationID,
						},
					},
				}, nil
			},
		},
		Orders: &commontest.MockSquareOrdersClient{
			BatchRetrieveFunc: commontest.OrdersFromSliceCheck(t, expectedLocationID, expectedOrders),
		},
	}

	store := &commontest.MockStore{
		GetRegistrationsByUserFunc: func(ctx context.Context, userID string) ([]*storage.Registration, error) {
			if userID != expectedUserID {
				t.Fatalf("expectedIncorrectUserID")
			}

			registrations := make([]*storage.Registration, 0, len(expectedRegistrations))
			for _, r := range expectedRegistrations {
				registrations = append(registrations, toStorageRegistration(r))
			}
			return registrations, nil
		},
	}

	service := NewService(true, false, logger, squareClient, commontest.CommonCatalogObjects().SquareData(), authorizer, store, &commontest.MockMailClient{}, nil)

	registrations, err := service.ListByUser(context.Background(), expectedToken)
	if err != nil {
		t.Fatalf("found unexpected error in call to SummaryByUser: %v", err)
	}
	for _, r := range registrations {
		expectedRegistration, ok := expectedRegistrations[r.ID]
		if !ok {
			t.Fatalf("registration with id %s not found in expected data", r.ID)
		}

		if !reflect.DeepEqual(r, expectedRegistration) {
			t.Fatalf("found registration %s, expected registration %s", spew.Sdump(r), spew.Sdump(expectedRegistration))
		}
	}
}
