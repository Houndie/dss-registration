package postgres

import (
	"bytes"
	"context"
	"fmt"
	"text/template"
)

var addAdminStmt string

func init() {
	tmplStmt := "INSERT INTO {{ .Table }} ({{ .UserIDCol }}) VALUES ($1);"

	tmpl, err := template.New("tmpl").Parse(tmplStmt)
	if err != nil {
		panic(fmt.Sprintf("error parsing add admin template: %v", err))
	}

	stmt := &bytes.Buffer{}
	err = tmpl.Execute(stmt, adminConsts)
	if err != nil {
		panic(fmt.Sprintf("error executing add admin template: %v", err))
	}

	addAdminStmt = stmt.String()
}

func (s *Store) AddAdmin(ctx context.Context, userID string) error {
	_, err := s.pool.Exec(ctx, addAdminStmt, userID)
	if err != nil {
		return fmt.Errorf("error adding new admin to db: %w", err)
	}
	return nil
}
