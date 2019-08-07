package datastore

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/datastore"
	"github.com/gofrs/uuid"
)

func TestMarkRegistrationPaid(t *testing.T) {
	if os.Getenv("DATASTORE_EMULATOR_HOST") == "" {
		t.Skip()
	}

	store, err := datastore.NewClient(context.Background(), os.Getenv("DATASTORE_PROJECT_ID"))
	if err != nil {
		t.Fatalf("could not connect to datastore emulator: %v", err)
	}

	referenceId, err := uuid.NewV4()
	if err != nil {
		t.Fatalf("error generating reference id for test: %v", err)
	}

	registration := &registrationEntity{
		FirstName:          "John",
		LastName:           "Doe",
		StreetAddress:      "1234 Any St",
		City:               "New York",
		State:              "New York",
		ZipCode:            "12345",
		Email:              "John.Doe@example.com",
		HomeScene:          "Frim Fram",
		IsStudent:          true,
		SoloJazz:           true,
		HousingRequest:     noHousing,
		WantsTShirt:        false,
		HasTeamCompetition: false,
		HasMixAndMatch:     false,
		WeekendPass:        danceOnlyPass,
		ReferenceId:        referenceId.String(),
		Paid:               false,
	}

	registrationKey := datastore.IncompleteKey(registrationKind, nil)
	completeKey, err := store.Put(context.Background(), registrationKey, registration)
	if err != nil {
		t.Fatalf("error inserting test registration into database: %v", err)
	}
	defer func() {
		err := store.Delete(context.Background(), completeKey)
		if err != nil {
			t.Log("WARNING:  Information leaked into database during test (test not cleaned up)")
			t.Fatalf("error cleaning up test item from store: %v", err)
		}
	}()

	err = NewDatastore(store).MarkRegistrationPaid(context.Background(), referenceId)
	if err != nil {
		t.Fatalf("error marking registration as paid: %v", err)
	}

	newRegistration := registrationEntity{}
	err = store.Get(context.Background(), completeKey, &newRegistration)
	if err != nil {
		t.Fatalf("error finding test item in datastore: %v", err)
	}
	if !newRegistration.Paid {
		t.Fatalf("could not mark registration as paid")
	}
}

func TestMarkRegistrationPaidNotExist(t *testing.T) {
	if os.Getenv("DATASTORE_EMULATOR_HOST") == "" {
		t.Skip()
	}

	store, err := datastore.NewClient(context.Background(), os.Getenv("DATASTORE_PROJECT_ID"))
	if err != nil {
		t.Fatalf("could not connect to datastore emulator: %v", err)
	}

	referenceId, err := uuid.NewV4()
	if err != nil {
		t.Fatalf("error generating reference id for test: %v", err)
	}

	err = NewDatastore(store).MarkRegistrationPaid(context.Background(), referenceId)
	if err == nil {
		t.Fatalf("expected error when marking registration that doesn't exist")
	}
}
