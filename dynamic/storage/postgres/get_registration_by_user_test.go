package postgres

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestGetRegistrationsByUser(t *testing.T) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DSS_PG_URL"))
	if err != nil {
		t.Fatalf("error connecting to database for test: %v", err)
	}
	defer pool.Close()

	store := NewStore(pool)
	defer func() {
		_, err := pool.Exec(context.Background(), "DELETE FROM "+registrationTable)
		if err != nil {
			t.Fatalf("error cleaning up after test: %v", err)
		}
	}()

	tests := []struct {
		name         string
		registration *storage.Registration
	}{
		{
			name: "all items",
			registration: &storage.Registration{
				ID:            "id-1",
				FirstName:     "Jimbo",
				LastName:      "Brown",
				StreetAddress: "123 Dirt",
				City:          "Nowhere",
				State:         "Missouri",
				ZipCode:       "12345",
				Email:         "birds@flyfast.com",
				HomeScene:     "Missouri",
				IsStudent:     true,
				PassType: &storage.WeekendPass{
					Tier:  storage.Tier1,
					Level: storage.Level1,
				},
				MixAndMatch: &storage.MixAndMatch{
					Role: storage.MixAndMatchRoleFollower,
				},
				SoloJazz: true,
				TeamCompetition: &storage.TeamCompetition{
					Name: "The dirty dirts",
				},
				TShirt: &storage.TShirt{
					Style: storage.TShirtStyleUnisex3XL,
				},
				Housing: &storage.ProvideHousing{
					Pets:     "a large bulldog",
					Quantity: 7,
					Details:  "dogs are great",
				},
				UserID:        "1234354",
				OrderIDs:      []string{"12352q35234", "98709812087123"},
				DiscountCodes: []string{"doe", "a deer"},
			},
		},
		{
			name: "dance only",
			registration: &storage.Registration{
				ID:        "id-2",
				FirstName: "Jimbo",
				LastName:  "Brown",
				Email:     "birds@flyfast.com",
				PassType:  &storage.DanceOnlyPass{},
				Housing: &storage.RequireHousing{
					PetAllergies: "cats",
					Details:      "dogs are great",
				},
				UserID:        "1234355634",
				OrderIDs:      []string{},
				DiscountCodes: []string{},
			},
		},
		{
			name: "nothing",
			registration: &storage.Registration{
				ID:        "id-3",
				FirstName: "Jimbo",
				LastName:  "Brown",
				Email:     "birds@flyfast.com",
				PassType:  &storage.NoPass{},
				Housing:   &storage.NoHousing{},
				UserID:    "12sdfa34355634",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			id, err := store.AddRegistration(context.Background(), test.registration)
			if err != nil {
				t.Fatalf("unexpected error adding registration: %v", err)
			}

			testRegistration, err := store.GetRegistrationsByUser(context.Background(), test.registration.UserID)
			if err != nil {
				t.Fatalf("unexpected error getting registration: %v", err)
			}

			if len(testRegistration) != 1 {
				t.Fatalf("expected one registration, found %d", len(testRegistration))
			}

			if testRegistration[0].ID != id {
				t.Fatalf("expected registration id %s, found %s", id, testRegistration[0].ID)
			}
			if testRegistration[0].FirstName != test.registration.FirstName {
				t.Fatalf("expected registration first name %s, found %s", test.registration.FirstName, testRegistration[0].FirstName)
			}
			if testRegistration[0].LastName != test.registration.LastName {
				t.Fatalf("expected registration last name %s, found %s", test.registration.LastName, testRegistration[0].LastName)
			}
			if testRegistration[0].StreetAddress != test.registration.StreetAddress {
				t.Fatalf("expected registration street address %s, found %s", test.registration.StreetAddress, testRegistration[0].StreetAddress)
			}
			if testRegistration[0].City != test.registration.City {
				t.Fatalf("expected registration city %s, found %s", test.registration.City, testRegistration[0].City)
			}
			if testRegistration[0].State != test.registration.State {
				t.Fatalf("expected registration state %s, found %s", test.registration.State, testRegistration[0].State)
			}
			if testRegistration[0].ZipCode != test.registration.ZipCode {
				t.Fatalf("expected registration zip code %s, found %s", test.registration.ZipCode, testRegistration[0].ZipCode)
			}
			if testRegistration[0].Email != test.registration.Email {
				t.Fatalf("expected registration email %s, found %s", test.registration.Email, testRegistration[0].Email)
			}
			if testRegistration[0].HomeScene != test.registration.HomeScene {
				t.Fatalf("expected registration home scene %s, found %s", test.registration.HomeScene, testRegistration[0].HomeScene)
			}
			if testRegistration[0].IsStudent != test.registration.IsStudent {
				t.Fatalf("expected registration student status %v, found %v", test.registration.IsStudent, testRegistration[0].IsStudent)
			}
			if testRegistration[0].SoloJazz != test.registration.SoloJazz {
				t.Fatalf("expected registration solo jazz status %v, found %v", test.registration.SoloJazz, testRegistration[0].SoloJazz)
			}
			if testRegistration[0].UserID != test.registration.UserID {
				t.Fatalf("expected registration user id %s, found %s", test.registration.UserID, testRegistration[0].UserID)
			}
			if !reflect.DeepEqual(testRegistration[0].PassType, test.registration.PassType) {
				t.Fatalf("expected registration pass type %#v, found %#v", test.registration.PassType, testRegistration[0].PassType)
			}
			if !reflect.DeepEqual(testRegistration[0].MixAndMatch, test.registration.MixAndMatch) {
				t.Fatalf("expected registration mix and match %#v, found %#v", test.registration.MixAndMatch, testRegistration[0].MixAndMatch)
			}
			if !reflect.DeepEqual(testRegistration[0].TeamCompetition, test.registration.TeamCompetition) {
				t.Fatalf("expected registration team competition %#v, found %#v", test.registration.TeamCompetition, testRegistration[0].TeamCompetition)
			}
			if !reflect.DeepEqual(testRegistration[0].TShirt, test.registration.TShirt) {
				t.Fatalf("expected registration tshirt %#v, found %#v", test.registration.TShirt, testRegistration[0].TShirt)
			}
			if !reflect.DeepEqual(testRegistration[0].Housing, test.registration.Housing) {
				t.Fatalf("expected registration housing %#v, found %#v", test.registration.Housing, testRegistration[0].Housing)
			}
			if len(testRegistration[0].OrderIDs) != len(test.registration.OrderIDs) {
				t.Fatalf("expected number of registration order ids %d, found %d", len(test.registration.OrderIDs), len(testRegistration[0].OrderIDs))
			}
			if len(testRegistration[0].OrderIDs) > 0 && !reflect.DeepEqual(testRegistration[0].OrderIDs, test.registration.OrderIDs) {
				t.Fatalf("expected registration order ids %#v, found %#v", test.registration.OrderIDs, testRegistration[0].OrderIDs)
			}
			if len(testRegistration[0].DiscountCodes) != len(test.registration.DiscountCodes) {
				t.Fatalf("expected number of registration discount codes %d, found %d", len(test.registration.DiscountCodes), len(testRegistration[0].DiscountCodes))
			}
			if len(testRegistration[0].DiscountCodes) > 0 && !reflect.DeepEqual(testRegistration[0].DiscountCodes, test.registration.DiscountCodes) {
				t.Fatalf("expected registration discount codes %#v, found %#v", test.registration.DiscountCodes, testRegistration[0].DiscountCodes)
			}
		})
	}

}

func TestGetRegistrationsByUserNone(t *testing.T) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DSS_PG_URL"))
	if err != nil {
		t.Fatalf("error connecting to database for test: %v", err)
	}
	defer pool.Close()

	store := NewStore(pool)

	registrations, err := store.GetRegistrationsByUser(context.Background(), "me")
	if err != nil {
		t.Fatalf("Unexpected error when getting registration")
	}
	if len(registrations) > 0 {
		t.Fatalf("expected no registrations, found %d", len(registrations))
	}
}
