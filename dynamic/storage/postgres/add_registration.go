package postgres

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/gofrs/uuid"
)

var addRegistrationStmt string

func init() {
	tmplStmt := `INSERT INTO {{ .Table }}(
			{{ .FirstNameCol }}, 
			{{ .LastNameCol }},
			{{ .StreetAddressCol }},
			{{ .CityCol }},
			{{ .StateCol }},
			{{ .ZipCodeCol }},
			{{ .EmailCol }},
			{{ .HomeSceneCol }},
			{{ .IsStudentCol }},
			{{ .PassTypeCol }},
			{{ .FullWeekendLevelCol }},
			{{ .FullWeekendTierCol }},
			{{ .MixAndMatchCol }},
			{{ .MixAndMatchRoleCol }},
			{{ .SoloJazzCol }},
			{{ .TeamCompetitionCol }},
			{{ .TeamCompetitionNameCol }},
			{{ .TShirtCol }},
			{{ .TShirtStyleCol }},
			{{ .HousingCol }},
			{{ .ProvideHousingPetsCol }},
			{{ .ProvideHousingQuantityCol }},
			{{ .ProvideHousingDetailsCol }},
			{{ .RequireHousingPetAllergiesCol }},
			{{ .RequireHousingDetailsCol }},
			{{ .UserIDCol }},
			{{ .OrderIDsCol }},
			{{ .DiscountCodesCol }})
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28)
		RETURNING {{ .IDCol }};`

	tmpl, err := template.New("tmpl").Parse(tmplStmt)
	if err != nil {
		panic(fmt.Sprintf("error parsing add admin template: %v", err))
	}

	stmt := &bytes.Buffer{}
	err = tmpl.Execute(stmt, registrationConsts)
	if err != nil {
		panic(fmt.Sprintf("error executing add admin template: %v", err))
	}

	addRegistrationStmt = stmt.String()
}

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
	if registration.TShirt != nil {
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
	if len(registration.OrderIDs) > 0 {
		orderIDs = registration.OrderIDs
	}

	discountCodes := []string{}
	if len(registration.DiscountCodes) > 0 {
		discountCodes = registration.DiscountCodes
	}

	var id uuid.UUID
	err = s.pool.QueryRow(ctx, addRegistrationStmt,
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
		discountCodes).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("error adding new registration to database: %w", err)
	}
	return id.String(), nil
}
