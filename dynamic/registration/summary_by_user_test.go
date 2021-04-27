package registration

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Houndie/dss-registration/dynamic/authorizer"
	"github.com/Houndie/dss-registration/dynamic/commontest"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/test_utility"
	"github.com/Houndie/square-go"
	"github.com/Houndie/square-go/locations"
	"github.com/Houndie/square-go/objects"
	"github.com/sirupsen/logrus"
)

func TestSummaryByUser(t *testing.T) {
	summary1ID := "summary1"
	summary2ID := "summary2"
	expectedSummaries := map[string]*Summary{
		summary1ID: &Summary{
			FirstName: "Joe",
			LastName:  "Dirt",
			Email:     "joedirt@verizon.net",
			CreatedAt: time.Now(),
			Paid:      true,
		},
		summary2ID: &Summary{
			FirstName: "John",
			LastName:  "Deer",
			Email:     "iliketractors@aol.com",
			CreatedAt: time.Now(),
			Paid:      false,
		},
		"summary3": &Summary{
			FirstName: "Joe",
			LastName:  "Blow",
			Email:     "gameshows@mtv.com",
			CreatedAt: time.Now(),
			Paid:      true,
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

	regToOrderMap := map[string][]string{
		summary1ID: []string{expectedOrders[0].ID},
		summary2ID: []string{expectedOrders[1].ID, expectedOrders[2].ID},
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

			registrations := make([]*storage.Registration, len(expectedSummaries))
			idx := 0
			for id, summary := range expectedSummaries {
				orderIDs := []string{}
				if o, ok := regToOrderMap[id]; ok {
					orderIDs = o
				}
				registrations[idx] = &storage.Registration{
					ID:        id,
					FirstName: summary.FirstName,
					LastName:  summary.LastName,
					Email:     summary.Email,
					CreatedAt: summary.CreatedAt,
					OrderIDs:  orderIDs,
				}
				idx++
			}
			return registrations, nil
		},
	}

	service := NewService(true, false, logger, squareClient, authorizer, store, &commontest.MockMailClient{})

	summaries, err := service.SummaryByUser(context.Background(), expectedToken)
	if err != nil {
		t.Fatalf("found unexpected error in call to SummaryByUser: %v", err)
	}
	for _, summary := range summaries {
		expectedSummary, ok := expectedSummaries[summary.ID]
		if !ok {
			t.Fatalf("summary with id %s not found in expected data", summary.ID)
		}

		if expectedSummary.FirstName != summary.FirstName {
			t.Fatalf("expected summary first name %s, found %s", expectedSummary.FirstName, summary.FirstName)
		}
		if expectedSummary.LastName != summary.LastName {
			t.Fatalf("expected summary last name %s, found %s", expectedSummary.LastName, summary.LastName)
		}
		if expectedSummary.Email != summary.Email {
			t.Fatalf("expected summary email %s, found %s", expectedSummary.Email, summary.Email)
		}
		if !expectedSummary.CreatedAt.Equal(summary.CreatedAt) {
			t.Fatalf("expected summary created at %v, found %v", expectedSummary.CreatedAt, summary.CreatedAt)
		}
		if expectedSummary.Paid != summary.Paid {
			t.Fatalf("expected summary paid %v, found %v", expectedSummary.Paid, summary.Paid)
		}
	}
}
