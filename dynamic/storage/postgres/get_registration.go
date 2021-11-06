package postgres

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"text/template"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
)

var getRegistrationStmt string

func init() {
	tmplStmt := `SELECT
			{{.CreatedAtCol}},
			{{.FirstNameCol}},
			{{.LastNameCol}},
			{{.StreetAddressCol}},
			{{.CityCol}},
			{{.StateCol}},
			{{.ZipCodeCol}},
			{{.EmailCol}},
			{{.HomeSceneCol}},
			{{.IsStudentCol}},
			{{.PassTypeCol}},
			{{.FullWeekendLevelCol}},
			{{.FullWeekendTierCol}},
			{{.PassManuallyPaidCol}},
			{{.MixAndMatchCol}},
			{{.MixAndMatchRoleCol}},
			{{.MixAndMatchManuallyPaidCol}},
			{{.SoloJazzCol}},
			{{.SoloJazzManuallyPaidCol}},
			{{.TeamCompetitionCol}},
			{{.TeamCompetitionNameCol}},
			{{.TeamCompetitionManuallyPaidCol}},
			{{.TShirtCol}},
			{{.TShirtStyleCol}},
			{{.TShirtManuallyPaidCol}},
			{{.HousingCol}},
			{{.ProvideHousingPetsCol}},
			{{.ProvideHousingQuantityCol}},
			{{.ProvideHousingDetailsCol}},
			{{.RequireHousingPetAllergiesCol}},
			{{.RequireHousingDetailsCol}},
			{{.UserIDCol}},
			{{.OrderIDsCol}},
			{{.DiscountCodesCol}},
			{{.EnabledCol}}
		FROM {{.Table}}
		WHERE {{.IDCol}} = $1
		LIMIT 1;`
	tmpl, err := template.New("tmpl").Parse(tmplStmt)
	if err != nil {
		panic(fmt.Sprintf("error parsing get registration template: %v", err))
	}

	stmt := &bytes.Buffer{}
	err = tmpl.Execute(stmt, registrationConsts)
	if err != nil {
		panic(fmt.Sprintf("error executing get registration template: %v", err))
	}

	getRegistrationStmt = stmt.String()
}

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
	var passManuallyPaid bool
	var mixAndMatch bool
	var mixAndMatchRole *string
	var mixAndMatchManuallyPaid bool
	var soloJazz bool
	var soloJazzManuallyPaid bool
	var teamCompetition bool
	var teamCompetitionName string
	var teamCompetitionManuallyPaid bool
	var tshirt bool
	var tshirtStyle *string
	var tshirtManuallyPaid bool
	var housing string
	var provideHousingPets string
	var provideHousingQuantity int
	var provideHousingDetails string
	var requireHousingPetAllergies string
	var requireHousingDetails string
	err = s.pool.QueryRow(ctx, getRegistrationStmt, uuidID).Scan(
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
		&passManuallyPaid,
		&mixAndMatch,
		&mixAndMatchRole,
		&mixAndMatchManuallyPaid,
		&soloJazz,
		&soloJazzManuallyPaid,
		&teamCompetition,
		&teamCompetitionName,
		&teamCompetitionManuallyPaid,
		&tshirt,
		&tshirtStyle,
		&tshirtManuallyPaid,
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

	r.PassType, err = fromDBPassType(passType, fullWeekendTier, fullWeekendLevel, passManuallyPaid)
	if err != nil {
		return nil, err
	}

	if mixAndMatch {
		role, ok := roleFromEnum[*mixAndMatchRole]
		if !ok {
			return nil, fmt.Errorf("unknown mix and match role found: %v", *mixAndMatchRole)
		}
		r.MixAndMatch = &storage.MixAndMatch{
			Role:         role,
			ManuallyPaid: mixAndMatchManuallyPaid,
		}
	}

	if soloJazz {
		r.SoloJazz = &storage.SoloJazz{
			ManuallyPaid: soloJazzManuallyPaid,
		}
	}

	if teamCompetition {
		r.TeamCompetition = &storage.TeamCompetition{
			Name:         teamCompetitionName,
			ManuallyPaid: teamCompetitionManuallyPaid,
		}
	}

	if tshirt {
		style, ok := styleFromEnum[*tshirtStyle]
		if !ok {
			return nil, fmt.Errorf("unknown tshirt style found: %v", *tshirtStyle)
		}
		r.TShirt = &storage.TShirt{
			Style:        style,
			ManuallyPaid: teamCompetitionManuallyPaid,
		}
	}

	r.Housing, err = fromDBHousingType(housing, provideHousingPets, provideHousingQuantity, provideHousingDetails, requireHousingPetAllergies, requireHousingDetails)
	if err != nil {
		return nil, err
	}

	return r, nil
}
