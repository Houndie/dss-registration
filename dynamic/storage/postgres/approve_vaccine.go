package postgres

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

	"github.com/Houndie/dss-registration/dynamic/storage"
	"github.com/gofrs/uuid"
)

var approveVaccineStmt string

func init() {
	tmplStmt := `UPDATE {{.Table}} SET {{.VaxApprovedCol}} = $1 WHERE {{.IDCol}} = $2;`

	tmpl, err := template.New("tmpl").Parse(tmplStmt)
	if err != nil {
		panic(fmt.Sprintf("error parsing approve vaccine template: %v", err))
	}

	stmt := &bytes.Buffer{}
	err = tmpl.Execute(stmt, registrationConsts)
	if err != nil {
		panic(fmt.Sprintf("error executing approve vaccine template: %v", err))
	}

	approveVaccineStmt = stmt.String()
}

func (s *Store) ApproveVaccine(ctx context.Context, id string, approval bool) error {
	uuidID, err := uuid.FromString(id)
	if err != nil {
		return storage.ErrNoRegistrationForID{ID: id}
	}

	ct, err := s.pool.Exec(ctx, approveVaccineStmt, approval, uuidID)
	if err != nil {
		return fmt.Errorf("error approving vaccine in database: %w", err)
	}

	if ct.RowsAffected() == 0 {
		return storage.ErrNoRegistrationForID{ID: id}
	}

	return nil
}
