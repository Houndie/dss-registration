package postgres

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/gofrs/uuid"
)

func (s *Store) AddRegistration(ctx context.Context, registration *storage.Registration) (string, error) {
	passType, fullWeekendLevel, fullWeekendTier, err := toDBPassType(registration.PassType)
	if err != nil {
		return "", err
	}

	mixAndMatch := false
	var mixAndMatchRole *string
	if registration.MixAndMatch != nil {
		mixAndMatch = true
		mixAndMatchRoleStr, ok := roleToEnum[registration.MixAndMatch.Role]
		if !ok {
			return "", fmt.Errorf("unknown mix and match role found: %v", registration.MixAndMatch.Role)
		}
		mixAndMatchRole = &mixAndMatchRoleStr
	}

	teamCompetition := false
	var teamCompetitionName string
	if registration.TeamCompetition != nil {
		teamCompetition = true
		teamCompetitionName = registration.TeamCompetition.Name
	}

	tshirt := false
	var tshirtStyle *string
	if registration.MixAndMatch != nil {
		tshirt = true
		tshirtStyleStr, ok := styleToEnum[registration.TShirt.Style]
		if !ok {
			return "", fmt.Errorf("unknown tshirt style found: %v", registration.TShirt.Style)
		}
		tshirtStyle = &tshirtStyleStr
	}

	housing, provideHousingPets, provideHousingQuantity, provideHousingDetails, requireHousingPetAllergies, requireHousingDetails, err := toDBHousingType(registration.Housing)
	if err != nil {
		return "", err
	}

	orderIDs := []string{}
	if len(registration.OrderIds) > 0 {
		orderIDs = registration.OrderIds
	}

	discountCodes := []string{}
	if len(registration.DiscountCodes) > 0 {
		discountCodes = registration.DiscountCodes
	}

	var id uuid.UUID
	err = s.pool.QueryRow(ctx, fmt.Sprintf("INSERT INTO %s(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28) RETURNING %s",
		registrationTable,
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
		registrationIDCol),
		registration.FirstName,
		registration.LastName,
		registration.StreetAddress,
		registration.City,
		registration.State,
		registration.ZipCode,
		registration.Email,
		registration.HomeScene,
		registration.IsStudent,
		passType,
		fullWeekendLevel,
		fullWeekendTier,
		mixAndMatch,
		mixAndMatchRole,
		registration.SoloJazz,
		teamCompetition,
		teamCompetitionName,
		tshirt,
		tshirtStyle,
		housing,
		provideHousingPets,
		provideHousingQuantity,
		provideHousingDetails,
		requireHousingPetAllergies,
		requireHousingDetails,
		registration.UserId,
		orderIDs,
		discountCodes).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("error adding new registration to database: %w", err)
	}
	return id.String(), nil
}
