package datastore

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/registration/add"
	"github.com/Houndie/dss-registration/dynamic/registration/common"
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

	userId, err := uuid.NewV4()
	if err != nil {
		t.Fatalf("error generating user id for test: %v", err)
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
		PassType:        &common.NoPass{},
		MixAndMatch:     nil,
		SoloJazz:        false,
		TeamCompetition: nil,
		TShirt:          nil,
		Housing:         &common.NoHousing{},
		UserId:          userId.String(),
	}

	_, err = NewDatastore(store).AddRegistration(context.Background(), registration)
	if err != nil {
		t.Fatalf("error inserting new registration into store: %v", err)
	}

	q := datastore.NewQuery(registrationKind).Filter("UserId =", userId.String())
	entities := []registrationEntity{}
	keys, err := store.GetAll(context.Background(), q, &entities)
	if err != nil {
		t.Log("WARNING:  Information leaked into database during test (test not cleaned up)")
		t.Fatalf("could find database item with id %v: %v", userId, err)
	}
	defer func() {
		for _, key := range keys {
			err := store.Delete(context.Background(), key)
			if err != nil {
				t.Log("WARNING:  Information leaked into database during test (test not cleaned up)")
				t.Fatalf("error cleaning up test item from store: %v", err)
			}
		}
	}()
	if len(keys) != 1 {
		t.Fatalf("found incorrect number of test item keys %d", len(keys))
	}
	if len(entities) != 1 {
		t.Fatalf("found incorrect number of test items %d", len(keys))
	}
	entity := entities[0]

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

	if entity.UserId != registration.UserId {
		t.Fatalf("found user id %v, expected %v", entity.UserId, registration.UserId)
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

	userId, err := uuid.NewV4()
	if err != nil {
		t.Fatalf("error generating user id for test: %v", err)
	}

	role := common.MixAndMatchRoleLeader
	name := "Awesome team"
	style := common.TShirtStyleBellaM
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
		PassType:      &common.DanceOnlyPass{},
		MixAndMatch: &common.MixAndMatch{
			Role: role,
		},
		SoloJazz: true,
		TeamCompetition: &common.TeamCompetition{
			Name: name,
		},
		TShirt: &common.TShirt{
			Style: style,
		},
		Housing: &common.RequireHousing{
			PetAllergies: petAllergies,
			Details:      details,
		},
		UserId: userId.String(),
	}

	_, err = NewDatastore(store).AddRegistration(context.Background(), registration)
	if err != nil {
		t.Fatalf("error inserting new registration into store: %v", err)
	}

	q := datastore.NewQuery(registrationKind).Filter("UserId =", userId.String())
	entities := []registrationEntity{}
	keys, err := store.GetAll(context.Background(), q, &entities)
	if err != nil {
		t.Log("WARNING:  Information leaked into database during test (test not cleaned up)")
		t.Fatalf("could find database item with id %v: %v", userId, err)
	}
	defer func() {
		for _, key := range keys {
			err := store.Delete(context.Background(), key)
			if err != nil {
				t.Log("WARNING:  Information leaked into database during test (test not cleaned up)")
				t.Fatalf("error cleaning up test item from store: %v", err)
			}
		}
	}()
	if len(keys) != 1 {
		t.Fatalf("found incorrect number of test item keys %d", len(keys))
	}
	if len(entities) != 1 {
		t.Fatalf("found incorrect number of test items %d", len(keys))
	}
	entity := entities[0]

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
}

func TestAddRegistration3(t *testing.T) {
	if os.Getenv("DATASTORE_EMULATOR_HOST") == "" {
		t.Skip()
	}

	store, err := datastore.NewClient(context.Background(), os.Getenv("DATASTORE_PROJECT_ID"))
	if err != nil {
		t.Fatalf("could not connect to datastore emulator: %v", err)
	}

	userId, err := uuid.NewV4()
	if err != nil {
		t.Fatalf("error generating user id for test: %v", err)
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
		PassType: &common.WeekendPass{
			Level: common.WeekendPassLevel2,
			Tier:  common.WeekendPassTier3,
		},
		MixAndMatch:     nil,
		SoloJazz:        true,
		TeamCompetition: nil,
		TShirt:          nil,
		Housing: &common.ProvideHousing{
			Pets:     pets,
			Quantity: quantity,
			Details:  details,
		},
		UserId: userId.String(),
	}

	_, err = NewDatastore(store).AddRegistration(context.Background(), registration)
	if err != nil {
		t.Fatalf("error inserting new registration into store: %v", err)
	}

	q := datastore.NewQuery(registrationKind).Filter("UserId =", userId.String())
	entities := []registrationEntity{}
	keys, err := store.GetAll(context.Background(), q, &entities)
	if err != nil {
		t.Log("WARNING:  Information leaked into database during test (test not cleaned up)")
		t.Fatalf("could find database item with id %v: %v", userId, err)
	}
	defer func() {
		for _, key := range keys {
			err := store.Delete(context.Background(), key)
			if err != nil {
				t.Log("WARNING:  Information leaked into database during test (test not cleaned up)")
				t.Fatalf("error cleaning up test item from store: %v", err)
			}
		}
	}()
	if len(keys) != 1 {
		t.Fatalf("found incorrect number of test item keys %d", len(keys))
	}
	if len(entities) != 1 {
		t.Fatalf("found incorrect number of test items %d", len(keys))
	}
	entity := entities[0]

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

	if entity.FullWeekendPassInfo.Level != int(common.WeekendPassLevel2) {
		t.Fatalf("found weekend pass level %d, expected %d", entity.FullWeekendPassInfo.Level, int(common.WeekendPassLevel2))
	}

	if entity.FullWeekendPassInfo.Tier != int(common.WeekendPassTier3) {
		t.Fatalf("found weekend pass tier %d, expected %d", entity.FullWeekendPassInfo.Tier, int(common.WeekendPassTier3))
	}
}