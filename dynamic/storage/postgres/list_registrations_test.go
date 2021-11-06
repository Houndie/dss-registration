package postgres

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/jackc/pgx/v4/pgxpool"
)

func TestRegistrations(t *testing.T) {
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

	registrations := []*storage.Registration{
		&storage.Registration{
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
			UserID:        "1234354",
			OrderIDs:      []string{"12352q35234", "98709812087123"},
			DiscountCodes: []string{"doe", "a deer"},
		},
		&storage.Registration{
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
		&storage.Registration{
			FirstName: "Jimbo",
			LastName:  "Brown",
			Email:     "birds@flyfast.com",
			PassType:  &storage.NoPass{},
			Housing:   &storage.NoHousing{},
			UserID:    "12sdfa34355634",
		},
	}

	for _, r := range registrations {
		r.ID, err = store.AddRegistration(context.Background(), r)
		if err != nil {
			t.Fatalf("unexpected error adding registration: %v", err)
		}
	}

	testRegistrations, err := store.ListRegistrations(context.Background())
	if err != nil {
		t.Fatalf("unexpected error getting registration: %v", err)
	}

	if len(testRegistrations) != len(registrations) {
		t.Fatal(len(testRegistrations))
	}

	for _, testRegistration := range testRegistrations {

		var controlRegistration *storage.Registration
		for _, r := range registrations {
			if testRegistration.ID == r.ID {
				controlRegistration = r
				break
			}
		}

		if controlRegistration == nil {
			t.Fatal(testRegistration.ID) // No registration with id found
		}

		if testRegistration.ID != controlRegistration.ID {
			t.Fatalf("expected registration id %s, found %s", controlRegistration.ID, testRegistration.ID)
		}
		if testRegistration.FirstName != controlRegistration.FirstName {
			t.Fatalf("expected registration first name %s, found %s", controlRegistration.FirstName, testRegistration.FirstName)
		}
		if testRegistration.LastName != controlRegistration.LastName {
			t.Fatalf("expected registration last name %s, found %s", controlRegistration.LastName, testRegistration.LastName)
		}
		if testRegistration.StreetAddress != controlRegistration.StreetAddress {
			t.Fatalf("expected registration street address %s, found %s", controlRegistration.StreetAddress, testRegistration.StreetAddress)
		}
		if testRegistration.City != controlRegistration.City {
			t.Fatalf("expected registration city %s, found %s", controlRegistration.City, testRegistration.City)
		}
		if testRegistration.State != controlRegistration.State {
			t.Fatalf("expected registration state %s, found %s", controlRegistration.State, testRegistration.State)
		}
		if testRegistration.ZipCode != controlRegistration.ZipCode {
			t.Fatalf("expected registration zip code %s, found %s", controlRegistration.ZipCode, testRegistration.ZipCode)
		}
		if testRegistration.Email != controlRegistration.Email {
			t.Fatalf("expected registration email %s, found %s", controlRegistration.Email, testRegistration.Email)
		}
		if testRegistration.HomeScene != controlRegistration.HomeScene {
			t.Fatalf("expected registration home scene %s, found %s", controlRegistration.HomeScene, testRegistration.HomeScene)
		}
		if testRegistration.IsStudent != controlRegistration.IsStudent {
			t.Fatalf("expected registration student status %v, found %v", controlRegistration.IsStudent, testRegistration.IsStudent)
		}
		if !reflect.DeepEqual(testRegistration.SoloJazz, controlRegistration.SoloJazz) {
			t.Fatalf("expected registration solo jazz status %v, found %v", controlRegistration.SoloJazz, testRegistration.SoloJazz)
		}
		if testRegistration.UserID != controlRegistration.UserID {
			t.Fatalf("expected registration user id %s, found %s", controlRegistration.UserID, testRegistration.UserID)
		}
		if !reflect.DeepEqual(testRegistration.PassType, controlRegistration.PassType) {
			t.Fatalf("expected registration pass type %#v, found %#v", controlRegistration.PassType, testRegistration.PassType)
		}
		if !reflect.DeepEqual(testRegistration.MixAndMatch, controlRegistration.MixAndMatch) {
			t.Fatalf("expected registration mix and match %#v, found %#v", controlRegistration.MixAndMatch, testRegistration.MixAndMatch)
		}
		if !reflect.DeepEqual(testRegistration.TeamCompetition, controlRegistration.TeamCompetition) {
			t.Fatalf("expected registration team competition %#v, found %#v", controlRegistration.TeamCompetition, testRegistration.TeamCompetition)
		}
		if !reflect.DeepEqual(testRegistration.TShirt, controlRegistration.TShirt) {
			t.Fatalf("expected registration tshirt %#v, found %#v", controlRegistration.TShirt, testRegistration.TShirt)
		}
		if !reflect.DeepEqual(testRegistration.Housing, controlRegistration.Housing) {
			t.Fatalf("expected registration housing %#v, found %#v", controlRegistration.Housing, testRegistration.Housing)
		}
		if len(testRegistration.OrderIDs) != len(controlRegistration.OrderIDs) {
			t.Fatalf("expected number of registration order ids %d, found %d", len(controlRegistration.OrderIDs), len(testRegistration.OrderIDs))
		}
		if len(testRegistration.OrderIDs) > 0 && !reflect.DeepEqual(testRegistration.OrderIDs, controlRegistration.OrderIDs) {
			t.Fatalf("expected registration order ids %#v, found %#v", controlRegistration.OrderIDs, testRegistration.OrderIDs)
		}
		if len(testRegistration.DiscountCodes) != len(controlRegistration.DiscountCodes) {
			t.Fatalf("expected number of registration discount codes %d, found %d", len(controlRegistration.DiscountCodes), len(testRegistration.DiscountCodes))
		}
		if len(testRegistration.DiscountCodes) > 0 && !reflect.DeepEqual(testRegistration.DiscountCodes, controlRegistration.DiscountCodes) {
			t.Fatalf("expected registration discount codes %#v, found %#v", controlRegistration.DiscountCodes, testRegistration.DiscountCodes)
		}
	}

}

func TestListRegistrationsNone(t *testing.T) {
	pool, err := pgxpool.Connect(context.Background(), os.Getenv("DSS_TEST_POSTGRESURL"))
	if err != nil {
		t.Fatalf("error connecting to database for test: %v", err)
	}
	defer pool.Close()

	store := NewStore(pool)

	registrations, err := store.ListRegistrations(context.Background())
	if err != nil {
		t.Fatalf("Unexpected error when getting registration")
	}
	if len(registrations) > 0 {
		t.Fatalf("expected no registrations, found %d", len(registrations))
	}
}
