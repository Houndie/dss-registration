package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
)

func (s *Store) GetRegistration(ctx context.Context, id string) (*storage.Registration, error) {
	uuidID, err := uuid.FromString(id)
	if err != nil {
		return nil, storage.ErrNoRegistrationForID{ID: id}
	}

	r := &storage.Registration{}
	r.ID = uuidID.String()
	var passType string
	var fullWeekendLevel *string
	var fullWeekendTier *string
	var mixAndMatch bool
	var mixAndMatchRole *string
	var teamCompetition bool
	var teamCompetitionName string
	var tshirt bool
	var tshirtStyle *string
	var housing string
	var provideHousingPets string
	var provideHousingQuantity int
	var provideHousingDetails string
	var requireHousingPetAllergies string
	var requireHousingDetails string
	err = s.pool.QueryRow(ctx, fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s FROM %s WHERE %s = $1 LIMIT 1;",
		registrationCreatedAtCol,
		registrationFirstNameCol,
		registrationLastNameCol,
		registrationStreetAddressCol,
		registrationCityCol,
		registrationStateCol,
		registrationZipCodeCol,
		registrationEmailCol,
		registrationHomeSceneCol,
		registrationIsStudentCol,
		registrationPassTypeCol,
		registrationFullWeekendLevelCol,
		registrationFullWeekendTierCol,
		registrationMixAndMatchCol,
		registrationMixAndMatchRoleCol,
		registrationSoloJazzCol,
		registrationTeamCompetitionCol,
		registrationTeamCompetitionNameCol,
		registrationTShirtCol,
		registrationTShirtStyleCol,
		registrationHousingCol,
		registrationProvideHousingPetsCol,
		registrationProvideHousingQuantityCol,
		registrationProvideHousingDetailsCol,
		registrationRequireHousingPetAllergiesCol,
		registrationRequireHousingDetailsCol,
		registrationUserIDCol,
		registrationOrderIDsCol,
		registrationDiscountCodesCol,
		registrationEnabledCol,
		registrationTable,
		registrationIDCol),
		uuidID).Scan(
		&r.CreatedAt,
		&r.FirstName,
		&r.LastName,
		&r.StreetAddress,
		&r.City,
		&r.State,
		&r.ZipCode,
		&r.Email,
		&r.HomeScene,
		&r.IsStudent,
		&passType,
		&fullWeekendLevel,
		&fullWeekendTier,
		&mixAndMatch,
		&mixAndMatchRole,
		&r.SoloJazz,
		&teamCompetition,
		&teamCompetitionName,
		&tshirt,
		&tshirtStyle,
		&housing,
		&provideHousingPets,
		&provideHousingQuantity,
		&provideHousingDetails,
		&requireHousingPetAllergies,
		&requireHousingDetails,
		&r.UserID,
		&r.OrderIDs,
		&r.DiscountCodes,
		&r.Enabled)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, storage.ErrNoRegistrationForID{ID: id}
		}
		return nil, err
	}

	r.PassType, err = fromDBPassType(passType, fullWeekendTier, fullWeekendLevel)
	if err != nil {
		return nil, err
	}

	if mixAndMatch {
		role, ok := roleFromEnum[*mixAndMatchRole]
		if !ok {
			return nil, fmt.Errorf("unknown mix and match role found: %v", *mixAndMatchRole)
		}
		r.MixAndMatch = &storage.MixAndMatch{Role: role}
	}

	if teamCompetition {
		r.TeamCompetition = &storage.TeamCompetition{Name: teamCompetitionName}
	}

	if tshirt {
		style, ok := styleFromEnum[*tshirtStyle]
		if !ok {
			return nil, fmt.Errorf("unknown tshirt style found: %v", *tshirtStyle)
		}
		r.TShirt = &storage.TShirt{Style: style}
	}

	r.Housing, err = fromDBHousingType(housing, provideHousingPets, provideHousingQuantity, provideHousingDetails, requireHousingPetAllergies, requireHousingDetails)
	if err != nil {
		return nil, err
	}

	return r, nil
}
