package postgres

import (
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/storage"
)

const (
	registrationTable                         = "registrations"
	registrationCreatedAtCol                  = "created_at"
	registrationEnabledCol                    = "enabled"
	registrationFirstNameCol                  = "first_name"
	registrationLastNameCol                   = "last_name"
	registrationStreetAddressCol              = "street_address"
	registrationCityCol                       = "city"
	registrationStateCol                      = "state"
	registrationZipCodeCol                    = "zip_code"
	registrationEmailCol                      = "email"
	registrationHomeSceneCol                  = "home_scene"
	registrationIsStudentCol                  = "is_student"
	registrationPassTypeCol                   = "pass_type"
	registrationFullWeekendLevelCol           = "full_weekend_level"
	registrationFullWeekendTierCol            = "full_weekend_tier"
	registrationMixAndMatchCol                = "mix_and_match"
	registrationMixAndMatchRoleCol            = "mix_and_match_role"
	registrationSoloJazzCol                   = "solo_jazz"
	registrationTeamCompetitionCol            = "team_competition"
	registrationTeamCompetitionNameCol        = "team_competition_name"
	registrationTShirtCol                     = "tshirt"
	registrationTShirtStyleCol                = "tshirt_style"
	registrationHousingCol                    = "housing"
	registrationProvideHousingPetsCol         = "provide_housing_pets"
	registrationProvideHousingQuantityCol     = "provide_housing_quantity"
	registrationProvideHousingDetailsCol      = "provide_housing_details"
	registrationRequireHousingPetAllergiesCol = "require_housing_pet_allergies"
	registrationRequireHousingDetailsCol      = "require_housing_details"
	registrationUserIDCol                     = "user_id"
	registrationOrderIDsCol                   = "order_ids"
	registrationDiscountCodesCol              = "discount_codes"
	registrationIDCol                         = "id"

	fullWeekendPassEnum = "Full Weekend"
	danceOnlyPassEnum   = "Dance Only"
	noPassEnum          = "No Pass"

	provideHousingEnum = "Provide"
	requireHousingEnum = "Require"
	noHousingEnum      = "No Housing"
)

type registrationConstsType struct {
	Table                         string
	CreatedAtCol                  string
	EnabledCol                    string
	FirstNameCol                  string
	LastNameCol                   string
	StreetAddressCol              string
	CityCol                       string
	StateCol                      string
	ZipCodeCol                    string
	EmailCol                      string
	HomeSceneCol                  string
	IsStudentCol                  string
	PassTypeCol                   string
	FullWeekendLevelCol           string
	FullWeekendTierCol            string
	MixAndMatchCol                string
	MixAndMatchRoleCol            string
	SoloJazzCol                   string
	TeamCompetitionCol            string
	TeamCompetitionNameCol        string
	TShirtCol                     string
	TShirtStyleCol                string
	HousingCol                    string
	ProvideHousingPetsCol         string
	ProvideHousingQuantityCol     string
	ProvideHousingDetailsCol      string
	RequireHousingPetAllergiesCol string
	RequireHousingDetailsCol      string
	UserIDCol                     string
	OrderIDsCol                   string
	DiscountCodesCol              string
	IDCol                         string
}

var registrationConsts = &registrationConstsType{
	Table:                         "registrations",
	CreatedAtCol:                  "created_at",
	EnabledCol:                    "enabled",
	FirstNameCol:                  "first_name",
	LastNameCol:                   "last_name",
	StreetAddressCol:              "street_address",
	CityCol:                       "city",
	StateCol:                      "state",
	ZipCodeCol:                    "zip_code",
	EmailCol:                      "email",
	HomeSceneCol:                  "home_scene",
	IsStudentCol:                  "is_student",
	PassTypeCol:                   "pass_type",
	FullWeekendLevelCol:           "full_weekend_level",
	FullWeekendTierCol:            "full_weekend_tier",
	MixAndMatchCol:                "mix_and_match",
	MixAndMatchRoleCol:            "mix_and_match_role",
	SoloJazzCol:                   "solo_jazz",
	TeamCompetitionCol:            "team_competition",
	TeamCompetitionNameCol:        "team_competition_name",
	TShirtCol:                     "tshirt",
	TShirtStyleCol:                "tshirt_style",
	HousingCol:                    "housing",
	ProvideHousingPetsCol:         "provide_housing_pets",
	ProvideHousingQuantityCol:     "provide_housing_quantity",
	ProvideHousingDetailsCol:      "provide_housing_details",
	RequireHousingPetAllergiesCol: "require_housing_pet_allergies",
	RequireHousingDetailsCol:      "require_housing_details",
	UserIDCol:                     "user_id",
	OrderIDsCol:                   "order_ids",
	DiscountCodesCol:              "discount_codes",
	IDCol:                         "id",
}

var (
	levelToEnum = map[storage.WeekendPassLevel]string{
		storage.Level1: "Level 1",
		storage.Level2: "Level 2",
		storage.Level3: "Level 3",
	}

	levelFromEnum = map[string]storage.WeekendPassLevel{
		"Level 1": storage.Level1,
		"Level 2": storage.Level2,
		"Level 3": storage.Level3,
	}

	tierToEnum = map[storage.WeekendPassTier]string{
		storage.Tier1: "Tier 1",
		storage.Tier2: "Tier 2",
		storage.Tier3: "Tier 3",
		storage.Tier4: "Tier 4",
		storage.Tier5: "Tier 5",
	}

	tierFromEnum = map[string]storage.WeekendPassTier{
		"Tier 1": storage.Tier1,
		"Tier 2": storage.Tier2,
		"Tier 3": storage.Tier3,
		"Tier 4": storage.Tier4,
		"Tier 5": storage.Tier5,
	}

	roleToEnum = map[storage.MixAndMatchRole]string{
		storage.MixAndMatchRoleLeader:   "Leader",
		storage.MixAndMatchRoleFollower: "Follower",
	}

	roleFromEnum = map[string]storage.MixAndMatchRole{
		"Leader":   storage.MixAndMatchRoleLeader,
		"Follower": storage.MixAndMatchRoleFollower,
	}

	styleToEnum = map[storage.TShirtStyle]string{
		storage.TShirtStyleUnisexS:   "Unisex S",
		storage.TShirtStyleUnisexM:   "Unisex M",
		storage.TShirtStyleUnisexL:   "Unisex L",
		storage.TShirtStyleUnisexXL:  "Unisex XL",
		storage.TShirtStyleUnisex2XL: "Unisex 2XL",
		storage.TShirtStyleUnisex3XL: "Unisex 3XL",
		storage.TShirtStyleBellaS:    "Bella S",
		storage.TShirtStyleBellaM:    "Bella M",
		storage.TShirtStyleBellaL:    "Bella L",
		storage.TShirtStyleBellaXL:   "Bella XL",
		storage.TShirtStyleBella2XL:  "Bella 2XL",
	}

	styleFromEnum = map[string]storage.TShirtStyle{
		"Unisex S":   storage.TShirtStyleUnisexS,
		"Unisex M":   storage.TShirtStyleUnisexM,
		"Unisex L":   storage.TShirtStyleUnisexL,
		"Unisex XL":  storage.TShirtStyleUnisexXL,
		"Unisex 2XL": storage.TShirtStyleUnisex2XL,
		"Unisex 3XL": storage.TShirtStyleUnisex3XL,
		"Bella S":    storage.TShirtStyleBellaS,
		"Bella M":    storage.TShirtStyleBellaM,
		"Bella L":    storage.TShirtStyleBellaL,
		"Bella XL":   storage.TShirtStyleBellaXL,
		"Bella 2XL":  storage.TShirtStyleBella2XL,
	}
)

func toDBPassType(p storage.PassType) (string, *string, *string, error) {
	switch pt := p.(type) {
	case *storage.WeekendPass:
		level, ok := levelToEnum[pt.Level]
		if !ok {
			return "", nil, nil, fmt.Errorf("unknown full weekend pass level found: %v", pt.Level)
		}
		tier, ok := tierToEnum[pt.Tier]
		if !ok {
			return "", nil, nil, fmt.Errorf("unknown full weekend pass tier found: %v", pt.Tier)
		}
		return fullWeekendPassEnum, &level, &tier, nil
	case *storage.DanceOnlyPass:
		return danceOnlyPassEnum, nil, nil, nil
	case *storage.NoPass:
		return noPassEnum, nil, nil, nil
	default:
		return "", nil, nil, fmt.Errorf("unknown pass type found %T", pt)
	}
}

func fromDBPassType(passType string, tierEnum *string, levelEnum *string) (storage.PassType, error) {
	switch passType {
	case fullWeekendPassEnum:
		level, ok := levelFromEnum[*levelEnum]
		if !ok {
			return nil, fmt.Errorf("unknown full weekend pass level enum found: %v", *levelEnum)
		}
		tier, ok := tierFromEnum[*tierEnum]
		if !ok {
			return nil, fmt.Errorf("unknown full weekend pass tier enum found: %v", *tierEnum)
		}
		return &storage.WeekendPass{
			Level: level,
			Tier:  tier,
		}, nil
	case danceOnlyPassEnum:
		return &storage.DanceOnlyPass{}, nil
	case noPassEnum:
		return &storage.NoPass{}, nil
	default:
		return nil, fmt.Errorf("unknown pass type enum found: %v", passType)
	}
}

func toDBHousingType(h storage.Housing) (string, string, int, string, string, string, error) {
	switch ht := h.(type) {
	case *storage.ProvideHousing:
		return provideHousingEnum, ht.Pets, ht.Quantity, ht.Details, "", "", nil
	case *storage.RequireHousing:
		return requireHousingEnum, "", 0, "", ht.PetAllergies, ht.Details, nil
	case *storage.NoHousing:
		return noHousingEnum, "", 0, "", "", "", nil
	default:
		return "", "", 0, "", "", "", fmt.Errorf("unknown housing type found: %T", ht)
	}
}

func fromDBHousingType(housingType, providePets string, provideQuantity int, provideDetails, requireAllergies, requireDetails string) (storage.Housing, error) {
	switch housingType {
	case provideHousingEnum:
		return &storage.ProvideHousing{
			Pets:     providePets,
			Quantity: provideQuantity,
			Details:  provideDetails,
		}, nil
	case requireHousingEnum:
		return &storage.RequireHousing{
			PetAllergies: requireAllergies,
			Details:      requireDetails,
		}, nil
	case noHousingEnum:
		return &storage.NoHousing{}, nil
	default:
		return nil, fmt.Errorf("unknown housing enum found: %v", housingType)
	}
}
