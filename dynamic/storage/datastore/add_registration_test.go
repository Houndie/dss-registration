package datastore

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/registration/add"
	"github.com/gofrs/uuid"
)

func TestAddRegistration1(t *testing.T) {
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

	registration := &add.StoreRegistration{
		FirstName:       "John",
		LastName:        "Doe",
		StreetAddress:   "123 Any Street",
		City:            "Pittsburgh",
		State:           "PA",
		ZipCode:         "12345",
		Email:           "John.Doe@example.com",
		HomeScene:       "Swing City",
		IsStudent:       false,
		PassType:        &add.NoPass{},
		MixAndMatch:     nil,
		SoloJazz:        false,
		TeamCompetition: nil,
		TShirt:          nil,
		Housing:         &add.NoHousing{},
		ReferenceId:     referenceId,
		Paid:            false,
	}

	key, err := NewDatastore(store).AddRegistration(context.Background(), registration)
	if err != nil {
		t.Fatalf("error inserting new registration into store: %v", err)
	}
	decodedKey, err := datastore.DecodeKey(key)
	if err != nil {
		t.Log("WARNING:  Information leaked into database during test (test not cleaned up)")
		t.Fatalf("could not decode key %s: %v", key, err)
	}
	defer func() {
		err := store.Delete(context.Background(), decodedKey)
		if err != nil {
			t.Log("WARNING:  Information leaked into database during test (test not cleaned up)")
			t.Fatalf("error cleaning up test item from store: %v", err)
		}
	}()

	entity := &registrationEntity{}
	err = store.Get(context.Background(), decodedKey, entity)
	if err != nil {
		t.Fatalf("error finding newly added registration in store: %v", err)
	}

	if registration.FirstName != entity.FirstName {
		t.Fatalf("found first name %s, expected %s", entity.FirstName, registration.FirstName)
	}

	if registration.LastName != entity.LastName {
		t.Fatalf("found last name %s, expected %s", entity.LastName, registration.LastName)
	}

	if registration.StreetAddress != entity.StreetAddress {
		t.Fatalf("found street address %s, expected %s", entity.StreetAddress, registration.StreetAddress)
	}

	if registration.City != entity.City {
		t.Fatalf("found city %s, expected %s", entity.City, registration.City)
	}

	if registration.State != entity.State {
		t.Fatalf("found state %s, expected %s", entity.State, registration.State)
	}

	if registration.ZipCode != entity.ZipCode {
		t.Fatalf("found zip code %s, expected %s", entity.ZipCode, registration.ZipCode)
	}

	if registration.Email != entity.Email {
		t.Fatalf("found email %s, expected %s", entity.Email, registration.Email)
	}

	if registration.HomeScene != entity.HomeScene {
		t.Fatalf("found home scene %s, expected %s", entity.HomeScene, registration.HomeScene)
	}

	if registration.IsStudent != entity.IsStudent {
		t.Fatalf("found student status %v, expected %v", entity.IsStudent, registration.IsStudent)
	}

	if registration.SoloJazz != entity.SoloJazz {
		t.Fatalf("found solo jazz comp %v, expected %v", entity.SoloJazz, registration.SoloJazz)
	}

	if entity.HousingRequest != noHousing {
		t.Fatalf("found housing request %s, expected %s", entity.HousingRequest, noHousing)
	}

	if entity.WantsTShirt != false {
		t.Fatalf("found tshirt status %v, expected false", entity.WantsTShirt)
	}

	if entity.HasTeamCompetition != false {
		t.Fatalf("found team competition status %v, expected false", entity.HasTeamCompetition)
	}

	if entity.HasMixAndMatch != false {
		t.Fatalf("found mix and match status %v, expected false", entity.HasMixAndMatch)
	}

	if entity.WeekendPass != noPass {
		t.Fatalf("found weekend pass status %s, expected %s", entity.WeekendPass, noPass)
	}

	if entity.ReferenceId != referenceId.String() {
		t.Fatalf("found reference id %s, expected %s", entity.ReferenceId, registration.ReferenceId)
	}

	if entity.Paid != registration.Paid {
		t.Fatalf("found paid status %v, expected %v", entity.Paid, registration.Paid)
	}
}

func TestAddRegistration2(t *testing.T) {
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

	role := add.MixAndMatchRoleLeader
	name := "Awesome team"
	style := add.TShirtStyleBellaM
	petAllergies := "some pet allergies"
	details := "plz help"
	registration := &add.StoreRegistration{
		FirstName:     "John",
		LastName:      "Doe",
		StreetAddress: "123 Any Street",
		City:          "Pittsburgh",
		State:         "PA",
		ZipCode:       "12345",
		Email:         "John.Doe@example.com",
		HomeScene:     "Swing City",
		IsStudent:     true,
		PassType:      &add.DanceOnlyPass{},
		MixAndMatch: &add.MixAndMatch{
			Role: role,
		},
		SoloJazz: true,
		TeamCompetition: &add.TeamCompetition{
			Name: name,
		},
		TShirt: &add.TShirt{
			Style: style,
		},
		Housing: &add.RequireHousing{
			PetAllergies: petAllergies,
			Details:      details,
		},
		ReferenceId: referenceId,
		Paid:        true,
	}

	key, err := NewDatastore(store).AddRegistration(context.Background(), registration)
	if err != nil {
		t.Fatalf("error inserting new registration into store: %v", err)
	}
	decodedKey, err := datastore.DecodeKey(key)
	if err != nil {
		t.Log("WARNING:  Information leaked into database during test (test not cleaned up)")
		t.Fatalf("could not decode key %s: %v", key, err)
	}
	defer func() {
		err := store.Delete(context.Background(), decodedKey)
		if err != nil {
			t.Log("WARNING:  Information leaked into database during test (test not cleaned up)")
			t.Fatalf("error cleaning up test item from store: %v", err)
		}
	}()

	entity := &registrationEntity{}
	err = store.Get(context.Background(), decodedKey, entity)
	if err != nil {
		t.Fatalf("error finding newly added registration in store: %v", err)
	}

	if registration.IsStudent != entity.IsStudent {
		t.Fatalf("found student status %v, expected %v", entity.IsStudent, registration.IsStudent)
	}

	if registration.SoloJazz != entity.SoloJazz {
		t.Fatalf("found solo jazz comp %v, expected %v", entity.SoloJazz, registration.SoloJazz)
	}

	if entity.HousingRequest != requiresHousing {
		t.Fatalf("found housing request %s, expected %s", entity.HousingRequest, requiresHousing)
	}

	if entity.RequireHousing.PetAllergies != petAllergies {
		t.Fatalf("found pet allergies %s, expected %s", entity.RequireHousing.PetAllergies, petAllergies)
	}

	if entity.RequireHousing.Details != details {
		t.Fatalf("found pet allergies %s, expected %s", entity.RequireHousing.Details, details)
	}

	if entity.WantsTShirt != true {
		t.Fatalf("found tshirt status %v, expected true", entity.WantsTShirt)
	}

	if entity.TShirtStyle != string(style) {
		t.Fatalf("found tshirt style %s, expected %s", entity.TShirtStyle, style)
	}

	if entity.HasTeamCompetition != true {
		t.Fatalf("found team competition status %v, expected true", entity.HasTeamCompetition)
	}

	if entity.TeamCompetitionName != name {
		t.Fatalf("found team competition name %s, expected %s", entity.TeamCompetitionName, name)
	}

	if entity.HasMixAndMatch != true {
		t.Fatalf("found mix and match status %v, expected true", entity.HasMixAndMatch)
	}

	if entity.MixAndMatchRole != string(role) {
		t.Fatalf("found mix and match role %s, expected %s", entity.MixAndMatchRole, role)
	}

	if entity.WeekendPass != danceOnlyPass {
		t.Fatalf("found weekend pass status %s, expected %s", entity.WeekendPass, danceOnlyPass)
	}

	if entity.Paid != registration.Paid {
		t.Fatalf("found paid status %v, expected %v", entity.Paid, registration.Paid)
	}
}

func TestAddRegistration3(t *testing.T) {
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

	pets := "some pets"
	details := "plz help"
	quantity := 7
	registration := &add.StoreRegistration{
		FirstName:     "John",
		LastName:      "Doe",
		StreetAddress: "123 Any Street",
		City:          "Pittsburgh",
		State:         "PA",
		ZipCode:       "12345",
		Email:         "John.Doe@example.com",
		HomeScene:     "Swing City",
		IsStudent:     true,
		PassType: &add.WeekendPass{
			Level: add.WeekendPassLevel2,
			Tier:  add.WeekendPassTier3,
		},
		MixAndMatch:     nil,
		SoloJazz:        true,
		TeamCompetition: nil,
		TShirt:          nil,
		Housing: &add.ProvideHousing{
			Pets:     pets,
			Quantity: quantity,
			Details:  details,
		},
		ReferenceId: referenceId,
		Paid:        true,
	}

	key, err := NewDatastore(store).AddRegistration(context.Background(), registration)
	if err != nil {
		t.Fatalf("error inserting new registration into store: %v", err)
	}
	decodedKey, err := datastore.DecodeKey(key)
	if err != nil {
		t.Log("WARNING:  Information leaked into database during test (test not cleaned up)")
		t.Fatalf("could not decode key %s: %v", key, err)
	}
	defer func() {
		err := store.Delete(context.Background(), decodedKey)
		if err != nil {
			t.Log("WARNING:  Information leaked into database during test (test not cleaned up)")
			t.Fatalf("error cleaning up test item from store: %v", err)
		}
	}()

	entity := &registrationEntity{}
	err = store.Get(context.Background(), decodedKey, entity)
	if err != nil {
		t.Fatalf("error finding newly added registration in store: %v", err)
	}

	if entity.HousingRequest != providesHousing {
		t.Fatalf("found housing request %s, expected %s", entity.HousingRequest, providesHousing)
	}

	if entity.ProvideHousing.Pets != pets {
		t.Fatalf("found pet allergies %s, expected %s", entity.ProvideHousing.Pets, pets)
	}

	if entity.ProvideHousing.Details != details {
		t.Fatalf("found pet allergies %s, expected %s", entity.ProvideHousing.Details, details)
	}

	if entity.ProvideHousing.Quantity != quantity {
		t.Fatalf("found housing quantity %d, expected %d", entity.ProvideHousing.Quantity, quantity)
	}

	if entity.WeekendPass != fullWeekendPass {
		t.Fatalf("found weekend pass status %s, expected %s", entity.WeekendPass, fullWeekendPass)
	}

	if entity.FullWeekendPassInfo.Level != int(add.WeekendPassLevel2) {
		t.Fatalf("found weekend pass level %d, expected %d", entity.FullWeekendPassInfo.Level, int(add.WeekendPassLevel2))
	}

	if entity.FullWeekendPassInfo.Tier != int(add.WeekendPassTier3) {
		t.Fatalf("found weekend pass tier %d, expected %d", entity.FullWeekendPassInfo.Tier, int(add.WeekendPassTier3))
	}
}
