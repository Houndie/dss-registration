package datastore

import (
	"fmt"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/Houndie/dss-registration/dynamic/storage"
)

const (
	registrationKind = "Registration"

	fullWeekendPass = "Full Weekend Pass"
	danceOnlyPass   = "Dance Only Pass"
	noPass          = "No Pass"

	requiresHousing = "Requires Housing"
	providesHousing = "Provides Housing"
	noHousing       = "No Housing"
)

type registrationEntity struct {
	FirstName      string
	LastName       string
	StreetAddress  string
	City           string
	State          string
	ZipCode        string
	Email          string
	HomeScene      string
	IsStudent      bool
	SoloJazz       bool
	HousingRequest string
	RequireHousing struct {
		PetAllergies string
		Details      string
	}
	ProvideHousing struct {
		Pets     string
		Quantity int
		Details  string
	}
	WantsTShirt         bool
	TShirtStyle         string
	HasTeamCompetition  bool
	TeamCompetitionName string
	HasMixAndMatch      bool
	MixAndMatchRole     string
	WeekendPass         string
	FullWeekendPassInfo struct {
		Level int
		Tier  int
	}
	UserId        string
	OrderIds      []string
	CreatedAt     string
	DiscountCodes []string
	Disabled      bool
}

func toRegistrationEntity(r *storage.Registration) (*datastore.Key, *registrationEntity, error) {
	registration := &registrationEntity{
		FirstName:     r.FirstName,
		LastName:      r.LastName,
		StreetAddress: r.StreetAddress,
		City:          r.City,
		State:         r.State,
		ZipCode:       r.ZipCode,
		Email:         r.Email,
		HomeScene:     r.HomeScene,
		IsStudent:     r.IsStudent,
		SoloJazz:      r.SoloJazz,
		UserId:        r.UserId,
		OrderIds:      r.OrderIds,
		DiscountCodes: r.DiscountCodes,
		CreatedAt:     r.CreatedAt.Format(time.RFC3339),
		Disabled:      r.Disabled,
	}

	switch p := r.PassType.(type) {
	case *storage.WeekendPass:
		registration.WeekendPass = fullWeekendPass
		registration.FullWeekendPassInfo.Level = int(p.Level)
		registration.FullWeekendPassInfo.Tier = int(p.Tier)
	case *storage.DanceOnlyPass:
		registration.WeekendPass = danceOnlyPass
	case *storage.NoPass:
		registration.WeekendPass = noPass
	default:
		return nil, nil, fmt.Errorf("Found unknown type of weekend pass")
	}

	if r.MixAndMatch != nil {
		registration.HasMixAndMatch = true
		registration.MixAndMatchRole = string(r.MixAndMatch.Role)
	}

	if r.TeamCompetition != nil {
		registration.HasTeamCompetition = true
		registration.TeamCompetitionName = r.TeamCompetition.Name
	}

	if r.TShirt != nil {
		registration.WantsTShirt = true
		registration.TShirtStyle = string(r.TShirt.Style)
	}

	switch h := r.Housing.(type) {
	case *storage.ProvideHousing:
		registration.HousingRequest = providesHousing
		registration.ProvideHousing.Pets = h.Pets
		registration.ProvideHousing.Quantity = h.Quantity
		registration.ProvideHousing.Details = h.Details
	case *storage.RequireHousing:
		registration.HousingRequest = requiresHousing
		registration.RequireHousing.PetAllergies = h.PetAllergies
		registration.RequireHousing.Details = h.Details
	case *storage.NoHousing:
		registration.HousingRequest = noHousing
	default:
		return nil, nil, fmt.Errorf("Found unknown type of housing")
	}

	var key *datastore.Key
	if r.ID != "" {
		var err error
		key, err = datastore.DecodeKey(r.ID)
		if err != nil {
			return nil, nil, fmt.Errorf("error decoding registration ID: %w", err)
		}
	}

	return key, registration, nil
}

func fromRegistrationEntity(key *datastore.Key, re *registrationEntity) (*storage.Registration, error) {
	createdAt, err := time.Parse(time.RFC3339, re.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("error converting registration created at %s to understandable time: %w", re.CreatedAt, err)
	}

	var passType storage.PassType
	switch re.WeekendPass {
	case fullWeekendPass:
		passType = &storage.WeekendPass{
			Level: storage.WeekendPassLevel(re.FullWeekendPassInfo.Level),
			Tier:  storage.WeekendPassTier(re.FullWeekendPassInfo.Tier),
		}
	case danceOnlyPass:
		passType = &storage.DanceOnlyPass{}
	case noPass:
		passType = &storage.NoPass{}
	}

	var mixAndMatch *storage.MixAndMatch
	if re.HasMixAndMatch {
		mixAndMatch = &storage.MixAndMatch{
			Role: storage.MixAndMatchRole(re.MixAndMatchRole),
		}
	}

	var teamCompetition *storage.TeamCompetition
	if re.HasTeamCompetition {
		teamCompetition = &storage.TeamCompetition{
			Name: re.TeamCompetitionName,
		}
	}

	var tShirt *storage.TShirt
	if re.WantsTShirt {
		tShirt = &storage.TShirt{
			Style: storage.TShirtStyle(re.TShirtStyle),
		}
	}

	var housing storage.Housing
	switch re.HousingRequest {
	case requiresHousing:
		housing = &storage.RequireHousing{
			PetAllergies: re.RequireHousing.PetAllergies,
			Details:      re.RequireHousing.Details,
		}
	case providesHousing:
		housing = &storage.ProvideHousing{
			Pets:     re.ProvideHousing.Pets,
			Quantity: re.ProvideHousing.Quantity,
			Details:  re.ProvideHousing.Details,
		}
	case noHousing:
		housing = &storage.NoHousing{}
	}

	return &storage.Registration{
		ID:              key.Encode(),
		FirstName:       re.FirstName,
		LastName:        re.LastName,
		StreetAddress:   re.StreetAddress,
		City:            re.City,
		State:           re.State,
		ZipCode:         re.ZipCode,
		Email:           re.Email,
		HomeScene:       re.HomeScene,
		IsStudent:       re.IsStudent,
		PassType:        passType,
		MixAndMatch:     mixAndMatch,
		SoloJazz:        re.SoloJazz,
		TeamCompetition: teamCompetition,
		TShirt:          tShirt,
		Housing:         housing,
		UserId:          re.UserId,
		OrderIds:        re.OrderIds,
		CreatedAt:       createdAt,
		DiscountCodes:   re.DiscountCodes,
	}, nil
}
