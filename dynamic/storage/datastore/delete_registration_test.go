package datastore

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/datastore"
)

func TestDeleteRegistration(t *testing.T) {
	if os.Getenv("DATASTORE_EMULATOR_HOST") == "" {
		t.Skip()
	}

	store, err := datastore.NewClient(context.Background(), os.Getenv("DATASTORE_PROJECT_ID"))
	if err != nil {
		t.Fatalf("could not connect to datastore emulator: %v", err)
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
	}

	registrationKey := datastore.IncompleteKey(registrationKind, nil)
	completeKey, err := store.Put(context.Background(), registrationKey, registration)
	if err != nil {
		t.Fatalf("unable to insert test data into datastore: %v", err)
	}

	err = NewDatastore(store).DeleteRegistration(context.Background(), completeKey.Encode())
	if err != nil {
		t.Fatalf("error deleting registration from store")
	}

	foundRegistration := &registrationEntity{}
	err = store.Get(context.Background(), completeKey, foundRegistration)
	if err == nil {
		t.Fatal("expected no such entity error, didn't get error at all")
	}

	if err != datastore.ErrNoSuchEntity {
		t.Fatalf("expected no such entity error, found %v", err)
	}
}

func TestDeleteRegistrationInvalidKey(t *testing.T) {
	if os.Getenv("DATASTORE_EMULATOR_HOST") == "" {
		t.Skip()
	}

	store, err := datastore.NewClient(context.Background(), os.Getenv("DATASTORE_PROJECT_ID"))
	if err != nil {
		t.Fatalf("could not connect to datastore emulator: %v", err)
	}

	err = NewDatastore(store).DeleteRegistration(context.Background(), "intentionally_garbage_key")
	if err == nil {
		t.Fatal("expected error from store, found none")
	}
}
