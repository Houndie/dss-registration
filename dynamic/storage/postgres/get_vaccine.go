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

var getVaccineStmt string

func init() {
	tmplStmt := `SELECT {{.VaxApprovedCol}} FROM {{.Table}} WHERE {{.IDCol}} = $1 LIMIT 1;`

	tmpl, err := template.New("tmpl").Parse(tmplStmt)
	if err != nil {
		panic(fmt.Sprintf("error parsing get vaccine template: %v", err))
	}

	stmt := &bytes.Buffer{}
	err = tmpl.Execute(stmt, registrationConsts)
	if err != nil {
		panic(fmt.Sprintf("error executing get vaccine template: %v", err))
	}

	getVaccineStmt = stmt.String()
}

func (s *Store) GetVaccine(ctx context.Context, id string) (bool, error) {
	uuidID, err := uuid.FromString(id)
	if err != nil {
		return false, storage.ErrNoRegistrationForID{ID: id}
	}

	vaxApproved := false
	err = s.pool.QueryRow(ctx, getVaccineStmt, uuidID).Scan(&vaxApproved)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, storage.ErrNoRegistrationForID{ID: id}
		}
		return false, fmt.Errorf("error getting vaccine from database: %w", err)
	}

	return vaxApproved, nil
}
