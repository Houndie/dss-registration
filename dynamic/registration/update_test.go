package registration

import (
	"context"
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/commontest"
	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/Houndie/dss-registration/dynamic/test_utility"
	"github.com/Houndie/square-go"
	"github.com/Houndie/square-go/locations"
	"github.com/Houndie/square-go/objects"
	"github.com/Houndie/square-go/orders"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

func TestUpdateAdmin(t *testing.T) {
	token := "test.token"
	userID := "userid"
	testID := "id"

	tests := []struct {
		name               string
		updateRegistration *Info
		oldRegistration    *storage.Registration
		newRegistration    *storage.Registration
	}{
		{
			name: "mix_and_match",
			updateRegistration: &Info{
				ID:       testID,
				PassType: &NoPass{},
				MixAndMatch: &MixAndMatch{
					AdminPaymentOverride: true,
				},
			},
			oldRegistration: &storage.Registration{
				ID:       testID,
				PassType: &storage.NoPass{},
				MixAndMatch: &storage.MixAndMatch{
					ManuallyPaid: false,
				},
				UserID: userID,
			},
			newRegistration: &storage.Registration{
				ID:       testID,
				PassType: &storage.NoPass{},
				MixAndMatch: &storage.MixAndMatch{
					ManuallyPaid: true,
				},
				UserID: userID,
			},
		},
		{
			name: "full_weekend",
			updateRegistration: &Info{
				ID: testID,
				PassType: &WeekendPass{
					AdminPaymentOverride: true,
				},
			},
			oldRegistration: &storage.Registration{
				ID:       testID,
				PassType: &storage.WeekendPass{},
				UserID:   userID,
			},
			newRegistration: &storage.Registration{
				ID: testID,
				PassType: &storage.WeekendPass{
					ManuallyPaid: true,
				},
				UserID: userID,
			},
		},
		{
			name: "dance_only",
			updateRegistration: &Info{
				ID: testID,
				PassType: &DanceOnlyPass{
					AdminPaymentOverride: true,
				},
			},
			oldRegistration: &storage.Registration{
				ID:       testID,
				PassType: &storage.DanceOnlyPass{},
				UserID:   userID,
			},
			newRegistration: &storage.Registration{
				ID: testID,
				PassType: &storage.DanceOnlyPass{
					ManuallyPaid: true,
				},
				UserID: userID,
			},
		},
		{
			name: "solo_jazz",
			updateRegistration: &Info{
				ID:       testID,
				PassType: &NoPass{},
				SoloJazz: &SoloJazz{
					AdminPaymentOverride: true,
				},
			},
			oldRegistration: &storage.Registration{
				ID:       testID,
				PassType: &storage.NoPass{},
				SoloJazz: &storage.SoloJazz{},
				UserID:   userID,
			},
			newRegistration: &storage.Registration{
				ID:       testID,
				PassType: &storage.NoPass{},
				SoloJazz: &storage.SoloJazz{
					ManuallyPaid: true,
				},
				UserID: userID,
			},
		},
		{
			name: "team",
			updateRegistration: &Info{
				ID:       testID,
				PassType: &NoPass{},
				TeamCompetition: &TeamCompetition{
					AdminPaymentOverride: true,
				},
			},
			oldRegistration: &storage.Registration{
				ID:              testID,
				PassType:        &storage.NoPass{},
				TeamCompetition: &storage.TeamCompetition{},
				UserID:          userID,
			},
			newRegistration: &storage.Registration{
				ID:       testID,
				PassType: &storage.NoPass{},
				TeamCompetition: &storage.TeamCompetition{
					ManuallyPaid: true,
				},
				UserID: userID,
			},
		},
		{
			name: "tshirt",
			updateRegistration: &Info{
				ID:       testID,
				PassType: &NoPass{},
				TShirt: &TShirt{
					AdminPaymentOverride: true,
				},
			},
			oldRegistration: &storage.Registration{
				ID:       testID,
				PassType: &storage.NoPass{},
				TShirt:   &storage.TShirt{},
				UserID:   userID,
			},
			newRegistration: &storage.Registration{
				ID:       testID,
				PassType: &storage.NoPass{},
				TShirt: &storage.TShirt{
					ManuallyPaid: true,
				},
				UserID: userID,
			},
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			logger := logrus.New()
			devnull, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				t.Fatalf("error opening null: %v", err)
			}
			logger.SetOutput(devnull)
			logger.AddHook(&test_utility.ErrorHook{T: t})

			authorizer := &commontest.MockAuthorizer{
				GetUserinfoFunc: commontest.UserinfoFromIDCheck(t, token, []string{testPermissionConfig.Update}, "me", []string{testPermissionConfig.Update}),
			}

			store := &commontest.MockStore{
				GetRegistrationFunc: func(ctx context.Context, id string) (*storage.Registration, error) {
					if id != testID {
						t.Fatal(id)
					}

					return tt.oldRegistration, nil
				},
				UpdateRegistrationFunc: func(ctx context.Context, r *storage.Registration) error {
					if !reflect.DeepEqual(r, tt.newRegistration) {
						t.Fatalf(spew.Sdump(r))
					}

					return nil
				},
			}

			client := &square.Client{
				Locations: &commontest.MockSquareLocationsClient{
					ListFunc: func(context.Context, *locations.ListRequest) (*locations.ListResponse, error) {
						return &locations.ListResponse{
							Locations: []*objects.Location{{ID: "7"}},
						}, nil
					},
				},
				Orders: &commontest.MockSquareOrdersClient{
					BatchRetrieveFunc: func(context.Context, *orders.BatchRetrieveRequest) (*orders.BatchRetrieveResponse, error) {
						return &orders.BatchRetrieveResponse{
							Orders: []*objects.Order{},
						}, nil

					},
				},
			}

			service := NewService(true, true, logger, client, commontest.CommonCatalogObjects().SquareData(), authorizer, store, nil, nil, testPermissionConfig)

			_, err = service.Update(context.Background(), token, tt.updateRegistration)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestUpdateNoAdminOverride(t *testing.T) {
	token := "test.token"
	userID := "userid"
	testID := "id"

	tests := []struct {
		name               string
		updateRegistration *Info
		oldRegistration    *storage.Registration
	}{
		{
			name: "mix_and_match",
			updateRegistration: &Info{
				ID:       testID,
				PassType: &NoPass{},
				MixAndMatch: &MixAndMatch{
					AdminPaymentOverride: true,
				},
			},
			oldRegistration: &storage.Registration{
				ID:       testID,
				PassType: &storage.NoPass{},
				MixAndMatch: &storage.MixAndMatch{
					ManuallyPaid: false,
				},
				UserID: userID,
			},
		},
		{
			name: "full_weekend",
			updateRegistration: &Info{
				ID: testID,
				PassType: &WeekendPass{
					AdminPaymentOverride: true,
				},
			},
			oldRegistration: &storage.Registration{
				ID:       testID,
				PassType: &storage.WeekendPass{},
				UserID:   userID,
			},
		},
		{
			name: "dance_only",
			updateRegistration: &Info{
				ID: testID,
				PassType: &DanceOnlyPass{
					AdminPaymentOverride: true,
				},
			},
			oldRegistration: &storage.Registration{
				ID:       testID,
				PassType: &storage.DanceOnlyPass{},
				UserID:   userID,
			},
		},
		{
			name: "solo_jazz",
			updateRegistration: &Info{
				ID:       testID,
				PassType: &NoPass{},
				SoloJazz: &SoloJazz{
					AdminPaymentOverride: true,
				},
			},
			oldRegistration: &storage.Registration{
				ID:       testID,
				PassType: &storage.NoPass{},
				SoloJazz: &storage.SoloJazz{},
				UserID:   userID,
			},
		},
		{
			name: "team",
			updateRegistration: &Info{
				ID:       testID,
				PassType: &NoPass{},
				TeamCompetition: &TeamCompetition{
					AdminPaymentOverride: true,
				},
			},
			oldRegistration: &storage.Registration{
				ID:              testID,
				PassType:        &storage.NoPass{},
				TeamCompetition: &storage.TeamCompetition{},
				UserID:          userID,
			},
		},
		{
			name: "tshirt",
			updateRegistration: &Info{
				ID:       testID,
				PassType: &NoPass{},
				TShirt: &TShirt{
					AdminPaymentOverride: true,
				},
			},
			oldRegistration: &storage.Registration{
				ID:       testID,
				PassType: &storage.NoPass{},
				TShirt:   &storage.TShirt{},
				UserID:   userID,
			},
		},
	}

	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			logger := logrus.New()
			devnull, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				t.Fatalf("error opening null: %v", err)
			}
			logger.SetOutput(devnull)
			logger.AddHook(&test_utility.ErrorHook{T: t})

			authorizer := &commontest.MockAuthorizer{
				GetUserinfoFunc: commontest.UserinfoFromIDCheck(t, token, []string{testPermissionConfig.Update}, userID, []string{}),
			}

			store := &commontest.MockStore{
				GetRegistrationFunc: func(ctx context.Context, id string) (*storage.Registration, error) {
					if id != testID {
						t.Fatal(id)
					}

					return tt.oldRegistration, nil
				},
			}

			client := &square.Client{
				Locations: &commontest.MockSquareLocationsClient{
					ListFunc: func(context.Context, *locations.ListRequest) (*locations.ListResponse, error) {
						return &locations.ListResponse{
							Locations: []*objects.Location{{ID: "7"}},
						}, nil
					},
				},
				Orders: &commontest.MockSquareOrdersClient{
					BatchRetrieveFunc: func(context.Context, *orders.BatchRetrieveRequest) (*orders.BatchRetrieveResponse, error) {
						return &orders.BatchRetrieveResponse{
							Orders: []*objects.Order{},
						}, nil

					},
				},
			}

			service := NewService(true, true, logger, client, commontest.CommonCatalogObjects().SquareData(), authorizer, store, nil, nil, testPermissionConfig)

			_, err = service.Update(context.Background(), token, tt.updateRegistration)
			if err == nil {
				t.Fatal("expected error")
			}
			var e ErrHasAdminOverride
			if !errors.As(err, &e) {
				t.Fatal(err)
			}
		})
	}
}
