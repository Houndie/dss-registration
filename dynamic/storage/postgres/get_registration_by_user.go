package postgres

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

	"github.com/Houndie/dss-registration/dynamic/storage"
)

var getRegistrationsByUserStmt string

func init() {
	tmplStmt := `SELECT
			{{.IDCol}},
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
			{{.MixAndMatchCol}},
			{{.MixAndMatchRoleCol}},
			{{.SoloJazzCol}},
			{{.TeamCompetitionCol}},
			{{.TeamCompetitionNameCol}},
			{{.TShirtCol}},
			{{.TShirtStyleCol}},
			{{.HousingCol}},
			{{.ProvideHousingPetsCol}},
			{{.ProvideHousingQuantityCol}},
			{{.ProvideHousingDetailsCol}},
			{{.RequireHousingPetAllergiesCol}},
			{{.RequireHousingDetailsCol}},
			{{.OrderIDsCol}},
			{{.DiscountCodesCol}},
			{{.EnabledCol}}
		FROM {{.Table}}
		WHERE {{.UserIDCol}} = $1;`
	tmpl, err := template.New("tmpl").Parse(tmplStmt)
	if err != nil {
		panic(fmt.Sprintf("error parsing get registration by user template: %v", err))
	}

	stmt := &bytes.Buffer{}
	err = tmpl.Execute(stmt, registrationConsts)
	if err != nil {
		panic(fmt.Sprintf("error executing get registration by user template: %v", err))
	}

	getRegistrationsByUserStmt = stmt.String()
}

func (s *Store) GetRegistrationsByUser(ctx context.Context, userID string) ([]*storage.Registration, error) {
	rows, err := s.pool.Query(ctx, getRegistrationsByUserStmt, userID)
	if err != nil {
		return nil, fmt.Errorf("error querying for registrations by user id: %w", err)
	}
	defer rows.Close()
	registrations := []*storage.Registration{}
	for rows.Next() {
		r := &storage.Registration{
			UserID: userID,
		}
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
		err := rows.Scan(
			&r.ID,
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
			&r.OrderIDs,
			&r.DiscountCodes,
			&r.Enabled)
		if err != nil {
			return nil, fmt.Errorf("error parsing registration row: %w", err)
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
		registrations = append(registrations, r)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error on closing registration query: %w", err)
	}

	return registrations, nil
}
