package postgres

import (
	"context"
	"fmt"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/gofrs/uuid"
)

func (s *Store) UpdateRegistration(ctx context.Context, registration *storage.Registration) error {
	id, err := uuid.FromString(registration.ID)
	if err != nil {
		return storage.ErrNoRegistrationForID{ID: registration.ID}
	}
	passType, fullWeekendLevel, fullWeekendTier, err := toDBPassType(registration.PassType)
	if err != nil {
		return err
	}

	mixAndMatch := false
	var mixAndMatchRole *string
	if registration.MixAndMatch != nil {
		mixAndMatch = true
		mixAndMatchRoleStr, ok := roleToEnum[registration.MixAndMatch.Role]
		if !ok {
			return fmt.Errorf("unknown mix and match role found: %v", registration.MixAndMatch.Role)
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
			return fmt.Errorf("unknown tshirt style found: %v", registration.TShirt.Style)
		}
		tshirtStyle = &tshirtStyleStr
	}

	housing, provideHousingPets, provideHousingQuantity, provideHousingDetails, requireHousingPetAllergies, requireHousingDetails, err := toDBHousingType(registration.Housing)
	if err != nil {
		return err
	}

	orderIDs := []string{}
	if len(registration.OrderIDs) > 0 {
		orderIDs = registration.OrderIDs
	}

	discountCodes := []string{}
	if len(registration.DiscountCodes) > 0 {
		discountCodes = registration.DiscountCodes
	}

	ct, err := s.pool.Exec(ctx, fmt.Sprintf("UPDATE %s SET %s = $1, %s = $2, %s = $3, %s = $4, %s = $5, %s = $6, %s = $7, %s = $8, %s = $9, %s = $10, %s = $11, %s = $12, %s = $13, %s = $14, %s = $15, %s = $16, %s = $17, %s = $18, %s = $19, %s = $20, %s = $21, %s = $22, %s = $23, %s = $24, %s = $25, %s = $26, %s = $27, %s = $28 WHERE %s = $29",
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
		registration.UserID,
		orderIDs,
		discountCodes,
		id)
	if err != nil {
		return fmt.Errorf("error adding new registration to database: %w", err)
	}
	if ct.RowsAffected() == 0 {
		return storage.ErrNoRegistrationForID{ID: registration.ID}
	}
	return nil
}
