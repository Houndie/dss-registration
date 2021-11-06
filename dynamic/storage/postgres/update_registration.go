package postgres

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/gofrs/uuid"
)

var updateRegistrationStmt string

func init() {
	tmplStmt := `UPDATE {{.Table}}
		SET 
			{{.FirstNameCol}} = $1,
			{{.LastNameCol}} = $2,
			{{.StreetAddressCol}} = $3,
			{{.CityCol}} = $4,
			{{.StateCol}} = $5,
			{{.ZipCodeCol}} = $6,
			{{.EmailCol}} = $7,
			{{.HomeSceneCol}} = $8,
			{{.IsStudentCol}} = $9,
			{{.PassTypeCol}} = $10,
			{{.FullWeekendLevelCol}} = $11,
			{{.FullWeekendTierCol}} = $12,
			{{.PassManuallyPaidCol}} = $13,
			{{.MixAndMatchCol}} = $14,
			{{.MixAndMatchRoleCol}} = $15,
			{{.MixAndMatchManuallyPaidCol}} = $16,
			{{.SoloJazzCol}} = $17,
			{{.SoloJazzManuallyPaidCol}} = $18,
			{{.TeamCompetitionCol}} = $19,
			{{.TeamCompetitionNameCol}} = $20,
			{{.TeamCompetitionManuallyPaidCol}} = $21,
			{{.TShirtCol}} = $22,
			{{.TShirtStyleCol}} = $23,
			{{.TShirtManuallyPaidCol}} = $24,
			{{.HousingCol}} = $25,
			{{.ProvideHousingPetsCol}} = $26,
			{{.ProvideHousingQuantityCol}} = $27,
			{{.ProvideHousingDetailsCol}} = $28,
			{{.RequireHousingPetAllergiesCol}} = $29,
			{{.RequireHousingDetailsCol}} = $30,
			{{.UserIDCol}} = $31,
			{{.OrderIDsCol}} = $32,
			{{.DiscountCodesCol}} = $33
		WHERE {{.IDCol}} = $34;`

	tmpl, err := template.New("tmpl").Parse(tmplStmt)
	if err != nil {
		panic(fmt.Sprintf("error parsing add admin template: %v", err))
	}

	stmt := &bytes.Buffer{}
	err = tmpl.Execute(stmt, registrationConsts)
	if err != nil {
		panic(fmt.Sprintf("error executing add admin template: %v", err))
	}

	updateRegistrationStmt = stmt.String()
}

func (s *Store) UpdateRegistration(ctx context.Context, registration *storage.Registration) error {
	id, err := uuid.FromString(registration.ID)
	if err != nil {
		return storage.ErrNoRegistrationForID{ID: registration.ID}
	}
	passType, fullWeekendLevel, fullWeekendTier, passManuallyPaid, err := toDBPassType(registration.PassType)
	if err != nil {
		return err
	}

	mixAndMatch := false
	var mixAndMatchRole *string
	mixAndMatchManuallyPaid := false
	if registration.MixAndMatch != nil {
		mixAndMatch = true
		mixAndMatchRoleStr, ok := roleToEnum[registration.MixAndMatch.Role]
		if !ok {
			return fmt.Errorf("unknown mix and match role found: %v", registration.MixAndMatch.Role)
		}
		mixAndMatchRole = &mixAndMatchRoleStr
		mixAndMatchManuallyPaid = registration.MixAndMatch.ManuallyPaid
	}

	soloJazz := false
	soloJazzManuallyPaid := false
	if registration.SoloJazz != nil {
		soloJazz = true
		soloJazzManuallyPaid = registration.SoloJazz.ManuallyPaid
	}

	teamCompetition := false
	var teamCompetitionName string
	teamCompetitionManuallyPaid := false
	if registration.TeamCompetition != nil {
		teamCompetition = true
		teamCompetitionName = registration.TeamCompetition.Name
		teamCompetitionManuallyPaid = registration.TeamCompetition.ManuallyPaid
	}

	tshirt := false
	var tshirtStyle *string
	tshirtManuallyPaid := false
	if registration.TShirt != nil {
		tshirt = true
		tshirtStyleStr, ok := styleToEnum[registration.TShirt.Style]
		if !ok {
			return fmt.Errorf("unknown tshirt style found: %v", registration.TShirt.Style)
		}
		tshirtStyle = &tshirtStyleStr
		tshirtManuallyPaid = registration.TShirt.ManuallyPaid
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

	ct, err := s.pool.Exec(ctx, updateRegistrationStmt,
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
		passManuallyPaid,
		mixAndMatch,
		mixAndMatchRole,
		mixAndMatchManuallyPaid,
		soloJazz,
		soloJazzManuallyPaid,
		teamCompetition,
		teamCompetitionName,
		teamCompetitionManuallyPaid,
		tshirt,
		tshirtStyle,
		tshirtManuallyPaid,
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
