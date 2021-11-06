package postgres

import (
	"context"
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestAddGetRegistration(t *testing.T) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DSS_TEST_POSTGRESURL"))
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
				SoloJazz: &storage.SoloJazz{},
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
				UserID:        "1234355634",
				OrderIDs:      []string{"12352q35234", "98709812087123"},
				DiscountCodes: []string{"doe", "a deer"},
			},
		},
		{
			name: "dance only",
			registration: &storage.Registration{
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
				FirstName: "Jimbo",
				LastName:  "Brown",
				Email:     "birds@flyfast.com",
				PassType:  &storage.NoPass{},
				Housing:   &storage.NoHousing{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			id, err := store.AddRegistration(context.Background(), test.registration)
			if err != nil {
				t.Fatalf("unexpected error adding registration: %v", err)
			}

			testRegistration, err := store.GetRegistration(context.Background(), id)
			if err != nil {
				t.Fatalf("unexpected error getting registration: %v", err)
			}

			if testRegistration.ID != id {
				t.Fatalf("expected registration id %s, found %s", id, testRegistration.ID)
			}
			if testRegistration.FirstName != test.registration.FirstName {
				t.Fatalf("expected registration first name %s, found %s", test.registration.FirstName, testRegistration.FirstName)
			}
			if testRegistration.LastName != test.registration.LastName {
				t.Fatalf("expected registration last name %s, found %s", test.registration.LastName, testRegistration.LastName)
			}
			if testRegistration.StreetAddress != test.registration.StreetAddress {
				t.Fatalf("expected registration street address %s, found %s", test.registration.StreetAddress, testRegistration.StreetAddress)
			}
			if testRegistration.City != test.registration.City {
				t.Fatalf("expected registration city %s, found %s", test.registration.City, testRegistration.City)
			}
			if testRegistration.State != test.registration.State {
				t.Fatalf("expected registration state %s, found %s", test.registration.State, testRegistration.State)
			}
			if testRegistration.ZipCode != test.registration.ZipCode {
				t.Fatalf("expected registration zip code %s, found %s", test.registration.ZipCode, testRegistration.ZipCode)
			}
			if testRegistration.Email != test.registration.Email {
				t.Fatalf("expected registration email %s, found %s", test.registration.Email, testRegistration.Email)
			}
			if testRegistration.HomeScene != test.registration.HomeScene {
				t.Fatalf("expected registration home scene %s, found %s", test.registration.HomeScene, testRegistration.HomeScene)
			}
			if testRegistration.IsStudent != test.registration.IsStudent {
				t.Fatalf("expected registration student status %v, found %v", test.registration.IsStudent, testRegistration.IsStudent)
			}
			if !reflect.DeepEqual(testRegistration.SoloJazz, test.registration.SoloJazz) {
				t.Fatalf("expected registration solo jazz status %v, found %v", test.registration.SoloJazz, testRegistration.SoloJazz)
			}
			if testRegistration.UserID != test.registration.UserID {
				t.Fatalf("expected registration user id %s, found %s", test.registration.UserID, testRegistration.UserID)
			}
			if !reflect.DeepEqual(testRegistration.PassType, test.registration.PassType) {
				t.Fatalf("expected registration pass type %#v, found %#v", test.registration.PassType, testRegistration.PassType)
			}
			if !reflect.DeepEqual(testRegistration.MixAndMatch, test.registration.MixAndMatch) {
				t.Fatalf("expected registration mix and match %#v, found %#v", test.registration.MixAndMatch, testRegistration.MixAndMatch)
			}
			if !reflect.DeepEqual(testRegistration.TeamCompetition, test.registration.TeamCompetition) {
				t.Fatalf("expected registration team competition %#v, found %#v", test.registration.TeamCompetition, testRegistration.TeamCompetition)
			}
			if !reflect.DeepEqual(testRegistration.TShirt, test.registration.TShirt) {
				t.Fatalf("expected registration tshirt %#v, found %#v", test.registration.TShirt, testRegistration.TShirt)
			}
			if !reflect.DeepEqual(testRegistration.Housing, test.registration.Housing) {
				t.Fatalf("expected registration housing %#v, found %#v", test.registration.Housing, testRegistration.Housing)
			}
			if len(testRegistration.OrderIDs) != len(test.registration.OrderIDs) {
				t.Fatalf("expected number of registration order ids %d, found %d", len(test.registration.OrderIDs), len(testRegistration.OrderIDs))
			}
			if len(testRegistration.OrderIDs) > 0 && !reflect.DeepEqual(testRegistration.OrderIDs, test.registration.OrderIDs) {
				t.Fatalf("expected registration order ids %#v, found %#v", test.registration.OrderIDs, testRegistration.OrderIDs)
			}
			if len(testRegistration.DiscountCodes) != len(test.registration.DiscountCodes) {
				t.Fatalf("expected number of registration discount codes %d, found %d", len(test.registration.DiscountCodes), len(testRegistration.DiscountCodes))
			}
			if len(testRegistration.DiscountCodes) > 0 && !reflect.DeepEqual(testRegistration.DiscountCodes, test.registration.DiscountCodes) {
				t.Fatalf("expected registration discount codes %#v, found %#v", test.registration.DiscountCodes, testRegistration.DiscountCodes)
			}
		})
	}

}

func TestGetRegistrationDoesntExist(t *testing.T) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DSS_TEST_POSTGRESURL"))
	if err != nil {
		t.Fatalf("error connecting to database for test: %v", err)
	}
	defer pool.Close()

	store := NewStore(pool)
	uuidid, err := uuid.NewV4()
	if err != nil {
		t.Fatalf("error generating uuid for test: %v", err)
	}
	tests := []struct {
		name string
		id   string
	}{
		{
			name: "uuid",
			id:   uuidid.String(),
		},
		{
			name: "random string",
			id:   "aisayiasfjoisaf",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := store.GetRegistration(context.Background(), test.id)
			if err == nil {
				t.Fatalf("expected error, found none")
			}
			serr := storage.ErrNoRegistrationForID{}
			if !errors.As(err, &serr) {
				t.Fatalf("expected no registration for id error, found: %v", err)
			}
			if serr.ID != test.id {
				t.Fatalf("found id %s, expected %s", serr.ID, test.id)
			}
		})
	}
}
